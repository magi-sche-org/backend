//go:generate mockgen -source=./event.go -destination=./mock/event.go -package=mockcontroller
package controller

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type EventController interface {
	Create(c echo.Context) error
	Retrieve(c echo.Context) error
	CreateAnswer(c echo.Context) error
	RetrieveUserAnswer(c echo.Context) error
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
	event, err = ec.eu.Create(ctx, event, user)
	if err != nil {
		if ae, ok := err.(*apperror.AppError); ok {
			return ae
		} else {
			return apperror.NewUnknownError(fmt.Errorf("unknown error on create event controller: %w", err), nil)
		}
	}
	res := eventToEventResponse(event, user)

	c.Response().Header().Set("Location", "/events/"+res.ID)
	return util.JSONResponse(c, http.StatusCreated, res)
	// return c.JSON(http.StatusCreated, res)
}

// Retrieve implements EventController.
func (ec *eventController) Retrieve(c echo.Context) error {
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
	event, err := ec.eu.RetrieveEventAllData(ctx, eventId)
	if err != nil {
		return err
	}
	res := eventToEventResponse(event, user)
	return util.JSONResponse(c, http.StatusOK, res)
}

// CreateAnswer implements EventController.
func (ec *eventController) CreateAnswer(c echo.Context) error {
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

	_, err = ec.eu.CreateUserAnswer(ctx, eventId, uea, user)
	if err != nil {
		if ae, ok := err.(*apperror.AppError); ok {
			return ae
		} else {
			return apperror.NewUnknownError(err, nil)
		}
	}
	// ここから再取得，本来はUnitだけ得て混ぜたいが，加工がめんどうなのでひとまず…
	event, err := ec.eu.RetrieveEventAllData(ctx, eventId)
	if err != nil {
		return err
	}
	eventRes := eventToEventResponse(event, user)
	i := slices.IndexFunc(eventRes.UserAnswers, func(answer entity.UserEventAnswerResponse) bool {
		return answer.IsYourAnswer
	})
	res := eventRes.UserAnswers[i]
	c.Response().Header().Set("Location", "/events/"+eventIdStr+"/user/answer")
	return util.JSONResponse(c, http.StatusCreated, res)
}

// RetrieveUserAnswer implements EventController.
func (ec *eventController) RetrieveUserAnswer(c echo.Context) error {
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
	// とりあえずイベントの情報を全取得して，絞ってあげる感じで…
	event, err := ec.eu.RetrieveEventAllData(ctx, eventId)
	if err != nil {
		return err
	}
	eventRes := eventToEventResponse(event, user)
	i := slices.IndexFunc(eventRes.UserAnswers, func(answer entity.UserEventAnswerResponse) bool {
		return answer.IsYourAnswer
	})
	return util.JSONResponse(c, http.StatusOK, eventRes.UserAnswers[i])

}

// 受け取ったEventRequestを，Eventに変換する
func eventRequestToEvent(er entity.EventRequest) entity.Event {
	eur := er.Units
	units := make([]entity.EventTimeUnit, len(eur))
	for i, u := range eur {
		units[i] = entity.EventTimeUnit{
			TimeSlot: u.TimeSlot,
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

// イベントの情報を，レスポンスに変換する
func eventToEventResponse(e entity.Event, user entity.User) entity.EventResponse {
	ers := make([]entity.EventTimeUnitResponse, len(e.Units))
	unitsMap := make(map[ulid.ULID]entity.EventTimeUnitResponse)
	for i, u := range e.Units {
		ers[i] = entity.EventTimeUnitResponse{
			ID:       util.ULIDToString(u.ID),
			StartsAt: u.TimeSlot,
			// SlotSeconds: u.SlotSeconds,
			EndsAt: u.TimeSlot.Add(time.Duration(e.UnitSeconds) * time.Second),
		}
		unitsMap[u.ID] = ers[i]
	}
	eas := make([]entity.UserEventAnswerResponse, len(e.UserAnswers))
	for i, u := range e.UserAnswers {
		eas[i] = ueaToUeaResponse(u, user, unitsMap)
	}
	slices.SortFunc(ers, func(a, b entity.EventTimeUnitResponse) int {
		return cmp.Compare(a.StartsAt.UnixNano(), b.StartsAt.UnixNano())
	})
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

// 受け取ったイベント参加の回答を，EventAnswerに変換する
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

// EventAnswerを，レスポンス形式に変換する。該当ユーザーかそうでないかを含む。
func ueaToUeaResponse(uea entity.UserEventAnswer, user entity.User, unitsMap map[ulid.ULID]entity.EventTimeUnitResponse) entity.UserEventAnswerResponse {
	uears := make([]entity.UserEventAnswerUnitResponse, len(uea.Units))
	for i, u := range uea.Units {
		uears[i] = entity.UserEventAnswerUnitResponse{
			EventTimeUnitID: util.ULIDToString(u.EventTimeUnitID),
			Availability:    u.Availability,
			StartsAt:        unitsMap[u.EventTimeUnitID].StartsAt,
			EndsAt:          unitsMap[u.EventTimeUnitID].EndsAt,
		}
	}
	slices.SortFunc(uears, func(a, b entity.UserEventAnswerUnitResponse) int {
		return cmp.Compare(a.StartsAt.UnixNano(), b.StartsAt.UnixNano())
	})
	return entity.UserEventAnswerResponse{
		ID:           util.ULIDToString(uea.ID),
		UserID:       util.ULIDToString(uea.UserID),
		UserNickname: uea.UserNickname,
		Note:         uea.Note,
		Units:        uears,
		IsYourAnswer: uea.UserID == user.ID,
	}
}
