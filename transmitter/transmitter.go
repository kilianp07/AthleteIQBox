package transmitter

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/kilianp07/AthleteIQBox/transmitter/services"
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
}

// New creates a new Transmitter instance with the given configuration.
// It returns a pointer to the Transmitter and an error, if any.
func New(configuration *Conf) (*Transmitter, error) {
	// Create the transmitter
	t := &Transmitter{
		conf:     configuration,
		services: make(map[string]services.Service),
	}

	return t, nil
}

// Start starts the transmitter by configuring it and initializing the device.
// It takes a sync.WaitGroup and a channel for error communication as parameters.
// It returns no values.
func (t *Transmitter) Start(wg *sync.WaitGroup, errchan chan error, successChan chan bool, ctx context.Context) {
	defer wg.Done()

	err := t.configure()
	if err != nil {
		errchan <- err
	}

	if err = t.device.Init(t.onStateChanged); err != nil {
		errchan <- err
	}

	successChan <- true
	select {
	case <-ctx.Done():
		return
	default:

	}

}

// configure configures the transmitter by creating the device, creating the services,
// configuring the services, and collecting the services UUIDs.
// It returns an error if any of the configuration steps fail.
func (t *Transmitter) configure() error {
	// Create the device
	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
		return err
	}
	t.device = d

	// Create the services
	err = t.creator()
	if err != nil {
		return err
	}

	// Configure the services
	for _, service := range t.services {
		err = service.Configure(t.device)
		if err != nil {
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
	log.Printf("State: %s", s)
	switch s {
	case gatt.StatePoweredOn:
		if err := d.AdvertiseNameAndServices(deviceName, t.servicesUUIDs); err != nil {
			log.Fatalf("Failed to advertise name and services, err: %s", err)
		}
	default:
		log.Println("State:", s)
	}
}

// creator is a method of the Transmitter struct that initializes and configures the services based on the provided configuration.
// It iterates over the services configuration and creates instances of each service using the factory function.
// The created services are stored in the services map of the Transmitter struct.
// Returns an error if there is any issue with creating the services.
func (t *Transmitter) creator() error {
	fmt.Println(t.conf)
	for id, conf := range t.conf.ServicesConf {
		service, err := factory(id, conf)
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
	// Check if the service exists
	if _, ok := t.services[service]; !ok {
		return fmt.Errorf("service %s does not exist", service)
	}

	// Update the service
	return t.services[service].Update(data)
}
