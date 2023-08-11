package usecase

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EventUsecase interface {
	// イベント作成
	Create(ctx context.Context, event entity.Event, owner entity.User) (entity.Event, error)
	// イベント情報取得
	RetrieveEventAllData(ctx context.Context, eventId ulid.ULID) (entity.Event, error)
	// イベント回答登録
	CreateUserAnswer(ctx context.Context, eventId ulid.ULID, answer entity.UserEventAnswer, user entity.User) (entity.UserEventAnswer, error)
	// // イベント回答取得
	// RetrieveUserAnswer(ctx context.Context, eventId ulid.ULID, user entity.User) (entity.UserEventAnswer, error)
	// CreateAnonymousUser(ctx context.Context) (entity.User, error)
	// Register(ctx context.Context, user entity.User) (entity.UserResponse, error)
}

type eventUsecase struct {
	er repository.EventRepository
	// uv validator.UserValidator
}

func NewEventUsecase(er repository.EventRepository) EventUsecase {
	return &eventUsecase{
		er: er,
		// uv: uv,
	}
}

// Create implements EventUsecase.
func (eu *eventUsecase) Create(ctx context.Context, reqEvent entity.Event, owner entity.User) (entity.Event, error) {
	reqEvent.OwnerID = owner.ID

	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return entity.Event{}, err
	}
	defer tx.Rollback()

	newEvent, err := eu.er.CreateEvent(ctx, tx, reqEvent)
	if err != nil {
		return entity.Event{}, err
	}

	units := reqEvent.Units
	for i, u := range units {
		units[i].TimeSlot = u.TimeSlot.Round(time.Second)
		reqEvent.Units[i].EventID = newEvent.ID
	}
	units, err = eu.er.CreateEventTimeUnits(ctx, tx, units)
	if err != nil {
		return entity.Event{}, err
	}
	newEvent.Units = units

	err = tx.Commit()
	if err != nil {
		return entity.Event{}, err
	}
	return newEvent, nil
}

// RetrieveEventAllData implements EventUsecase.
func (eu *eventUsecase) RetrieveEventAllData(ctx context.Context, eventId ulid.ULID) (entity.Event, error) {
	// イベントが存在するか確認・取得
	event, err := eu.er.FetchEvent(ctx, nil, eventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Event{}, apperror.NewNotFoundError(err, "event not found")
		}
		return entity.Event{}, err
	}
	// イベント時間単位取得
	units, err := eu.er.FetchEventTimeUnits(ctx, nil, eventId)
	if err != nil {
		return entity.Event{}, err
	}
	event.Units = units

	answers, err := eu.er.FetchEventAnswersWithUnits(ctx, nil, eventId)
	if err != nil {
		return entity.Event{}, err
	}
	log.Println(answers)
	event.UserAnswers = answers
	return event, nil
}

// CreateUserAnswer implements EventUsecase.
func (eu *eventUsecase) CreateUserAnswer(ctx context.Context, eventId ulid.ULID, reqAnswer entity.UserEventAnswer, user entity.User) (entity.UserEventAnswer, error) {
	reqAnswer.EventID = eventId
	reqAnswer.UserID = user.ID
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return entity.UserEventAnswer{}, err
	}
	defer tx.Rollback()

	// イベントが存在するか確認・取得
	event, err := eu.er.FetchEvent(ctx, tx, eventId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.UserEventAnswer{}, apperror.NewNotFoundError(err, "event not found")
		}
		return entity.UserEventAnswer{}, err
	}
	// イベント時間単位取得
	units, err := eu.er.FetchEventTimeUnits(ctx, tx, eventId)
	if err != nil {
		return entity.UserEventAnswer{}, err
	}
	event.Units = units

	// eventのunitとanswerのunitが一致しているか確認(単純なvalidateができない)
	unitsMap := make(map[ulid.ULID]time.Time)
	for _, u := range event.Units {
		unitsMap[u.ID] = u.TimeSlot
	}
	for _, a := range reqAnswer.Units {
		if _, ok := unitsMap[a.EventTimeUnitID]; ok {
			delete(unitsMap, a.EventTimeUnitID)
		} else {
			return entity.UserEventAnswer{}, apperror.NewInvalidRequestBodyError(nil, "eventTimeUnitID is not valid")
		}
	}
	if len(unitsMap) != 0 {
		return entity.UserEventAnswer{}, apperror.NewInvalidRequestBodyError(nil, "eventTimeUnitID is not valid")
	}

	// イベント参加回答登録
	newAnswer, err := eu.er.UpdateEventAnswer(ctx, tx, reqAnswer)
	if err != nil {
		return entity.UserEventAnswer{}, err
	}
	ansUnits := reqAnswer.Units
	for i, _ := range ansUnits {
		ansUnits[i].UserEventAnswerID = newAnswer.ID
	}
	// イベント参加回答時間単位登録
	ansUnits, err = eu.er.RegisterAnswerUnits(ctx, tx, ansUnits)
	if err != nil {
		return entity.UserEventAnswer{}, err
	}
	newAnswer.Units = ansUnits

	// commit!
	err = tx.Commit()
	if err != nil {
		return entity.UserEventAnswer{}, err
	}

	return newAnswer, nil
}

