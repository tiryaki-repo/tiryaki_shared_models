package validator

import (
	valp "github.com/go-playground/validator/v10"
)

type validator struct {
	vp *valp.Validate
}

func NewValidator() *validator {
	return &validator{vp: valp.New()}
}

func (v *validator) Validate(s any) error {
	return v.vp.Struct(s)
}
