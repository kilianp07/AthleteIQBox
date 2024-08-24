package entrypoint

import (
	"log"

	"github.com/kilianp07/AthleteIQBox/gps"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type configuration struct {
	Gps gps.Configuration `json:"gps"`
}

func Start(confFile string) {
	var (
		err    error
		k      = koanf.New(".")
		parser = json.Parser()
		conf   = configuration{}
		r      gps.Reader
	)

	// Load JSON config.
	if err := k.Load(file.Provider(confFile), parser); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.UnmarshalWithConf("", &conf, koanf.UnmarshalConf{Tag: "json"}); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	if r, err = gps.New(conf.Gps); err != nil {
		log.Fatalf("error creating gps reader: %v", err)
	}

	err = r.Start()
	if err != nil {
		log.Fatalf("error starting gps reader: %v", err)
		return
	}

	// Block forever.
	select {}
}
