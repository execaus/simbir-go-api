package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	stringEmpty = ""
	stringNull  = "null"
)

func getNumberParam(c *gin.Context, key string) (int32, error) {
	stringID := c.Param(key)
	if stringID == stringEmpty || stringID == stringNull {
		return 0, errors.New("param is not valid")
	}

	id, err := strconv.ParseInt(stringID, 10, 32)
	if err != nil {
		return 0, errors.New("param is not valid")
	}

	if id <= 0 {
		return 0, errors.New("param is not valid")
	}

	return int32(id), nil
}

//func getStringParam(c *gin.Context, key string) (string, error) {
//	value := c.Param(key)
//	if value == stringEmpty || value == stringNull {
//		return "", errors.New("param is not valid")
//	}
//
//	return value, nil
//}
