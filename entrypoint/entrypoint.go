package entrypoint

import (
	"context"
	"log"

	"github.com/kilianp07/AthleteIQBox/buttons"
	"github.com/kilianp07/AthleteIQBox/gps"
	"github.com/kilianp07/AthleteIQBox/transmitter"
	utils "github.com/kilianp07/AthleteIQBox/utils/logger"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type configuration struct {
	Gps       gps.Configuration         `json:"gps"`
	Bluetooth transmitter.Conf          `json:"bluetooth"`
	Logger    utils.LoggerConfiguration `json:"logger"`
	Buttons   buttons.Configuration     `json:"buttons"`
}

func Start(confFile string) {
	var (
		err    error
		k      = koanf.New(".")
		parser = json.Parser()
		conf   = configuration{}
		logger *utils.Logger
	)

	// Load JSON config.
	if err := k.Load(file.Provider(confFile), parser); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.UnmarshalWithConf("", &conf, koanf.UnmarshalConf{Tag: "json"}); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	// Configure logger.
	logger = utils.NewLogger(conf.Logger)

	// Configure buttons
	buttonsManager, err := buttons.NewManager(conf.Buttons)
	if err != nil {
		logger.Errorf("error creating buttons manager: %v", err)
		return
	}

	if err = gps.Configure(conf.Gps); err != nil {
		logger.Errorf("error configuring gps reader: %v", err)
		return
	}

	t, err := transmitter.New(conf.Bluetooth)
	if err != nil {
		logger.Errorf("error creating transmitter: %v", err)
		return
	}

	if err := t.Start(context.Background()); err != nil {
		logger.Errorf("Error starting transmitter: %v", err)
		return
	}

	// Block and wait for button presses indefinitely
	logger.Debugf("Entering button press loop")
	for r := range buttonsManager.Run() {
		logger.Infof("Received button press event: %v", r)

		if r {
			logger.Debugf("Handling start button press")
			logger.Infof("Transmitter started successfully")

			if err := gps.Start(); err != nil {
				logger.Errorf("Error starting GPS reader: %v", err)
				return
			}
			logger.Infof("GPS reader started successfully")
		} else {
			logger.Infof("Handling stop button press")
			if err := gps.Stop(); err != nil {
				logger.Errorf("Error stopping GPS reader: %v", err)
				return
			}
			logger.Infof("GPS reader stopped successfully")
		}
	}

}
