package constants

import (
	"errors"
	"fmt"
)

const (
	TransportTypeCar     = "CAR"
	TransportTypeBike    = "BIKE"
	TransportTypeScooter = "SCOOTER"
)

var transportTypes = []string{TransportTypeBike, TransportTypeCar, TransportTypeScooter}

func CheckTransportType(value string) error {
	for _, transportType := range transportTypes {
		if value == transportType {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("%s is not transport type", value))
}
