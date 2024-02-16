package services

import "github.com/paypal/gatt"

type Service interface {
	Configure(d gatt.Device) error
	Update(values any) error
	GetServiceUUID() string
}
