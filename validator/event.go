package validator

import (
	"github.com/geekcamp-vol11-team30/backend/entity"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IEventValidator interface {
	Validate(event entity.Event) error
}

type eventValidator struct{}

func NewEventValidator() IEventValidator {
	return &eventValidotor{}
}

func (ev *eventValidator) Validate(event entity.Event) error {
	return validation.ValidateStruct(&event,
		validation.Field(&event.Name,
			validation.Required,
		),
		validation.Field(&event.Description,
			validation.Required,
		),
		validation.Field(&event.DurationAbout,
			validation.Required,
		),
		validation.Field(&event.UnitSeconds,
			validation.Required,
		),
	)
}
