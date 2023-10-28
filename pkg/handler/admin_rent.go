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

	isRemoved, err := h.services.Rent.IsRemoved(int32(rentID))
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
		IsDeleted: isRemoved,
	})
}
