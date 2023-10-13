package handler

import (
	"github.com/gin-gonic/gin"
	"simbir-go-api/models"
)

// @Summary      Registration
// @Description  Registers a new user and returns the authorization token jwt.
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      200
// @Param        input body models.SignUpInput true "-"
// @Failure      400  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Account/SignUp [post]
func (h *Handler) signUp(c *gin.Context) {
	var input *models.SignUpInput

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

	_, err = h.services.Account.SignUp(input.Username, input.Password)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	token, err := h.services.Account.GenerateToken(input.Username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, token)
}

// @Summary      Authorization
// @Description  Generates and returns a new jwt user token.
// @Tags         account
// @Accept       json
// @Produce      json
// @Success      200
// @Param        input body models.SignInInput true "-"
// @Failure      400  {object}  handler.Error
// @Failure      500  {object}  handler.Error
// @Router       /Account/SignIn [post]
func (h *Handler) signIn(c *gin.Context) {
	var input *models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		h.sendInvalidRequest(c, err.Error())
		return
	}

	isExist, err := h.services.Account.IsExist(input.Username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isExist {
		h.sendInvalidRequest(c, accountIsNotExist)
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

	token, err := h.services.Account.GenerateToken(account.Username)
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	h.sendOKWithBody(c, token)
}
