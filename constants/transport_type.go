package constants

import (
	"errors"
	"fmt"
	"strings"
)

const (
	TransportTypeCar     = "CAR"
	TransportTypeBike    = "BIKE"
	TransportTypeScooter = "SCOOTER"
	TransportTypeAll     = "ALL"
)

var transportTypes = []string{TransportTypeBike, TransportTypeCar, TransportTypeScooter}

func CheckTransportType(value string) error {
	for _, transportType := range transportTypes {
		if strings.ToUpper(value) == transportType {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("%s is not transport type", value))
}

func CheckTransportTypeWithAll(value string) error {
	if strings.ToUpper(value) == TransportTypeAll {

		return nil
	}
	for _, transportType := range transportTypes {
		if value == transportType {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("%s is not transport type", value))
}
