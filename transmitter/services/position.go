package services

import (
	"fmt"
	"log"

	"github.com/kilianp07/AthleteIQBox/transmitter/data"
	"github.com/kilianp07/AthleteIQBox/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/paypal/gatt"
)

type PositionConf struct {
	LatitudeUUID             string `json:"latitude_uuid"`
	LattitudeUserDescription string `json:"latitude_user_description"`

	LongitudeUUID            string `json:"longitude_uuid"`
	LongitudeUserDescription string `json:"longitude_user_description"`

	ServiceUUID         string `json:"service_uuid"`
	UserDescriptionUUID string `json:"user_description_uuid"`
}

type Position struct {
	conf PositionConf
	data data.Position

	// Bluetooth related attributes
	service *gatt.Service
	lonChar *gatt.Characteristic
	latChar *gatt.Characteristic
}

func NewPosition(conf any) (*Position, error) {
	var (
		positionConf PositionConf
	)

	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &positionConf,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)

	if decoder.Decode(conf) != nil {
		return nil, fmt.Errorf("invalid conf type: %T", conf)
	}
	return &Position{
		conf: positionConf,
	}, nil
}

func (p *Position) Configure(d gatt.Device) error {
	// Create a new service
	p.service = gatt.NewService(gatt.MustParseUUID(p.conf.ServiceUUID))

	// Add the characteristics to the service
	// Create longitude characteristic
	p.lonChar = p.service.AddCharacteristic(gatt.MustParseUUID(p.conf.LatitudeUUID))
	p.lonChar.HandleReadFunc(p.writeLat)
	p.lonChar.AddDescriptor(gatt.MustParseUUID(p.conf.UserDescriptionUUID)).SetValue(utils.StringToByte(p.conf.LattitudeUserDescription))

	// Create latitude characteristic
	p.latChar = p.service.AddCharacteristic(gatt.MustParseUUID(p.conf.LongitudeUUID))
	p.latChar.HandleReadFunc(p.writeLong)
	p.latChar.AddDescriptor(gatt.MustParseUUID(p.conf.UserDescriptionUUID)).SetValue(utils.StringToByte(p.conf.LongitudeUserDescription))

	// Register the service
	return d.AddService(p.service)
}

func (p *Position) Update(values any) error {
	var (
		position data.Position
		ok       bool
	)

	if position, ok = values.(data.Position); !ok {
		return fmt.Errorf("invalid data type: %T", values)
	}
	p.data = position
	return nil
}

func (p *Position) writeLat(rsp gatt.ResponseWriter, _ *gatt.ReadRequest) {
	value, err := utils.Float64ToByte(p.data.Latitude)
	if err != nil {
		rsp.SetStatus(gatt.StatusUnexpectedError)
		return
	}
	_, err = rsp.Write(value)
	if err != nil {
		log.Printf("Error writing latitude: %v", err)
	}
}

func (p *Position) writeLong(rsp gatt.ResponseWriter, _ *gatt.ReadRequest) {
	value, err := utils.Float64ToByte(p.data.Longitude)
	if err != nil {
		rsp.SetStatus(gatt.StatusUnexpectedError)
		return
	}
	_, err = rsp.Write(value)
	if err != nil {
		log.Printf("Error writing longitude: %v", err)
	}
}

func (p *Position) GetServiceUUID() string {
	return p.conf.ServiceUUID
}
