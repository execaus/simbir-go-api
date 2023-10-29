package handler

import (
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
	"math"
	"simbir-go-api/models"
)

// AdminGetRent
// @Summary      Rent information
// @Description  Return rent information by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.AdminGetRentOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Rent/{id} [get]
func (h *Handler) AdminGetRent(c *gin.Context) {
	rentID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Rent.IsExist(rentID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, rentIsNotExist)
		return
	}

	rent, err := h.services.Rent.Get(rentID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.AdminGetRentOutput{
		Rent: models.GetRentOutput{
			ID:         rent.ID,
			Account:    rent.Account,
			Transport:  rent.Transport,
			TimeStart:  rent.TimeStart,
			TimeEnd:    rent.TimeEnd,
			PriceUnit:  rent.PriceUnit,
			PriceType:  rent.PriceType,
			FinalPrice: rent.FinalPrice,
		},
		IsDeleted: rent.IsDeleted,
	})
}

// AdminGetUserRentHistory
// @Summary      User rent history
// @Description  Return user rent history by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.AdminGetUserHistoryOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/UserHistory/{id} [get]
func (h *Handler) AdminGetUserRentHistory(c *gin.Context) {
	userID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExistByID(userID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, accountIsNotExist)
		return
	}

	rents, err := h.services.Rent.GetListFromUserID(userID, 0, math.MaxInt32)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	var output models.AdminGetUserHistoryOutput
	for _, rent := range rents {
		output.Rents = append(output.Rents, models.AdminGetRentOutput{
			Rent: models.GetRentOutput{
				ID:         rent.ID,
				Account:    rent.Account,
				Transport:  rent.Transport,
				TimeStart:  rent.TimeStart,
				TimeEnd:    rent.TimeEnd,
				PriceUnit:  rent.PriceUnit,
				PriceType:  rent.PriceType,
				FinalPrice: rent.FinalPrice,
			},
			IsDeleted: rent.IsDeleted,
		})
	}

	h.sendOKWithBody(c, output)
}

// AdminGetTransportRentHistory
// @Summary      Transport rent history
// @Description  Return transport rent history by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.AdminGetTransportHistoryOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/TransportHistory/{id} [get]
func (h *Handler) AdminGetTransportRentHistory(c *gin.Context) {
	transportID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Transport.IsExistByID(transportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, transportIsNotExist)
		return
	}

	rents, err := h.services.Rent.GetListFromTransportID(transportID, 0, math.MaxInt32)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	var output models.AdminGetTransportHistoryOutput
	for _, rent := range rents {
		output.Rents = append(output.Rents, models.AdminGetRentOutput{
			Rent: models.GetRentOutput{
				ID:         rent.ID,
				Account:    rent.Account,
				Transport:  rent.Transport,
				TimeStart:  rent.TimeStart,
				TimeEnd:    rent.TimeEnd,
				PriceUnit:  rent.PriceUnit,
				PriceType:  rent.PriceType,
				FinalPrice: rent.FinalPrice,
			},
			IsDeleted: rent.IsDeleted,
		})
	}

	h.sendOKWithBody(c, output)
}

