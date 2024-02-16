package transmitter

import (
	"fmt"

	"github.com/kilianp07/AthleteIQBox/transmitter/services"
)

type ServiceConstructor func(conf any) (*services.Position, error)

var constructors = map[string]ServiceConstructor{
	"position": services.NewPosition,
}

func factory(id string, conf any) (services.Service, error) {
	constructor, ok := constructors[id]
	if !ok {
		return nil, fmt.Errorf("invalid service type %s", id)
	}
	return constructor(conf)
}
