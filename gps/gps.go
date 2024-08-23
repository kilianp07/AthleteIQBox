package gps

import (
	"fmt"

	"github.com/go-viper/mapstructure/v2"
	"github.com/kilianp07/AthleteIQBox/gps/nmea"
)

var (
	readers = make(map[string]func() Reader)
)

func init() {
	readers["nmea"] = func() Reader {
		return nmea.New()
	}
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
	var (
		constructor func() Reader
		r           Reader
		ok          bool
		err         error
	)
	if constructor, ok = readers[conf.ID]; !ok {
		return nil, fmt.Errorf("invalid reader: %s", conf.ID)
	}

	r = constructor()

	if err = mapstructure.Decode(conf.Conf, r.Conf()); err != nil {
		return nil, fmt.Errorf("failed to decode configuration: %w", err)
	}

	if err = r.Configure(); err != nil {
		return nil, err
	}

	return r, nil
}