// AdminCreateRent
// @Summary      Create rent
// @Description  Create new rent.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Param        input body models.AdminCreateRentInput true "-"
// @Success      200  {object}  models.AdminCreateRentOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Rent [post]
func (h *Handler) AdminCreateRent(c *gin.Context) {
	var input models.AdminCreateRentInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if input.TimeEnd != nil {
		if input.TimeStart.After(*input.TimeEnd) {
			exloggo.Error(invalidTimeRange)
			h.sendInvalidRequest(c, invalidTimeRange)
			return
		}
	}

	isTransportExist, err := h.services.Transport.IsExistByID(input.TransportID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isTransportExist {
		h.sendInvalidRequest(c, transportIsNotExist)
		return
	}

	isAccountExist, err := h.services.Account.IsExistByID(input.UserID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isAccountExist {
		h.sendGeneralException(c, accountIsNotExist)
		return
	}

	rent, err := h.services.Rent.Create(&models.Rent{
		Account:   input.UserID,
		Transport: input.TransportID,
		TimeStart: input.TimeStart,
		TimeEnd:   input.TimeEnd,
		PriceUnit: input.PriceUnit,
		PriceType: input.PriceType,
	})
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.AdminCreateRentOutput{
		Rent: models.AdminGetRentOutput{
			Rent: models.GetRentOutput{
				ID:         rent.ID,
				Account:    rent.Account,
				Transport:  rent.Transport,
				TimeStart:  rent.TimeStart,
				TimeEnd:    rent.TimeEnd,
				PriceUnit:  rent.PriceUnit,
				PriceType:  rent.PriceType,
				FinalPrice: rent.FinalPrice,
			},
			IsDeleted: rent.IsDeleted,
		},
	})
}

// AdminEndRent
// @Summary      End rent
// @Description  Completion of the lease of transportation under the lease id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200
// @Param        id query number true "-"
// @Param        input body models.AdminEndRentInput true "-"
// @Success      200  {object}  models.AdminEndRentOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Admin/Rent/End/{id} [post]
func (h *Handler) AdminEndRent(c *gin.Context) {
	var input models.AdminEndRentInput

	rentID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isRentExist, err := h.services.Rent.IsExist(rentID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isRentExist {
		h.sendInvalidRequest(c, rentIsNotExist)
		return
	}

	rent, err := h.services.Rent.End(rentID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	account, err := h.services.Account.GetByID(rent.Account)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	account.Balance -= *rent.FinalPrice
	_, err = h.services.Account.Update(account)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	transport, err := h.services.Transport.Get(rent.Transport)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	transport.Latitude = *input.Latitude
	transport.Longitude = *input.Longitude
	_, err = h.services.Transport.Update(transport)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.AdminEndRentOutput{Rent: models.AdminGetRentOutput{
		Rent: models.GetRentOutput{
			ID:         rent.ID,
			Account:    rent.Account,
			Transport:  rent.Transport,
			TimeStart:  rent.TimeStart,
			TimeEnd:    rent.TimeEnd,
			PriceUnit:  rent.PriceUnit,
			PriceType:  rent.PriceType,
			FinalPrice: rent.FinalPrice,
		},
		IsDeleted: rent.IsDeleted,
	}})
}

// AdminUpdateRent
// @Summary      Update rent
// @Description  Changing a lease record by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200
// @Param        id query number true "-"
// @Param        input body models.AdminUpdateRentInput true "-"
// @Success      200  {object}  models.AdminUpdateRentOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Admin/Rent/{id} [put]
func (h *Handler) AdminUpdateRent(c *gin.Context) {
	var input models.AdminUpdateRentInput

	rentID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	if err = c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isRentExist, err := h.services.Rent.IsExist(rentID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isRentExist {
		h.sendInvalidRequest(c, rentIsNotExist)
		return
	}

	rent, err := h.services.Rent.Update(&models.Rent{
		ID:        rentID,
		Account:   input.UserID,
		Transport: input.TransportID,
		TimeStart: input.TimeStart,
		TimeEnd:   input.TimeEnd,
		PriceUnit: input.PriceUnit,
		PriceType: input.PriceType,
	})
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.AdminUpdateRentOutput{
		Rent: models.AdminGetRentOutput{
			Rent: models.GetRentOutput{
				ID:         rent.ID,
				Account:    rent.Account,
				Transport:  rent.Transport,
				TimeStart:  rent.TimeStart,
				TimeEnd:    rent.TimeEnd,
				PriceUnit:  rent.PriceUnit,
				PriceType:  rent.PriceType,
				FinalPrice: rent.FinalPrice,
			},
			IsDeleted: rent.IsDeleted,
		},
	})
}

// AdminDeleteRent
// @Summary      Delete rent
// @Description  Deleting rental information by id.
// @Tags         admin-rent
// @Accept       json
// @Produce      json
// @Success      200
// @Param        id query number true "-"
// @Success      204
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Admin/Rent/{id} [put]
func (h *Handler) AdminDeleteRent(c *gin.Context) {
	rentID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Rent.IsExist(rentID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, rentIsNotExist)
		return
	}

	isRemoved, err := h.services.Rent.IsRemoved(rentID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if !isRemoved {
		h.sendInvalidRequest(c, rentIsRemoved)
		return
	}

	if err = h.services.Rent.Remove(rentID); err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOK(c)
}
