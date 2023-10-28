package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/constants"
	"simbir-go-api/models"
	"time"
)

// GetRentTransport
// @Summary      Get transport list
// @Description  Getting transport available for rent by parameters.
// @Tags         rent
// @Accept       json
// @Produce      json
// @Param        lat query number true "-"
// @Param        long query number true "-"
// @Param        radius query number true "-"
// @Param        type query string true "-"
// @Success      200  {object}  models.GetRentTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Rent/Transport [get]
func (h *Handler) GetRentTransport(c *gin.Context) {
	var input models.GetRentTransportInput

	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	transports, err := h.services.Transport.GetFromRadius(&models.Point{
		Longitude: *input.Longitude,
		Latitude:  *input.Latitude,
	}, *input.Radius, input.TransportType)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	var outputTransports []models.GetTransportOutput
	for _, transport := range transports {
		outputTransports = append(outputTransports, models.GetTransportOutput{
			OwnerID:       transport.OwnerID,
			CanBeRented:   transport.CanBeRented,
			TransportType: transport.TransportType,
			Model:         transport.Model,
			Color:         transport.Color,
			Identifier:    transport.Identifier,
			Description:   transport.Description,
			Latitude:      transport.Latitude,
			Longitude:     transport.Longitude,
			MinutePrice:   transport.MinutePrice,
			DayPrice:      transport.DayPrice,
		})
	}

	h.sendOKWithBody(c, &models.GetRentTransportOutput{Transports: outputTransports})
}

// GetRent
// @Summary      Rent information
// @Description  Return rent by id.
// @Tags         rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetRentOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Rent/{id} [get]
func (h *Handler) GetRent(c *gin.Context) {
	rentID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	username, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	isRenter, err := h.services.Rent.IsRenter(int32(rentID), username)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isRenter {
		h.sendAccessDenied(c, accountIsNotRenter)
		return
	}

	isExist, err := h.services.Rent.IsExist(int32(rentID))
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, rentIsNotExist)
		return
	}

	rent, err := h.services.Rent.Get(int32(rentID))
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if rent.Account.Username != username {
		h.sendInvalidRequest(c, accountIsNotTransportOwner)
		return
	}

	h.sendOKWithBody(c, &models.GetRentOutput{
		ID:        rent.ID,
		Account:   rent.Account.Username,
		Transport: rent.Transport.Identifier,
		TimeStart: rent.TimeStart,
		TimeEnd:   rent.TimeEnd,
		PriceUnit: rent.PriceUnit,
		PriceType: rent.PriceType,
	})
}

// GetRentMyHistory
// @Summary      Account rent history
// @Description  Returns the rental history of the current account.
// @Tags         rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetRentsOutput
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Rent/MyHistory [get]
func (h *Handler) GetRentMyHistory(c *gin.Context) {
	username, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	rents, err := h.services.Rent.GetListFromUsername(username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	var output models.GetRentsMyHistoryOutput
	for _, rent := range rents {
		output.Rents = append(output.Rents, models.GetRentOutput{
			ID:        rent.ID,
			Account:   rent.Account.Username,
			Transport: rent.Transport.Identifier,
			TimeStart: rent.TimeStart,
			TimeEnd:   rent.TimeEnd,
			PriceUnit: rent.PriceUnit,
			PriceType: rent.PriceType,
		})
	}

	h.sendOKWithBody(c, output)
}

// GetRentTransportHistory
// @Summary      Transport rent history
// @Description  Returns the transport rental history.
// @Tags         rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetRentTransportHistoryOutput
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Rent/TransportHistory/{id} [get]
func (h *Handler) GetRentTransportHistory(c *gin.Context) {
	transportID, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	username, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	isOwner, err := h.services.Transport.IsOwner(transportID, username)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isOwner {
		h.sendAccessDenied(c, accountIsNotTransportOwner)
		return
	}

	rents, err := h.services.Rent.GetListFromTransport(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	var output models.GetRentTransportHistoryOutput
	for _, rent := range rents {
		output.Rents = append(output.Rents, models.GetRentOutput{
			ID:        rent.ID,
			Account:   rent.Account.Username,
			Transport: rent.Transport.Identifier,
			TimeStart: rent.TimeStart,
			TimeEnd:   rent.TimeEnd,
			PriceUnit: rent.PriceUnit,
			PriceType: rent.PriceType,
		})
	}

	h.sendOKWithBody(c, output)
}

func (h *Handler) GetRentTransportNew(c *gin.Context) {
	var input models.GetRentTransportNewInput

	transportID, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err = c.ShouldBindQuery(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	username, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	isExist, err := h.services.Transport.IsExist(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, transportIsNotExist)
		return
	}

	isRemoved, err := h.services.Transport.IsRemoved(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isRemoved {
		h.sendResourceDeleted(c, transportIsDeleted)
		return
	}

	isAccessRent, err := h.services.Transport.IsAccessRent(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isAccessRent {
		h.sendInvalidRequest(c, transportIsNotRent)
		return
	}

	isOwner, err := h.services.Transport.IsOwner(transportID, username)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if isOwner {
		h.sendAccessDenied(c, accountIsTransportOwner)
		return
	}

	isRented, err := h.services.Rent.TransportIsRented(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if isRented {
		h.sendInvalidRequest(c, transportInRent)
		return
	}

	transport, err := h.services.Transport.Get(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	var priceUnit float64
	if input.RentType == constants.RentTypeMinutes {
		if transport.MinutePrice == nil {
			h.sendConditionNotMet(c, noMinutesPrice)
			return
		}
		priceUnit = *transport.MinutePrice
	} else if input.RentType == constants.RentTypeDays {
		if transport.DayPrice == nil {
			h.sendConditionNotMet(c, noDaysPrice)
			return
		}
		priceUnit = *transport.DayPrice
	}

	timeStart := time.Now()
	rent, err := h.services.Rent.Create(username, transportID, timeStart, nil, priceUnit, input.RentType)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.GetRentTransportNewOutput{
		Rent: models.GetRentOutput{
			ID:        rent.ID,
			Account:   rent.Account.Username,
			Transport: rent.Transport.Identifier,
			TimeStart: rent.TimeStart,
			TimeEnd:   rent.TimeEnd,
			PriceUnit: rent.PriceUnit,
			PriceType: rent.PriceType,
		},
	})
}
