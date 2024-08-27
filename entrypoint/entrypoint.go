package entrypoint

import (
	"context"
	"log"

	"github.com/kilianp07/AthleteIQBox/gps/reader"
	"github.com/kilianp07/AthleteIQBox/transmitter"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type configuration struct {
	Gps       reader.Configuration `json:"gps"`
	Bluetooth transmitter.Conf     `json:"bluetooth"`
}

func Start(confFile string) {
	var (
		err    error
		k      = koanf.New(".")
		parser = json.Parser()
		conf   = configuration{}
		r      reader.Reader
	)

	// Load JSON config.
	if err := k.Load(file.Provider(confFile), parser); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.UnmarshalWithConf("", &conf, koanf.UnmarshalConf{Tag: "json"}); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	if r, err = reader.New(conf.Gps); err != nil {
		log.Fatalf("error creating gps reader: %v", err)
	}

	err = r.Start()

	t, err := transmitter.New(conf.Bluetooth)
	if err != nil {
		log.Fatalf("error creating transmitter: %v", err)
	}

	if err = t.Start(context.Background()); err != nil {
		log.Fatalf("error starting transmitter: %v", err)
	}
	if err != nil {
		log.Fatalf("error starting gps reader: %v", err)
		return
	}

	// Block forever.
	select {}
}
