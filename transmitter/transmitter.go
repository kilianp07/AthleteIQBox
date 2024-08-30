package transmitter

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
	"github.com/kilianp07/AthleteIQBox/transmitter/services"
	"github.com/kilianp07/AthleteIQBox/utils"
	logger "github.com/kilianp07/AthleteIQBox/utils/logger"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

const (
	deviceName = "AthleteIQBox"
)

type Transmitter struct {
	services      map[string]services.Service
	conf          *Conf
	device        gatt.Device
	servicesUUIDs []gatt.UUID
	logger        *logger.Logger
}

// New creates a new Transmitter instance with the given configuration.
// It returns a pointer to the Transmitter and an error, if any.
func New(conf any) (*Transmitter, error) {

	configuration := Conf{}
	err := mapstructure.Decode(conf, &configuration)
	if err != nil {
		return nil, err
	}

	// Create the transmitter
	t := &Transmitter{
		conf:     &configuration,
		services: make(map[string]services.Service),
		logger:   logger.GetLogger("Transmitter"),
	}

	return t, nil
}

// Start starts the transmitter by configuring it and initializing the device.
// It takes a sync.WaitGroup and a channel for error communication as parameters.
// It returns no values.
func (t *Transmitter) Start(ctx context.Context) error {
	var err error

	if err = t.configure(); err != nil {
		return err
	}

	if err = t.device.Init(t.onStateChanged); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return nil
	default:
	}

	return nil
}

// configure configures the transmitter by creating the device, creating the services,
// configuring the services, and collecting the services UUIDs.
// It returns an error if any of the configuration steps fail.
func (t *Transmitter) configure() error {

	// Create the device
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		t.logger.Fatalf("Failed to open device, err: %s", err)
		return err
	}
	t.device = d

	// Create the services
	err = t.creator()
	if err != nil {
		return err
	}

	// Configure the services
	for id, service := range t.services {
		decoder, err := utils.NewDecoder(service.Conf())
		if err != nil {
			return err
		}

		err = decoder.Decode(t.conf.ServicesConf[id])
		if err != nil {
			return err
		}

		s, err := service.Configure()
		if err != nil {
			return err
		}
		if err := t.device.AddService(s); err != nil {
			t.logger.Errorf("Failed to add service, err: %s", err)
			return err
		}
	}

	// Collect the services UUIDs
	for _, service := range t.services {
		t.servicesUUIDs = append(t.servicesUUIDs, gatt.MustParseUUID(service.GetServiceUUID()))
	}
	return nil
}

// onStateChanged is a callback function that is called when the state of the gatt.Device changes.
// It logs the current state and performs actions based on the state.
// If the state is gatt.StatePoweredOn, it advertises the device name and services.
// Otherwise, it logs the state.
func (t *Transmitter) onStateChanged(d gatt.Device, s gatt.State) {
	t.logger.Debugf("State: %s", s)
	switch s {
	case gatt.StatePoweredOn:
		if err := d.AdvertiseNameAndServices(deviceName, t.servicesUUIDs); err != nil {
			t.logger.Fatalf("Failed to advertise name and services, err: %s", err)
		}
	default:
		t.logger.Debugf("State:", s)
	}
}

// creator is a method of the Transmitter struct that initializes and configures the services based on the provided configuration.
// It iterates over the services configuration and creates instances of each service using the factory function.
// The created services are stored in the services map of the Transmitter struct.
// Returns an error if there is any issue with creating the services.
func (t *Transmitter) creator() error {
	for id := range t.conf.ServicesConf {
		service, err := factory(id)
		if err != nil {
			return err
		}
		t.services[id] = service
	}
	return nil
}

// Update updates the specified service with the provided data.
// It returns an error if the service does not exist or if there is an error updating the service.
func (t *Transmitter) Update(service string, data any) error {
	// Update the service
	return t.services[service].Update(data)
}
