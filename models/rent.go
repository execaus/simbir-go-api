package models

import (
	"simbir-go-api/constants"
	"strings"
)

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
