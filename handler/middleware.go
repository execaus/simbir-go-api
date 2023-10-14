package handler

import (
	"errors"
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	accountContextName            = "account-context"
	authorizationHeader           = "Authorization"
	headerAuthorizationPartsCount = 2
)

func (h *Handler) accountIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		exloggo.Error(accountTokenNotFound)
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

	c.Set(accountContextName, username)
}

func getAccountContext(c *gin.Context) (string, error) {
	var username = c.GetString(accountContextName)
	if username == "" {
		exloggo.Error(accountContextEmpty)
		return "", errors.New(accountContextEmpty)
	}
	return username, nil
}
