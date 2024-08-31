package buttons

import (
	"fmt"

	utils "github.com/kilianp07/AthleteIQBox/utils/logger"
	"periph.io/x/host/v3"
)

const (
	StartButton = "start"
	StopButton  = "stop"
)

var (
	requiredButtons   = []string{StartButton, StopButton}
	ErrButtonNotFound = "button %s not found in configuration"
)

type Manager struct {
	buttons map[string]Button
	Conf    Configuration
	logger  *utils.Logger
	runChan chan bool
}

func NewManager(conf Configuration) (*Manager, error) {
	m := &Manager{
		Conf:    conf,
		logger:  utils.GetLogger("Buttons Manager"),
		buttons: make(map[string]Button),
		runChan: make(chan bool, 1),
	}

	m.logger.Infof("Initializing buttons manager")

	// Check if all required buttons are present in the configuration
	for _, button := range requiredButtons {
		if _, ok := conf.Switches[button]; !ok {
			return nil, fmt.Errorf(ErrButtonNotFound, button)
		}
	}

	// Initialize the drivers
	_, err := host.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to load drivers: %w", err)
	}

	for name, conf := range conf.Switches {
		conf.name = name
		b, err := NewReader(conf)
		if err != nil {
			//return nil, fmt.Errorf("failed to initialize button %s: %w", name, err)
			m.logger.Errorf("failed to initialize button %s: %v", name, err)
		}

		m.logger.Debugf("Button %s initialized", name)
		m.buttons[name] = b
	}

	for name, b := range m.buttons {
		m.logger.Infof("Starting button %s", name)
		b.Start()
	}

	go m.run()

	return m, nil
}

func (m *Manager) Run() chan bool {
	return m.runChan
}

func (m *Manager) run() {
	var isRunning bool

	for {
		select {
		case run := <-m.buttons[StartButton].Trigger():
			if isRunning {
				m.logger.Infof("Start button already pressed, ignoring")
			} else {
				isRunning = true
				m.logger.Debugf("Start button pressed: %v", run)
				m.runChan <- true // Send the button state (true for pressed, false for released)
			}
		case stop := <-m.buttons[StopButton].Trigger():
			if !isRunning {
				m.logger.Infof("Stop button already pressed, ignoring")
			} else {
				isRunning = false
				m.logger.Debugf("Stop button pressed: %v", stop)
				m.runChan <- false // Send the button state (true for pressed, false for released)
			}
		}
	}
}
