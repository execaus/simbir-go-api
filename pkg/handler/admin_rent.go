package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
)

// GetAdminRent
// @Summary      Rent information
// @Description  Return rent information by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAdminRentOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Rent/{id} [get]
func (h *Handler) GetAdminRent(c *gin.Context) {
	rentID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Rent.IsExist(int32(rentID))
	if err != nil {
		h.sendGeneralException(c, serverError)
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

	h.sendOKWithBody(c, &models.GetAdminRentOutput{
		Rent: models.GetRentOutput{
			ID:        rent.ID,
			Account:   rent.Account.Username,
			Transport: rent.Transport.TransportType,
			TimeStart: rent.TimeStart,
			TimeEnd:   rent.TimeEnd,
			PriceUnit: rent.PriceUnit,
			PriceType: rent.PriceType,
		},
		IsDeleted: rent.IsDeleted,
	})
}

// GetAdminUserRentHistory
// @Summary      User rent history
// @Description  Return user rent history by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAdminUserHistoryOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/UserHistory/{id} [get]
func (h *Handler) GetAdminUserRentHistory(c *gin.Context) {
	username, err := getStringParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExist(username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, accountIsNotExist)
		return
	}

	rents, err := h.services.Rent.GetListFromUsername(username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	var output models.GetAdminUserHistoryOutput
	for _, rent := range rents {
		output.Rents = append(output.Rents, models.GetAdminRentOutput{
			Rent: models.GetRentOutput{
				ID:        rent.ID,
				Account:   rent.Account.Username,
				Transport: rent.Transport.Identifier,
				TimeStart: rent.TimeStart,
				TimeEnd:   rent.TimeEnd,
				PriceUnit: rent.PriceUnit,
				PriceType: rent.PriceType,
			},
			IsDeleted: rent.IsDeleted,
		})
	}

	h.sendOKWithBody(c, output)
}

// GetAdminTransportRentHistory
// @Summary      Transport rent history
// @Description  Return transport rent history by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAdminTransportHistoryOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/TransportHistory/{id} [get]
func (h *Handler) GetAdminTransportRentHistory(c *gin.Context) {
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

	rents, err := h.services.Rent.GetListFromTransport(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	var output models.GetAdminTransportHistoryOutput
	for _, rent := range rents {
		output.Rents = append(output.Rents, models.GetAdminRentOutput{
			Rent: models.GetRentOutput{
				ID:        rent.ID,
				Account:   rent.Account.Username,
				Transport: rent.Transport.Identifier,
				TimeStart: rent.TimeStart,
				TimeEnd:   rent.TimeEnd,
				PriceUnit: rent.PriceUnit,
				PriceType: rent.PriceType,
			},
			IsDeleted: rent.IsDeleted,
		})
	}

	h.sendOKWithBody(c, output)
}
