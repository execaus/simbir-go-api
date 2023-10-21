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
		h.sendInvalidRequest(c, transportIsNotExist)
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

// GetTransport
// @Summary      Get transport
// @Description  Getting information about transport by id.
// @Tags         transport
// @Accept       json
// @Produce      json
// @Success      200
// @Param        id path string true "-"
// @Success      200  {object}  models.GetTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Transport/{id} [get]
func (h *Handler) GetTransport(c *gin.Context) {
	transportID, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
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

	isRemoved, err := h.services.Transport.IsRemoved(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isRemoved {
		h.sendResourceDeleted(c, transportIsDeleted)
		return
	}

	transport, err := h.services.Transport.Get(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.GetTransportOutput{
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

// UpdateTransport
// @Summary      Update transport
// @Description  Update transport by id.
// @Tags         transport
// @Accept       json
// @Produce      json
// @Param        id path string true "-"
// @Param        input body models.UpdateTransportInput true "-"
// @Success      200  {object}  models.UpdateTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Transport/{id} [put]
func (h *Handler) UpdateTransport(c *gin.Context) {
	var input models.UpdateTransportInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	username, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	transportID, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
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

	isRemoved, err := h.services.Transport.IsRemoved(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isRemoved {
		h.sendResourceDeleted(c, transportIsDeleted)
		return
	}

	isOwner, err := h.services.Transport.IsOwner(transportID, username)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isOwner {
		h.sendInvalidRequest(c, accountNotTransportOwner)
		return
	}

	currentTransport, err := h.services.Transport.Get(transportID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	transport, err := h.services.Transport.Update(transportID, &models.Transport{
		OwnerID:       username,
		CanBeRented:   input.CanBeRented,
		TransportType: currentTransport.TransportType,
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
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.UpdateTransportOutput{Transport: &models.GetTransportOutput{
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
	}})
}

// DeleteTransport
// @Summary      Delete transport
// @Description  Deleting vehicles by id.
// @Tags         transport
// @Accept       json
// @Produce      json
// @Param        id path string true "-"
// @Success      204
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      410  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Transport/{id} [delete]
func (h *Handler) DeleteTransport(c *gin.Context) {
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

	isExist, err := h.services.Transport.IsExist(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
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

	isOwner, err := h.services.Transport.IsOwner(transportID, username)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isOwner {
		h.sendInvalidRequest(c, accountNotTransportOwner)
		return
	}

	if err = h.services.Transport.Remove(transportID); err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOK(c)
}
