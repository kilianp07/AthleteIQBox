package recorder

import (
	"fmt"
	"reflect"

	"github.com/kilianp07/AthleteIQBox/data"
	"github.com/kilianp07/AthleteIQBox/gps/recorder/sqlite"
	"github.com/kilianp07/AthleteIQBox/utils"
)

var (
	recorders = make(map[string]func() Recorder)
)

func init() {
	recorders["sqlite"] = func() Recorder {
		return sqlite.New()
	}
}

type Recorder interface {
	Configure(chan data.Position) error
	RuntimeErr() <-chan error
	Conf() any
	Start() error
	Stop() error
}

func New(conf Configuration, positionChan chan data.Position) (Recorder, error) {

	constructor, ok := recorders[conf.ID]
	if !ok {
		return nil, fmt.Errorf("invalid recorder: %s", conf.ID)
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

	if err := r.Configure(positionChan); err != nil {
		return nil, err
	}

	return r, nil
}
