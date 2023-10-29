package input_validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/relvacode/iso8601"
)

var IsISO8601Date validator.Func = func(fl validator.FieldLevel) bool {
	_, err := iso8601.ParseString(fl.Field().String())
	return err == nil
}
