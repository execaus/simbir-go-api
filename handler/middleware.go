package handler

import (
	"errors"
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
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

	username, err := h.services.ParseToken(headerParts[1])
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

	c.Set(accountContextName, username)
	c.Set(accountTokenName, headerParts[1])
}

func getAccountContext(c *gin.Context) (string, error) {
	username := c.GetString(accountContextName)
	if username == "" {
		exloggo.Error(accountContextEmpty)
		return "", errors.New(accountContextEmpty)
	}
	return username, nil
}

func getAccountToken(c *gin.Context) (string, error) {
	token := c.GetString(accountTokenName)
	if token == "" {
		exloggo.Error(accountContextEmpty)
		return "", errors.New(accountContextEmpty)
	}
	return token, nil
}
