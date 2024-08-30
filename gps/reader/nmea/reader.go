package nmea

import (
	"bufio"
	"fmt"
	"time"

	"github.com/adrianmo/go-nmea"
	"github.com/kilianp07/AthleteIQBox/data"
	utils "github.com/kilianp07/AthleteIQBox/utils/logger"
	"github.com/tarm/serial"
)

type Reader struct {
	conf   Configuration
	period time.Duration

	runCh      chan bool
	errCh      chan error
	positionCh chan data.Position
	logger     *utils.Logger
}

func New() *Reader {
	return &Reader{
		conf:   Configuration{},
		logger: utils.GetLogger("NMEA Reader"),
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
	// Initialize the position channel
	r.positionCh = make(chan data.Position, 1)

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

func (r *Reader) Position() chan data.Position {
	return r.positionCh
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

	latLonFilled := false
	altitudeFilled := false
	speedFilled := false

	actual := data.Position{}
	for {
		select {
		case running := <-r.runCh:
			if !running {
				r.logger.Infof("Received an order to stop reading from GPS.")
				return
			}

		default:
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					r.logger.Errorf("Scanner error: %v\n", err)
				}
				return
			}

			line := scanner.Text()

			sentence, err := nmea.Parse(line)
			if err != nil {
				r.logger.Errorf("Error parsing NMEA sentence: %v\n", err)
				r.errCh <- fmt.Errorf("error parsing NMEA sentence: %w", err)
				continue
			}

			switch sentence.DataType() {
			case nmea.TypeGGA:
				gga := sentence.(nmea.GGA)
				actual.Altitude_M = gga.Altitude
				altitudeFilled = true

			case nmea.TypeRMC:
				rmc := sentence.(nmea.RMC)
				actual.Latitude = rmc.Latitude
				actual.Longitude = rmc.Longitude
				actual.Course = rmc.Course
				latLonFilled = true

			case nmea.TypeVTG:
				vtg := sentence.(nmea.VTG)
				actual.Speed_kMh = vtg.GroundSpeedKPH
				speedFilled = true
			}

			if latLonFilled && altitudeFilled && speedFilled {
				actual.Timestamp = time.Now().Unix()
				select {
				case r.positionCh <- actual.Copy():
					r.logger.Debugf("Sent position: %v\n", actual)
				default:
					r.logger.Errorf("Position channel is full, dropping data")
				}

				latLonFilled = false
				altitudeFilled = false
				speedFilled = false
				actual = data.Position{}
			}
		}
	}
}
