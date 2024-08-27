package reader

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/kilianp07/AthleteIQBox/gps/reader/nmea"
	"github.com/kilianp07/AthleteIQBox/utils"
)

var (
	readers   = make(map[string]func() Reader)
	readersMu sync.RWMutex
)

func init() {
	// Register the "nmea" reader
	readersMu.Lock()
	readers["nmea"] = func() Reader {
		return nmea.New()
	}
	readersMu.Unlock()
}

type Reader interface {
	Configure() error
	Start() error
	Stop() error
	RuntimeErr() chan error
	Conf() any
}

type Configuration struct {
	ID   string `json:"id"`
	Conf any    `json:"conf"`
}

func New(conf Configuration) (Reader, error) {
	readersMu.RLock()
	constructor, ok := readers[conf.ID]
	readersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("invalid reader: %s", conf.ID)
	}

	r := constructor()

	// Check if r.Conf returns a pointer
	if reflect.ValueOf(r.Conf()).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("Conf() method of reader %s did not return a pointer", conf.ID)
	}

	decoder, err := utils.NewDecoder(r.Conf())
	if err != nil {
		return nil, fmt.Errorf("failed to create decoder: %w", err)
	}

	if err := decoder.Decode(conf.Conf); err != nil {
		return nil, fmt.Errorf("failed to decode configuration: %w", err)
	}

	if err := r.Configure(); err != nil {
		return nil, err
	}

	return r, nil
}
