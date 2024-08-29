package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kilianp07/AthleteIQBox/data"
	// Import the sqlite3 package to register the database driver
	_ "github.com/mattn/go-sqlite3"
)

type Recorder struct {
	conf       configuration
	positionCh chan data.Position
	stopCh     chan bool
	errCh      chan error
	db         *sql.DB
}

func New() *Recorder {
	return &Recorder{
		conf:   configuration{},
		stopCh: make(chan bool, 1),
	}
}

func (r *Recorder) Conf() any {
	return &r.conf
}
func (r *Recorder) RuntimeErr() <-chan error {
	return r.errCh
}

func (r *Recorder) Configure(positionCh chan data.Position) error {
	r.positionCh = positionCh

	if positionCh == nil {
		return fmt.Errorf("position channel is nil")
	}

	// Check if the last caracter of filepath is a /
	if string(r.conf.Filepath[len(r.conf.Filepath)-1]) != string(os.PathSeparator) {
		// If not, append one
		r.conf.Filepath += string(os.PathSeparator)
	}

	// Construct filename using the current date and filepath
	filename := fmt.Sprintf("%s%s.db", r.conf.Filepath, time.Now().Format(time.RFC3339))

	fmt.Println(filename)
	// Create the database file
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	r.db = db

	p := data.Position{}
	_, err = db.Exec(p.SQLCreateQuery())
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func (r *Recorder) Start() error {
	// (re)set the stop channel
	r.stopCh <- false

	go r.run()
	return nil
}

func (r *Recorder) Stop() error {
	log.Println("Stop() called")
	r.stopCh <- true
	return nil
}

func (r *Recorder) run() {
	ticker := time.NewTicker(r.conf.Period)
	log.Printf("recorder period set to %v", r.conf.Period)
	defer ticker.Stop()

	log.Print("recorder started")
	for range ticker.C {
		select {
		case p := <-r.positionCh:
			log.Printf("recording position: %v", p)
			_, err := r.db.Exec(p.SQLInsertQuery())
			if err != nil {
				log.Printf("failed to insert position: %v", err)
				r.errCh <- fmt.Errorf("failed to insert position: %w", err)
			}
		case r := <-r.stopCh:
			if r {
				log.Print("recorder stopping")
				return
			}
		default:
			log.Print("no position to record")
		}
	}
}