// // RetrieveUserAnswer implements EventUsecase.
// func (eu *eventUsecase) RetrieveUserAnswer(ctx context.Context, eventId ulid.ULID, user entity.User) (entity.UserEventAnswer, error) {
// 	// reqAnswer.EventID = eventId
// 	// reqAnswer.UserID = user.ID
// 	tx, err := boil.BeginTx(ctx, nil)
// 	if err != nil {
// 		return entity.UserEventAnswer{}, err
// 	}
// 	defer tx.Rollback()

// 	uea, err := eu.er.FetchEventAnswer(ctx, tx, eventId, user.ID)
// 	if err != nil {
// 		return entity.UserEventAnswer{}, err
// 	}

// 	// // イベントが存在するか確認・取得
// 	// event, err := eu.er.FetchEvent(ctx, tx, eventId)
// 	// if err != nil {
// 	// 	if errors.Is(err, sql.ErrNoRows) {
// 	// 		return entity.UserEventAnswer{}, apperror.NewNotFoundError(err, "event not found")
// 	// 	}
// 	// 	return entity.UserEventAnswer{}, err
// 	// }
// 	// // イベント時間単位取得
// 	// units, err := eu.er.FetchEventTimeUnits(ctx, tx, eventId)
// 	// if err != nil {
// 	// 	return entity.UserEventAnswer{}, err
// 	// }
// 	// event.Units = units

// 	// // eventのunitとanswerのunitが一致しているか確認(単純なvalidateができない)
// 	// unitsMap := make(map[ulid.ULID]time.Time)
// 	// for _, u := range event.Units {
// 	// 	unitsMap[u.ID] = u.TimeSlot
// 	// }
// 	// for _, a := range reqAnswer.Units {
// 	// 	if _, ok := unitsMap[a.EventTimeUnitID]; ok {
// 	// 		delete(unitsMap, a.EventTimeUnitID)
// 	// 	} else {
// 	// 		return entity.UserEventAnswer{}, apperror.NewInvalidRequestBodyError(nil, "eventTimeUnitID is not valid")
// 	// 	}
// 	// }
// 	// if len(unitsMap) != 0 {
// 	// 	return entity.UserEventAnswer{}, apperror.NewInvalidRequestBodyError(nil, "eventTimeUnitID is not valid")
// 	// }

// 	// // イベント参加回答登録
// 	// newAnswer, err := eu.er.UpdateEventAnswer(ctx, tx, reqAnswer)
// 	// if err != nil {
// 	// 	return entity.UserEventAnswer{}, err
// 	// }
// 	// ansUnits := reqAnswer.Units
// 	// for i, _ := range ansUnits {
// 	// 	ansUnits[i].UserEventAnswerID = newAnswer.ID
// 	// }
// 	// // イベント参加回答時間単位登録
// 	// ansUnits, err = eu.er.RegisterAnswerUnits(ctx, tx, ansUnits)
// 	// if err != nil {
// 	// 	return entity.UserEventAnswer{}, err
// 	// }
// 	// newAnswer.Units = ansUnits

// 	// commit!
// 	err = tx.Commit()
// 	if err != nil {
// 		return entity.UserEventAnswer{}, err
// 	}

// 	return uea, nil
// }
