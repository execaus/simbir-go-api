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

// AdminGetTransports
// @Summary      Get transport list
// @Description  Getting list information about transport.
// @Tags         admin-transport
// @Accept       json
// @Produce      json
// @Param        count query number true "-"
// @Param        start query number true "-"
// @Param        transportType query string true "-"
// @Success      200  {object}  models.AdminGetTransportsOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Transport [get]
func (h *Handler) AdminGetTransports(c *gin.Context) {
	var input models.AdminGetTransportsInput

	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	transports, err := h.services.Transport.GetList(input.Start, input.Count, input.TransportType)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	output := models.AdminGetTransportsOutput{Transports: make([]*models.AdminGetTransportOutput, len(transports))}
	for i, transport := range transports {
		output.Transports[i] = &models.AdminGetTransportOutput{
			Transport: &models.GetTransportOutput{
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
			},
			IsDeleted: transport.IsDeleted,
		}
	}

	h.sendOKWithBody(c, output)
}

// AdminCreateTransport
// @Summary      Create transport
// @Description  Adding new transportation.
// @Tags         admin-transport
// @Accept       json
// @Produce      json
// @Success      200
// @Param        input body models.AdminCreateTransportInput true "-"
// @Success      200  {object}  models.AdminCreateTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Transport [post]
func (h *Handler) AdminCreateTransport(c *gin.Context) {
	var input models.AdminCreateTransportInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
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
		OwnerID:       input.OwnerID,
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
		IsDeleted:     false,
	})
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.AdminCreateTransportOutput{
		Transport: &models.GetTransportOutput{
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
		},
	})
}

// AdminUpdateTransport
// @Summary      Update transport
// @Description  Update transport by id.
// @Tags         admin-transport
// @Accept       json
// @Produce      json
// @Param        id path string true "-"
// @Param        input body models.AdminUpdateTransportInput true "-"
// @Success      200  {object}  models.AdminUpdateTransportOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Transport/{id} [put]
func (h *Handler) AdminUpdateTransport(c *gin.Context) {
	var input models.AdminUpdateTransportInput

	transportID, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
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

	if !isExist {
		h.sendInvalidRequest(c, transportIsNotExist)
		return
	}

	transport, err := h.services.Transport.Update(transportID, &models.Transport{
		OwnerID:       input.OwnerID,
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
		IsDeleted:     false,
	})
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.AdminUpdateTransportOutput{
		Transport: &models.GetTransportOutput{
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
		},
	})
}

// AdminDeleteTransport
// @Summary      Delete transport
// @Description  Deleting vehicles by id.
// @Tags         admin-transport
// @Accept       json
// @Produce      json
// @Param        id path string true "-"
// @Success      204
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Transport/{id} [delete]
func (h *Handler) AdminDeleteTransport(c *gin.Context) {
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

	if err = h.services.Transport.Remove(transportID); err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOK(c)
}
