package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
)

// SignUp
// @Summary      Registration
// @Description  Registers a new user and returns the authorization token jwt.
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      200
// @Param        input body models.SignUpInput true "-"
// @Success      200  {object}  models.SignUpOutput
// @Failure      400  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Account/SignUp [post]
func (h *Handler) SignUp(c *gin.Context) {
	var input *models.SignUpInput

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

	account, err := h.services.Account.SignUp(input.Username, input.Password)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	token, err := h.services.Account.GenerateToken(account.ID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.SignUpOutput{
		Token: token,
	})
}

// SignIn
// @Summary      Authorization
// @Description  Generates and returns a new jwt user token.
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      200
// @Param        input body models.SignInInput true "-"
// @Success      200  {object}  models.SignInOutput
// @Failure      400  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Account/SignIn [post]
func (h *Handler) SignIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExistByUsername(input.Username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, invalidAuthData)
		return
	}

	isRemoved, err := h.services.Account.IsRemovedByUsername(input.Username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if isRemoved {
		h.sendInvalidRequest(c, invalidAuthData)
		return
	}

	account, err := h.services.Account.Authorize(input.Username, input.Password)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if account == nil {
		h.sendInvalidRequest(c, invalidAuthData)
		return
	}

	token, err := h.services.Account.GenerateToken(account.ID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, &models.SignInOutput{
		Token: token,
	})
}

// GetAccount
// @Summary      Account information
// @Description  Returns the full data of the request author.
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAccountOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Account/Me [get]
func (h *Handler) GetAccount(c *gin.Context) {
	userID, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
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

	h.sendOKWithBody(c, &models.GetAccountOutput{
		ID:       userID,
		Username: account.Username,
		IsAdmin:  account.IsAdmin(),
		Balance:  account.Balance,
	})
}

// UpdateAccount
// @Summary      Update account
// @Description  Updates account data and returns a new authorization token.
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        input body models.UpdateAccountInput true "-"
// @Success      200  {object}  models.UpdateAccountOutput
// @Failure      400  {object}  handler.Error
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Account/Update [put]
func (h *Handler) UpdateAccount(c *gin.Context) {
	var input models.UpdateAccountInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	token, err := getAccountToken(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	userID, err := getAccountContext(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	currentAccount, err := h.services.Account.GetByID(userID)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if currentAccount.Username != input.Username {
		isExist, err := h.services.Account.IsExistByUsername(input.Username)
		if err != nil {
			h.sendGeneralException(c, serverError)
			return
		}

		if isExist {
			h.sendInvalidRequest(c, usernameIsBusy)
			return
		}
	}

	updatedAccount, err := h.services.Account.Update(&models.Account{
		ID:       userID,
		Username: input.Username,
		Password: input.Password,
		Balance:  currentAccount.Balance,
		Roles:    currentAccount.Roles,
	})
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if err = h.services.Account.BlockToken(token); err != nil {
		h.sendGeneralException(c, err.Error())
		return
	}

	h.sendOKWithBody(c, &models.UpdateAccountOutput{
		Account: models.GetAccountOutput{
			ID:       updatedAccount.ID,
			Username: updatedAccount.Username,
			IsAdmin:  updatedAccount.IsAdmin(),
			Balance:  updatedAccount.Balance,
		},
	})
}

// SignOut
// @Summary      Account logout
// @Description  Blocks the authorization token.
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      204
// @Failure      401  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Security     BearerAuth
// @Router       /Account/SignOut [post]
func (h *Handler) SignOut(c *gin.Context) {
	token, err := getAccountToken(c)
	if err != nil {
		h.sendUnAuthenticated(c, serverError)
		return
	}

	if err = h.services.Account.BlockToken(token); err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOK(c)
}
