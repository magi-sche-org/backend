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
	return &eventValidator{}
}

func (ev *eventValidator) Validate(event entity.Event) error {
	return validation.ValidateStruct(&event,
		validation.Field(&event.Name,
			validation.Required.Error("タイトルは必須です"),
			validation.Length(1, 100).Error("タイトルは100文字以内です"),
		),
		validation.Field(&event.Description,
			validation.Required.Error("説明は必須です"),
			validation.Length(0, 1000).Error("説明は1000文字以内です"),
		),
		validation.Field(&event.DurationAbout,
			validation.Required.Error("開催時間は必須です"),
			validation.Length(1, 100).Error("開催時間は100文字以内です"),
		),
		validation.Field(&event.UnitSeconds,
			validation.Required.Error("時間単位は30分,1時間,1日です"),
			//validation.Min(30).Error("時間単位は30分,1時間,1日です"),
		),
	)
}
