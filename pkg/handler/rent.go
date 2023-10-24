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
