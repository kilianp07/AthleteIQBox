package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/kilianp07/AthleteIQBox/data"
	utils "github.com/kilianp07/AthleteIQBox/utils/logger"

	// Import the sqlite3 package to register the database driver
	_ "github.com/mattn/go-sqlite3"
)

type Recorder struct {
	conf       configuration
	positionCh chan data.Position
	stopCh     chan bool
	errCh      chan error
	db         *sql.DB
	logger     *utils.Logger
}

func New() *Recorder {
	return &Recorder{
		conf:   configuration{},
		stopCh: make(chan bool, 1),
		logger: utils.GetLogger("SQLITE Recorder"),
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
	r.logger.Infof("Creating database as %s", filename)

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
	r.logger.Infof("Recorder stopping")
	r.stopCh <- true
	return nil
}

func (r *Recorder) run() {
	ticker := time.NewTicker(r.conf.Period)
	defer ticker.Stop()

	r.logger.Infof("recorder started")
	for range ticker.C {
		select {
		case p := <-r.positionCh:
			r.logger.Debugf("recording position: %v", p)
			_, err := r.db.Exec(p.SQLInsertQuery())
			if err != nil {
				r.logger.Errorf("failed to insert position: %v", err)
				r.errCh <- fmt.Errorf("failed to insert position: %w", err)
			}
		case stop := <-r.stopCh:
			if stop {
				r.logger.Infof("recorder stopping")
				return
			}
		default:
			r.logger.Infof("no position to record")
		}
	}
}
