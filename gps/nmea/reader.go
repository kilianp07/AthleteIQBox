package nmea

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/adrianmo/go-nmea"
	"github.com/tarm/serial"
)

type Point struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Altitude   float64 `json:"altitude_m"`
	Satellites int64   `json:"satellites"`
}

type Reader struct {
	conf   Configuration
	period time.Duration

	runCh chan bool
	errCh chan error
}

func New() *Reader {
	return &Reader{
		conf: Configuration{},
	}
}

func (r *Reader) Conf() any {
	return &r.conf
}
func (r *Reader) Configure() error {

	// Parse the period
	period, err := time.ParseDuration(r.conf.Period)
	if err != nil {
		return fmt.Errorf("failed to parse period: %w", err)
	}
	r.period = period

	// Initialize the run channel
	r.runCh = make(chan bool, 1)
	// Initialize the error channel
	r.errCh = make(chan error, 1)

	// Try to open the serial port
	s, err := serial.OpenPort(r.conf.serialConfig.ToSerial())
	if err != nil {
		return fmt.Errorf("failed to open serial port: %v", err)
	}
	defer s.Close()

	return nil
}

func (r *Reader) Start() error {
	var (
		s   *serial.Port
		err error
	)

	if s, err = serial.OpenPort(r.conf.serialConfig.ToSerial()); err != nil {
		return fmt.Errorf("failed to open serial port: %v", err)
	}

	r.runCh <- true

	go r.run(s)

	return nil
}

func (r *Reader) Stop() error {
	r.runCh <- false

	return nil
}

func (r *Reader) RuntimeErr() chan error {
	return r.errCh
}

func (r *Reader) run(s *serial.Port) {

	ticker := time.NewTicker(r.period)
	defer ticker.Stop()

	scanner := bufio.NewScanner(s)
	defer s.Close()
	for {
		select {
		case <-ticker.C:
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					log.Printf("Scanner error: %v\n", err)
				}
				return
			}

			line := scanner.Text()

			// Parse the NMEA sentence
			sentence, err := nmea.Parse(line)
			if err != nil {
				log.Printf("Error parsing NMEA sentence: %v\n", err)
				continue
			}

			switch sentence.DataType() {
			case nmea.TypeGGA:
				gga := sentence.(nmea.GGA)
				fmt.Printf("GGA: Latitude: %f, Longitude: %f, Altitude: %f meters, Satellites: %d\n",
					gga.Latitude, gga.Longitude, gga.Altitude, gga.NumSatellites)

			case nmea.TypeRMC:
				rmc := sentence.(nmea.RMC)
				fmt.Printf("RMC: Latitude: %f, Longitude: %f, Speed: %f knots, Date: %s\n",
					rmc.Latitude, rmc.Longitude, rmc.Speed, rmc.Date.String())

			case nmea.TypeGLL:
				gll := sentence.(nmea.GLL)
				fmt.Printf("GLL: Latitude: %f, Longitude: %f\n", gll.Latitude, gll.Longitude)

			case nmea.TypeGSA:
				gsa := sentence.(nmea.GSA)
				fmt.Printf("GSA: PDOP: %f, HDOP: %f, VDOP: %f\n", gsa.PDOP, gsa.HDOP, gsa.VDOP)

			case nmea.TypeGSV:
				gsv := sentence.(nmea.GSV)
				fmt.Printf("GSV: Number of Satellites in View: %d\n", gsv.NumberSVsInView)

			case nmea.TypeVTG:
				vtg := sentence.(nmea.VTG)
				fmt.Printf("VTG: True Track: %f degrees, Speed: %f knots\n", vtg.TrueTrack, vtg.GroundSpeedKnots)
			}

		case running := <-r.runCh:
			if !running {
				log.Println("Received an order to stop reading from gps.")
				return
			}
		}
	}

}
