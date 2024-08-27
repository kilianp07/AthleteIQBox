package gps

import (
	"fmt"

	"github.com/kilianp07/AthleteIQBox/gps/reader"
	"github.com/kilianp07/AthleteIQBox/gps/recorder"
	"github.com/kilianp07/AthleteIQBox/utils"
)

var (
	driver reader.Reader
	saver  recorder.Recorder
)

func Configure(configuration any) error {

	var conf Configuration
	decoder, err := utils.NewDecoder(&conf)
	if err != nil {
		return fmt.Errorf("error creating decoder: %v", err)
	}

	err = decoder.Decode(configuration)
	if err != nil {
		return fmt.Errorf("error decoding configuration: %v", err)
	}

	if driver, err = reader.New(conf.Reader); err != nil {
		return fmt.Errorf("error creating reader: %v", err)
	}

	if saver, err = recorder.New(conf.Recorder, driver.Position()); err != nil {
		return fmt.Errorf("error creating recorder: %v", err)
	}

	return err
}

func Start() error {
	if err := driver.Start(); err != nil {
		return fmt.Errorf("error starting gps reader: %v", err)
	}

	if err := saver.Start(); err != nil {
		return fmt.Errorf("error starting gps recorder: %v", err)
	}

	return nil
}

func Stop() error {
	if err := driver.Stop(); err != nil {
		return fmt.Errorf("error stopping gps reader: %v", err)
	}

	if err := saver.Stop(); err != nil {
		return fmt.Errorf("error stopping gps recorder: %v", err)
	}

	return nil
}
