package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
)

// PaymentHesoyam
// @Summary      Hesoyam
// @Description  Adds 250,000 cash units to the balance of the account with id={accountId}..
// @Tags         payment
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.PaymentHesoyamAdminOutput
// @Success      200  {object}  models.PaymentHesoyamUserOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      403  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Payment/Hesoyam/{id} [post]
func (h *Handler) PaymentHesoyam(c *gin.Context) {
	targetUserID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	userID, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	isExist, err := h.services.Account.IsExistByID(targetUserID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	account, err := h.services.Account.GetByID(userID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, accountIsNotExist)
		return
	}

	if targetUserID != userID {
		if !account.IsAdmin() {
			h.sendAccessDenied(c, accessDeniedOnlyAdmin)
			return
		}
	}

	updatedAccount, err := h.services.Account.Hesoyam(targetUserID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if account.IsAdmin() {
		h.sendOKWithBody(c, &models.PaymentHesoyamAdminOutput{Account: models.AdminGetAccountOutput{
			Account: models.GetAccountOutput{
				ID:       updatedAccount.ID,
				Username: updatedAccount.Username,
				IsAdmin:  updatedAccount.IsAdmin(),
				Balance:  updatedAccount.Balance,
			},
			IsDeleted: updatedAccount.IsDeleted,
		}})
		return
	}

	h.sendOKWithBody(c, &models.PaymentHesoyamUserOutput{Account: models.GetAccountOutput{
		ID:       updatedAccount.ID,
		Username: updatedAccount.Username,
		IsAdmin:  updatedAccount.IsAdmin(),
		Balance:  updatedAccount.Balance,
	}})
}
