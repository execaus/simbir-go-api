package models

import (
	"simbir-go-api/constants"
	"strings"
)

type Transport struct {
	OwnerID       string
	CanBeRented   bool
	TransportType string
	Model         string
	Color         string
	Identifier    string
	Description   *string
	Latitude      float64
	Longitude     float64
	MinutePrice   *float64
	DayPrice      *float64
	IsDeleted     bool
}

type CreateTransportInput struct {
	CanBeRented   bool     `json:"canBeRented" binding:"required"`
	TransportType string   `json:"transportType" binding:"required"`
	Model         string   `json:"model" binding:"required"`
	Color         string   `json:"color" binding:"required"`
	Identifier    string   `json:"identifier" binding:"required"`
	Description   *string  `json:"description"`
	Latitude      *float64 `json:"latitude" binding:"required,min=-180,max=180"`
	Longitude     *float64 `json:"longitude" binding:"required,min=-180,max=180"`
	MinutePrice   *float64 `json:"minutePrice" binding:"min=1"`
	DayPrice      *float64 `json:"dayPrice" binding:"min=1"`
}

func (t *CreateTransportInput) Validate() error {
	t.TransportType = strings.ToUpper(t.TransportType)
	return constants.CheckTransportType(t.TransportType)
}

type CreateTransportOutput struct {
	Transport *GetTransportOutput `json:"transport"`
}

type GetTransportOutput struct {
	OwnerID       string   `json:"ownerId"`
	CanBeRented   bool     `json:"canBeRented"`
	TransportType string   `json:"transportType"`
	Model         string   `json:"model"`
	Color         string   `json:"color"`
	Identifier    string   `json:"identifier"`
	Description   *string  `json:"description"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	MinutePrice   *float64 `json:"minutePrice"`
	DayPrice      *float64 `json:"dayPrice"`
}

type UpdateTransportInput struct {
	CanBeRented bool     `json:"canBeRented" binding:"required"`
	Model       string   `json:"model" binding:"required"`
	Color       string   `json:"color" binding:"required"`
	Identifier  string   `json:"identifier" binding:"required"`
	Description *string  `json:"description"`
	Latitude    *float64 `json:"latitude" binding:"required,min=-180,max=180"`
	Longitude   *float64 `json:"longitude" binding:"required,min=-180,max=180"`
	MinutePrice *float64 `json:"minutePrice" binding:"min=1"`
	DayPrice    *float64 `json:"dayPrice" binding:"min=1"`
}

type UpdateTransportOutput struct {
	Transport *GetTransportOutput `json:"transport"`
}

type AdminGetTransportOutput struct {
	Transport *GetTransportOutput `json:"transport"`
	IsDeleted bool                `json:"isDeleted"`
}

type AdminGetTransportsInput struct {
	Start         int32  `form:"start" binding:"min=0"`
	Count         int32  `form:"count" binding:"min=1"`
	TransportType string `form:"transportType" binding:"required"`
}

func (i *AdminGetTransportsInput) Validate() error {
	if err := constants.CheckTransportTypeWithAll(i.TransportType); err != nil {
		return err
	}

	i.TransportType = strings.ToUpper(i.TransportType)
	return nil
}

type AdminGetTransportsOutput struct {
	Transports []*AdminGetTransportOutput `json:"transports"`
}
