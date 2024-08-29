package gps

import (
	"github.com/kilianp07/AthleteIQBox/gps/reader"
	"github.com/kilianp07/AthleteIQBox/gps/recorder"
)

type Configuration struct {
	Reader   reader.Configuration   `json:"reader"`
	Recorder recorder.Configuration `json:"recorder"`
}
