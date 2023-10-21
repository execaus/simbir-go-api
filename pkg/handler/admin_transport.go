package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
)

// AdminGetTransport
// @Summary      Get transport
// @Description  Getting information about transport by id.
// @Tags         admin-transport
// @Accept       json
// @Produce      json
// @Param        id path string true "-"
// @Success      200  {object}  models.AdminGetTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Transport/{id} [get]
func (h *Handler) AdminGetTransport(c *gin.Context) {
	transportID, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	transport, err := h.services.Transport.Get(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	isExist, err := h.services.Transport.IsExist(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, transportIsNotExist)
		return
	}

	h.sendOKWithBody(c, &models.AdminGetTransportOutput{
		Transport: &models.GetTransportOutput{
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
			OwnerID:       transport.OwnerID,
		},
		IsDeleted: transport.IsDeleted,
	})
}
