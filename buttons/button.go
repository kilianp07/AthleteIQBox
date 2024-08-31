package buttons

import (
	"fmt"

	utils "github.com/kilianp07/AthleteIQBox/utils/logger"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
)

type Button interface {
	Trigger() chan bool
	Start()
}

type Reader struct {
	logger      *utils.Logger
	conf        ButtonConfiguration
	triggerChan chan bool
	pin         gpio.PinIO
}

// NewReader creates a new Button reader with the given configuration.
// It initializes the reader with the provided ButtonConfiguration and sets up the necessary GPIO pin.
// The function returns a Button interface and an error if any occurred during initialization.
func NewReader(conf ButtonConfiguration) (*Reader, error) {
	r := Reader{
		conf:        conf,
		logger:      utils.GetLogger(fmt.Sprintf("Button %s reader", conf.name)),
		triggerChan: make(chan bool, 1),
	}

	r.logger.Infof("Initializing button %s reader", conf.name)

	p := gpioreg.ByName(conf.Gpio)
	if p == nil {
		return nil, fmt.Errorf("failed to find GPIO %s", conf.Gpio)
	}

	if err := p.In(gpio.PullUp, gpio.RisingEdge); err != nil {
		r.logger.Errorf("failed to set pin %s as input: %v", conf.Gpio, err)
		//return nil, fmt.Errorf("failed to set pin %s as input: %w", conf.Gpio, err)
	}

	r.pin = p

	return &r, nil
}

// Trigger returns a channel that will be triggered when the button is pressed.
func (r *Reader) Trigger() chan bool {
	return r.triggerChan
}

func (r *Reader) Start() {
	r.logger.Debugf("Starting button %s reader", r.conf.name)
	go r.run()
}

func (r *Reader) run() {
	r.logger.Debugf("Button %s is starting to watch", r.conf.name)
	var lastState bool
	for {
		// Wait for an edge (change in signal)
		r.pin.WaitForEdge(-1)

		// Check if pin is low (button pressed)
		if r.pin.Read() == gpio.Low && !lastState {
			// Button was pressed, send signal
			lastState = true
			r.logger.Infof("Button %s pressed", r.conf.name)
			select {
			case r.triggerChan <- true:
			default:
				r.logger.Warnf("Button %s trigger channel is full", r.conf.name)
			}
		} else if r.pin.Read() != gpio.Low && lastState {
			// Button was released, reset the state
			lastState = false
		}
	}
}
