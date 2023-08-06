package validator

import (
	"github.com/geekcamp-vol11-team30/backend/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserValidator interface {
	Validate(user entity.User) error
}

type userValidator struct {
}

func NewUserValidator() UserValidator {
	return &userValidator{}
}

// Validate implements UserValidator.
func (*userValidator) Validate(user entity.User) error {
	// panic("unimplemented")
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name,
			// validation.Required,
			validation.RuneLength(0, 255),
		),
		// validation.Field(&user.SlackID,
		// 	validation.Required,
		// 	validation.RuneLength(1, 255),
		// ),
	)
}
