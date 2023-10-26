package models

import (
	"simbir-go-api/constants"
	"strings"
	"time"
)

type Rent struct {
	ID        int32
	Account   Account
	Transport Transport
	TimeStart time.Time
	TimeEnd   *time.Time
	PriceUnit float64
	PriceType string
	IsDeleted bool
}

type GetRentTransportInput struct {
	Latitude      *float64 `form:"lat" binding:"required,min=-180,max=180"`
	Longitude     *float64 `form:"long" binding:"required,min=-180,max=180"`
	Radius        *float64 `form:"radius" binding:"required,min=-180,max=180"`
	TransportType string   `form:"type" binding:"required"`
}

func (i *GetRentTransportInput) Validate() error {
	err := constants.CheckTransportTypeWithAll(i.TransportType)
	if err != nil {
		return err
	}
	i.TransportType = strings.ToUpper(i.TransportType)
	return nil
}

type GetRentTransportOutput struct {
	Transports []GetTransportOutput `json:"transports"`
}

type GetRentOutput struct {
	ID        int32      `json:"ID"`
	Account   string     `json:"account"`
	Transport string     `json:"transport"`
	TimeStart time.Time  `json:"timeStart"`
	TimeEnd   *time.Time `json:"timeEnd"`
	PriceUnit float64    `json:"priceUnit"`
	PriceType string     `json:"priceType"`
}

type GetRentsOutput struct {
	Rents []GetRentOutput `json:"rents"`
}
