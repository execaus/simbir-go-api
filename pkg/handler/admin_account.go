package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/constants"
	"simbir-go-api/models"
)

// AdminGetAccounts
// @Summary      Accounts information
// @Description  Returns a list of accounts by sampling condition.
// @Tags         admin-account
// @Accept       json
// @Produce      json
// @Param        count query number true "-"
// @Param        start query number true "-"
// @Success      200  {object}  models.AdminGetAccountsOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Account [get]
func (h *Handler) AdminGetAccounts(c *gin.Context) {
	var input models.AdminGetAccountsInput

	if err := c.ShouldBindQuery(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	accounts, err := h.services.Account.GetList(input.Start, input.Count)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	output := models.AdminGetAccountsOutput{Accounts: make([]*models.GetAccountOutput, 0)}
	for _, account := range accounts {
		output.Accounts = append(output.Accounts, &models.GetAccountOutput{
			ID:       account.ID,
			Username: account.Username,
			IsAdmin:  account.IsAdmin(),
			Balance:  account.Balance,
		})
	}

	h.sendOKWithBody(c, output)
}

// AdminGetAccount
// @Summary      Account information
// @Description  Returns account information.
// @Tags         admin-account
// @Accept       json
// @Produce      json
// @Param        username path string true "-"
// @Success      200  {object}  models.AdminGetAccountOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Account/{id} [get]
func (h *Handler) AdminGetAccount(c *gin.Context) {
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

	account, err := h.services.Account.GetByID(userID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.AdminGetAccountOutput{
		Account: models.GetAccountOutput{
			ID:       account.ID,
			Username: account.Username,
			IsAdmin:  account.IsAdmin(),
			Balance:  account.Balance,
		},
	})
}

// AdminCreateAccount
// @Summary      Create account
// @Description  Creating a new account by the administrator.
// @Tags         admin-account
// @Accept       json
// @Produce      json
// @Param        input body models.AdminCreateAccountInput true "-"
// @Success      200  {object}  models.AdminCreateAccountOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Account/ [post]
func (h *Handler) AdminCreateAccount(c *gin.Context) {
	var input models.AdminCreateAccountInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExistByUsername(input.Username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isExist {
		h.sendInvalidRequest(c, accountIsExist)
		return
	}

	var role string
	if input.IsAdmin {
		role = constants.RoleAdmin
	} else {
		role = constants.RoleUser
	}

	account, err := h.services.Account.Create(input.Username, input.Password, role, *input.Balance)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.AdminCreateAccountOutput{Account: models.GetAccountOutput{
		ID:       account.ID,
		Username: account.Username,
		IsAdmin:  account.IsAdmin(),
		Balance:  account.Balance,
	}})
}

// AdminUpdateAccount
// @Summary      Update account
// @Description  Account administrator changes account by username.
// @Tags         admin-account
// @Accept       json
// @Produce      json
// @Param        username path string true "-"
// @Param        input body models.AdminUpdateAccountInput true "-"
// @Success      200  {object}  models.AdminUpdateAccountOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      410  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Account/{id} [put]
func (h *Handler) AdminUpdateAccount(c *gin.Context) {
	var input models.AdminUpdateAccountInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	userID, err := getNumberParam(c, "id")
	if err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isRemoved, err := h.services.Account.IsRemovedByID(userID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isRemoved {
		h.sendResourceDeleted(c, accountIsDeleted)
		return
	}

	currentAccount, err := h.services.Account.GetByID(userID)
	if err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	if currentAccount.Username != input.Username {
		isExist, err := h.services.Account.IsExistByUsername(input.Username)
		if err != nil {
			h.sendGeneralException(c, serverError)
			return
		}

		if isExist {
			h.sendInvalidRequest(c, accountIsExist)
			return
		}
	}

	roles := []string{constants.RoleUser}
	if input.IsAdmin {
		roles = append(roles, constants.RoleAdmin)
	}

	updatedAccount, err := h.services.Account.Update(&models.Account{
		Username: input.Username,
		Password: input.Password,
		Balance:  *input.Balance,
		Roles:    roles,
	})
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.AdminUpdateAccountOutput{Account: models.GetAccountOutput{
		ID:       updatedAccount.ID,
		Username: updatedAccount.Username,
		IsAdmin:  updatedAccount.IsAdmin(),
		Balance:  updatedAccount.Balance,
	}})
}

// AdminRemoveAccount
// @Summary      Delete account
// @Description  Deleting account by id.
// @Tags         admin-account
// @Accept       json
// @Produce      json
// @Param        id path string true "-"
// @Success      204
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      410  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Admin/Account/{id} [delete]
func (h *Handler) AdminRemoveAccount(c *gin.Context) {
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

	isRemoved, err := h.services.Account.IsRemovedByID(userID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isRemoved {
		h.sendResourceDeleted(c, accountIsDeleted)
		return
	}

	if err = h.services.Account.Remove(userID); err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOK(c)
}
