package models

import (
	"simbir-go-api/constants"
	"strings"
	"time"
)

type Rent struct {
	ID         int32
	Account    int32
	Transport  int32
	TimeStart  time.Time
	TimeEnd    *time.Time
	PriceUnit  float64
	PriceType  string
	FinalPrice *float64
	IsDeleted  bool
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
	ID         int32      `json:"ID"`
	Account    int32      `json:"account"`
	Transport  int32      `json:"transport"`
	TimeStart  time.Time  `json:"timeStart"`
	TimeEnd    *time.Time `json:"timeEnd"`
	PriceUnit  float64    `json:"priceUnit"`
	PriceType  string     `json:"priceType"`
	FinalPrice *float64   `json:"finalPrice"`
}

type GetRentsOutput struct {
	Rents []GetRentOutput `json:"rents"`
}

type GetRentsMyHistoryOutput struct {
	Rents []GetRentOutput `json:"rents"`
}

type GetRentTransportHistoryOutput struct {
	Rents []GetRentOutput `json:"rents"`
}

type GetRentTransportNewInput struct {
	RentType string `json:"rentType"`
}

func (i *GetRentTransportNewInput) Validate() error {
	err := constants.CheckRentType(i.RentType)
	if err != nil {
		return err
	}
	i.RentType = strings.ToUpper(i.RentType)
	return nil

}

type GetRentTransportNewOutput struct {
	Rent GetRentOutput `json:"rent"`
}

type EndRentInput struct {
	Latitude  *float64 `json:"lat" binding:"required,min=-180,max=180"`
	Longitude *float64 `form:"long" binding:"required,min=-180,max=180"`
}

type EndRentOutput struct {
	Rent GetRentOutput `json:"rent"`
}

type GetAdminRentOutput struct {
	Rent      GetRentOutput `json:"rent"`
	IsDeleted bool          `json:"isDeleted"`
}

type GetAdminUserHistoryOutput struct {
	Rents []GetAdminRentOutput `json:"rents"`
}

type GetAdminTransportHistoryOutput struct {
	Rents []GetAdminRentOutput `json:"rents"`
}

type CreateAdminRentInput struct {
}
