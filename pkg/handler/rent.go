package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
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
