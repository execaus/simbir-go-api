package constants

import (
	"errors"
	"fmt"
	"strings"
)

const (
	RentTypeMinutes = "MINUTES"
	RentTypeDays    = "DAYS"
)

var rentTypes = []string{RentTypeMinutes, RentTypeDays}

func CheckRentType(value string) error {
	for _, rentType := range rentTypes {
		if strings.ToUpper(value) == rentType {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("%s is not rent type", value))
}
