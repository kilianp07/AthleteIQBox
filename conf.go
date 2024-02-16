package main

import (
	"github.com/kilianp07/AthleteIQBox/transmitter"
)

type Conf struct {
	Server transmitter.Conf `json:",inline"`
}
