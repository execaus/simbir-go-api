package handler

import (
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	accessDenied          = "accessDenied"
	generalException      = "generalException"
	itemNotFound          = "itemNotFound"
	invalidRequest        = "invalidRequest"
	resourceModified      = "resourceModified"
	unAuthenticated       = "unAuthenticated"
	sendEmail             = "sendEmail"
	sendEmailError        = "smtp server error"
	conditionNotMet       = "conditionNotMet"
	resourceDeleted       = "resourceDeleted"
	notConfirmed          = "notConfirmed"
	timeoutExpired        = "timeoutExpired"
	headerException       = "headerException"
	sendEmailErrorStatus  = 506
	headerExceptionStatus = 432
)

const (
	serverError                = "internal server error"
	accountIsExist             = "user with this name already exists"
	accountIsNotExist          = "user with this name does not exist"
	accountIsDeleted           = "account is deleted"
	invalidAuthData            = "invalid user data"
	accountContextEmpty        = "account context was not found"
	accountTokenNotFound       = "the user token was not detected in the request header"
	invalidAuthorizationHeader = "invalid \"Authorization\" header"
	usernameIsBusy             = "username is busy"
	accessDeniedOnlyAdmin      = "access denied, available only to administrator"
	transportIsNotExist        = "transport is not exist"
	transportIsExist           = "transport is exist"
	accountIsNotTransportOwner = "account is not transport owner"
	accountIsTransportOwner    = "account is transport owner"
	transportIsDeleted         = "transport is deleted"
	accountIsNotRenter         = "account is not renter"
	rentIsNotExist             = "rentIsNotExist"
	transportIsNotRent         = "transport is not rent"
	transportInRent            = "the transport in rent"
	noMinutesPrice             = "no price per minutes has been set"
	noDaysPrice                = "no price per days has been set"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) sendOK(c *gin.Context) {
	h.sendResponse(c, http.StatusNoContent, gin.H{})
}

func (h *Handler) sendOKWithBody(c *gin.Context, data interface{}) {
	h.sendResponse(c, http.StatusOK, data)
}

func (h *Handler) sendAccessDenied(c *gin.Context, message string) {
	data := Error{
		Code:    accessDenied,
		Message: message,
	}

	h.sendResponse(c, http.StatusForbidden, data)
}

func (h *Handler) sendGeneralException(c *gin.Context, message string) {
	data := Error{
		Code:    generalException,
		Message: message,
	}
	h.sendResponse(c, http.StatusInternalServerError, data)
}

func (h *Handler) sendItemNotFound(c *gin.Context, message string) {
	data := Error{
		Code:    itemNotFound,
		Message: message,
	}
	exloggo.RequestResult(c, message)
	h.sendResponse(c, http.StatusNotFound, data)
}

func (h *Handler) sendInvalidRequest(c *gin.Context, message string) {
	data := Error{
		Code:    invalidRequest,
		Message: message,
	}
	exloggo.RequestResult(c, message)
	h.sendResponse(c, http.StatusBadRequest, data)
}

func (h *Handler) sendResourceModified(c *gin.Context, message string) {
	data := Error{
		Code:    resourceModified,
		Message: message,
	}
	h.sendResponse(c, http.StatusInternalServerError, data)
}

func (h *Handler) sendUnAuthenticated(c *gin.Context, message string) {
	data := Error{
		Code:    unAuthenticated,
		Message: message,
	}
	h.sendResponse(c, http.StatusUnauthorized, data)
}

func (h *Handler) sendSendEmailError(c *gin.Context) {
	data := Error{
		Code:    sendEmail,
		Message: sendEmailError,
	}
	h.sendResponse(c, sendEmailErrorStatus, data)
}

func (h *Handler) sendConditionNotMet(c *gin.Context, message string) {
	data := Error{
		Code:    conditionNotMet,
		Message: message,
	}
	exloggo.RequestResult(c, message)
	h.sendResponse(c, http.StatusPreconditionFailed, data)
}

func (h *Handler) sendResourceDeleted(c *gin.Context, message string) {
	data := Error{
		Code:    resourceDeleted,
		Message: message,
	}
	exloggo.RequestResult(c, message)
	h.sendResponse(c, http.StatusGone, data)
}

func (h *Handler) sendNotConfirmed(c *gin.Context, message string) {
	data := Error{
		Code:    notConfirmed,
		Message: message,
	}
	exloggo.RequestResult(c, message)
	h.sendResponse(c, http.StatusPreconditionFailed, data)
}

func (h *Handler) sendTimeoutExpired(c *gin.Context, message string) {
	data := Error{
		Code:    timeoutExpired,
		Message: message,
	}
	h.sendResponse(c, http.StatusPreconditionFailed, data)
}

func (h *Handler) sendHeaderException(c *gin.Context, message string) {
	data := Error{
		Code:    headerException,
		Message: message,
	}
	exloggo.RequestResult(c, message)
	h.sendResponse(c, headerExceptionStatus, data)
}

func (h *Handler) sendResponse(c *gin.Context, status int, data interface{}) {
	c.AbortWithStatusJSON(status, data)
}
