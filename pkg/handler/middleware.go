package handler

import (
	"errors"
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
	"simbir-go-api/constants"
	"strconv"
	"strings"
)

const (
	accountContextName            = "account-context"
	accountTokenName              = "account-token"
	authorizationHeader           = "Authorization"
	invalidJWTToken               = "invalid token"
	headerAuthorizationPartsCount = 2
)

func (h *Handler) accountIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		exloggo.Info(accountTokenNotFound)
		h.sendUnAuthenticated(c, invalidAuthorizationHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != headerAuthorizationPartsCount {
		exloggo.Error(header)
		h.sendUnAuthenticated(c, invalidAuthorizationHeader)
		return
	}

	userID, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	isAccess, err := h.services.Account.IsValidToken(headerParts[1])
	if err != nil {
		h.sendGeneralException(c, serverError)
		return
	}

	if !isAccess {
		h.sendUnAuthenticated(c, invalidJWTToken)
		return
	}

	c.Set(accountContextName, strconv.Itoa(int(userID)))
	c.Set(accountTokenName, headerParts[1])
}

func (h *Handler) onlyAdmin(c *gin.Context) {
	h.accountIdentity(c)

	if c.IsAborted() {
		return
	}

	userID, err := getAccountContext(c)
	if err != nil {
		exloggo.Error(err.Error())
		h.sendGeneralException(c, serverError)
		return
	}

	roles, err := h.services.Account.GetRoles(userID)
	if err != nil {
		exloggo.Error(err.Error())
		h.sendGeneralException(c, serverError)
		return
	}

	for _, role := range roles {
		if role == constants.RoleAdmin {
			c.Next()
			return
		}
	}

	h.sendAccessDenied(c, accessDeniedOnlyAdmin)
}

func getAccountContext(c *gin.Context) (int32, error) {
	userID := c.GetString(accountContextName)
	if userID == "" {
		exloggo.Error(accountContextEmpty)
		return 0, errors.New(accountContextEmpty)
	}

	id, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		exloggo.Error(err.Error())
		return 0, errors.New(serverError)
	}

	return int32(id), nil
}

func getAccountToken(c *gin.Context) (string, error) {
	token := c.GetString(accountTokenName)
	if token == "" {
		exloggo.Error(accountContextEmpty)
		return "", errors.New(accountContextEmpty)
	}
	return token, nil
}
