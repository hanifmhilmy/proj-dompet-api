package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var (
	ErrValidator = errors.New("Validator")
)

func Validate(data interface{}) (bool, error) {
	v := validator.New()
	err := v.Struct(data)
	if err != nil {
		err = errors.Wrap(err, ErrValidator.Error())
		return false, err
	}
	return true, nil
}
