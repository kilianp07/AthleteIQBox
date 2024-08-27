package services

import "github.com/paypal/gatt"

type Service interface {
	Configure() (*gatt.Service, error)
	Update(values any) error
	GetServiceUUID() string
	Conf() any
}
