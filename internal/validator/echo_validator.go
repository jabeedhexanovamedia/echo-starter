package validator

import "github.com/go-playground/validator/v10"

type EchoValidator struct {
	validator *validator.Validate
}

func NewEchoValidator(v *validator.Validate) *EchoValidator {
	return &EchoValidator{validator: v}
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

//* USAGE: in any handler
// if err := c.Validate(&req); err != nil {
// 	return c.JSON(400, err)
// }
