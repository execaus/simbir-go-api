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
// @Router       /Admin/Account/{username} [get]
func (h *Handler) AdminGetAccount(c *gin.Context) {
	username, err := getStringParam(c, "username")
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

	account, err := h.services.Account.GetByUsername(username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.AdminGetAccountOutput{
		Account: models.GetAccountOutput{
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
// @Success      200
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

	isExist, err := h.services.Account.IsExist(input.Username)
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
		Username: account.Username,
		IsAdmin:  account.IsAdmin(),
		Balance:  account.Balance,
	}})
}
