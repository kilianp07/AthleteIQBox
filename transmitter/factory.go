package transmitter

import (
	"fmt"

	"github.com/kilianp07/AthleteIQBox/transmitter/services"
	"github.com/kilianp07/AthleteIQBox/transmitter/services/wifi"
)

type ServiceConstructor func() services.Service

var constructors = map[string]ServiceConstructor{
	"wifi": wifi.NewWiFi,
}

func factory(id string) (services.Service, error) {
	constructor, ok := constructors[id]
	if !ok {
		return nil, fmt.Errorf("invalid service type %s", id)
	}
	return constructor(), nil
}
