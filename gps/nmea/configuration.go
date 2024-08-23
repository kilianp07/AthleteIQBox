package nmea

import (
	"time"

	"github.com/tarm/serial"
)

type Configuration struct {
	serialConfig `json:"squash"`
	Period       string `json:"period"`
}

type serialConfig struct {
	Name        string        `json:"name"`
	Baud        int           `json:"baud"`
	ReadTimeout time.Duration `json:"read_timeout"`
	Size        byte          `json:"size"`
	Parity      byte          `json:"parity"`
	StopBits    byte          `json:"stop_bits"`
}

func (s *serialConfig) ToSerial() *serial.Config {
	return &serial.Config{
		Name:        s.Name,
		Baud:        s.Baud,
		ReadTimeout: s.ReadTimeout,
		Size:        s.Size,
		Parity:      serial.Parity(s.Parity),
		StopBits:    serial.StopBits(s.StopBits),
	}
}
