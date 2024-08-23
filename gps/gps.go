package gps

import (
	"fmt"
	"sync"

	"github.com/go-viper/mapstructure/v2"
	"github.com/kilianp07/AthleteIQBox/gps/nmea"
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

	if err := mapstructure.Decode(conf.Conf, r.Conf()); err != nil {
		return nil, fmt.Errorf("failed to decode configuration: %w", err)
	}

	if err := r.Configure(); err != nil {
		return nil, err
	}

	return r, nil
}
