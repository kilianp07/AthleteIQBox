package entrypoint

import (
	"context"
	"log"

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

	if err = gps.Configure(conf.Gps); err != nil {
		logger.Errorf("error configuring gps reader: %v", err)
		return
	}

	if err := gps.Start(); err != nil {
		logger.Errorf("error starting gps reader: %v", err)
		return
	}

	t, err := transmitter.New(conf.Bluetooth)
	if err != nil {
		logger.Errorf("error creating transmitter: %v", err)
		return
	}

	if err = t.Start(context.Background()); err != nil {
		logger.Errorf("error starting transmitter: %v", err)
		return
	}

	// Block forever.
	select {}
}
