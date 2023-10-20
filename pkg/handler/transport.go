package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
)

// CreateTransport
// @Summary      Create transport
// @Description  Adding new transportation.
// @Tags         transport
// @Accept       json
// @Produce      json
// @Success      200
// @Param        input body models.CreateTransportInput true "-"
// @Success      200  {object}  models.CreateTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Transport [post]
func (h *Handler) CreateTransport(c *gin.Context) {
	var input models.CreateTransportInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	username, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	if err = input.Validate(); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Transport.IsExist(input.Identifier)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isExist {
		h.sendInvalidRequest(c, transportIsExist)
		return
	}

	transport, err := h.services.Transport.Create(&models.Transport{
		OwnerID:       username,
		CanBeRented:   input.CanBeRented,
		TransportType: input.TransportType,
		Model:         input.Model,
		Color:         input.Color,
		Identifier:    input.Identifier,
		Description:   input.Description,
		Latitude:      *input.Latitude,
		Longitude:     *input.Longitude,
		MinutePrice:   input.MinutePrice,
		DayPrice:      input.DayPrice,
	})
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.CreateTransportOutput{Transport: transport})
}
