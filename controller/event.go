//go:generate mockgen -source=./event.go -destination=./mock/event.go -package=mockcontroller
package controller

import (
	"log"
	"net/http"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
)

type EventController interface {
	Create(c echo.Context) error
	Retrieve(c echo.Context) error
	Attend(c echo.Context) error
}

type eventController struct {
	eu usecase.EventUsecase
}

func NewEventController(eu usecase.EventUsecase) EventController {
	return &eventController{
		eu: eu,
	}
}

// Register implements UserController.
func (ec *eventController) Create(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := appcontext.Extract(ctx).GetUser()
	if err != nil {
		return err
	}

	er := entity.EventRequest{}
	if err := c.Bind(&er); err != nil {
		return apperror.NewInvalidRequestBodyError(err, nil)
	}
	event := eventRequestToEvent(er)
	// event.OwnerID = user.ID
	event, err = ec.eu.Create(ctx, event, user)
	if err != nil {
		if ae, ok := err.(*apperror.AppError); ok {
			return ae
		} else {
			return apperror.NewUnknownError(err, nil)
		}
	}
	res := eventToEventResponse(event)

	c.Response().Header().Set("Location", "/events/"+res.ID)
	return c.JSON(http.StatusCreated, res)
}

// Retrieve implements EventController.
func (ec *eventController) Retrieve(c echo.Context) error {
	eventIdStr := c.Param("event_id")
	eventId, err := util.ULIDFromString(eventIdStr)
	if err != nil {
		return apperror.NewInvalidRequestPathError(err, nil)
	}
	ctx := c.Request().Context()
	// user, err := appcontext.Extract(ctx).GetUser()
	// if err != nil {
	// 	return err
	// }
	event, err := ec.eu.RetrieveEventAllData(ctx, eventId)
	if err != nil {
		return err
	}
	res := eventToEventResponse(event)
	return c.JSON(http.StatusOK, res)
}

// Attend implements EventController.
func (ec *eventController) Attend(c echo.Context) error {
	eventIdStr := c.Param("event_id")
	eventId, err := util.ULIDFromString(eventIdStr)
	if err != nil {
		return apperror.NewInvalidRequestPathError(err, nil)
	}

	ctx := c.Request().Context()
	user, err := appcontext.Extract(ctx).GetUser()
	if err != nil {
		return err
	}

	uear := entity.UserEventAnswerRequest{}
	if err := c.Bind(&uear); err != nil {
		return apperror.NewInvalidRequestBodyError(err, nil)
	}
	uea, err := ueaRequestToUea(uear)
	if err != nil {
		return err
	}
	log.Printf("uear: %+v uea: %+v\n", uear, uea)

	uea, err = ec.eu.Attend(ctx, eventId, uea, user)
	if err != nil {
		if ae, ok := err.(*apperror.AppError); ok {
			return ae
		} else {
			return apperror.NewUnknownError(err, nil)
		}
	}
	res := ueaToUeaResponse(uea)

	c.Response().Header().Set("Location", "/events/"+eventIdStr+"/attend")
	return c.JSON(http.StatusCreated, res)
}

func eventRequestToEvent(er entity.EventRequest) entity.Event {
	eur := er.Units
	units := make([]entity.EventTimeUnit, len(eur))
	for i, u := range eur {
		units[i] = entity.EventTimeUnit{
			TimeSlot:    u.TimeSlot,
			SlotSeconds: u.SlotSeconds,
		}
	}
	return entity.Event{
		Name:          er.Name,
		Description:   er.Description,
		DurationAbout: er.DurationAbout,
		UnitSeconds:   er.UnitSeconds,
		Units:         units,
	}
}

func eventToEventResponse(e entity.Event) entity.EventResponse {
	ers := make([]entity.EventTimeUnitResponse, len(e.Units))
	for i, u := range e.Units {
		ers[i] = entity.EventTimeUnitResponse{
			ID:          util.ULIDToString(u.ID),
			TimeSlot:    u.TimeSlot,
			SlotSeconds: u.SlotSeconds,
		}
	}
	eas := make([]entity.UserEventAnswerResponse, len(e.UserAnswers))
	for i, u := range e.UserAnswers {
		eas[i] = ueaToUeaResponse(u)
	}
	return entity.EventResponse{
		ID:            util.ULIDToString(e.ID),
		OwnerID:       util.ULIDToString(e.OwnerID),
		Name:          e.Name,
		Description:   e.Description,
		DurationAbout: e.DurationAbout,
		UnitSeconds:   e.UnitSeconds,
		Units:         ers,
		UserAnswers:   eas,
	}
}

func ueaRequestToUea(uear entity.UserEventAnswerRequest) (entity.UserEventAnswer, error) {
	ueau := uear.Units
	units := make([]entity.UserEventAnswerUnit, len(ueau))
	for i, u := range ueau {
		id, err := util.ULIDFromString(u.EventTimeUnitID)
		if err != nil {
			return entity.UserEventAnswer{}, apperror.NewInvalidRequestBodyError(err, "eventTimeUnitIdの型が不正です")
		}
		units[i] = entity.UserEventAnswerUnit{
			EventTimeUnitID: id,
			Availability:    u.Availability,
		}
	}
	return entity.UserEventAnswer{
		UserNickname: uear.UserNickname,
		Note:         uear.Note,
		Units:        units,
	}, nil
}

func ueaToUeaResponse(uea entity.UserEventAnswer) entity.UserEventAnswerResponse {
	uears := make([]entity.UserEventAnswerUnitResponse, len(uea.Units))
	for i, u := range uea.Units {
		uears[i] = entity.UserEventAnswerUnitResponse{
			EventTimeUnitID: util.ULIDToString(u.EventTimeUnitID),
			Availability:    u.Availability,
		}
	}
	return entity.UserEventAnswerResponse{
		ID:           util.ULIDToString(uea.ID),
		UserID:       util.ULIDToString(uea.UserID),
		UserNickname: uea.UserNickname,
		Note:         uea.Note,
		Units:        uears,
	}
}
