package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository"
	"github.com/geekcamp-vol11-team30/backend/util"
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
	cfg *config.Config
	er  repository.EventRepository
	// uv validator.UserValidator
}

func NewEventUsecase(cfg *config.Config, er repository.EventRepository) EventUsecase {
	return &eventUsecase{
		cfg: cfg,
		er:  er,
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
		return entity.Event{}, fmt.Errorf("error on create event: %w", err)
	}

	units := reqEvent.Units
	for i, u := range units {
		units[i].TimeSlot = u.TimeSlot.Round(time.Second)
		reqEvent.Units[i].EventID = newEvent.ID
	}
	units, err = eu.er.CreateEventTimeUnits(ctx, tx, units)
	if err != nil {
		return entity.Event{}, fmt.Errorf("error on create event time units: %w", err)
	}
	newEvent.Units = units

	err = tx.Commit()
	if err != nil {
		return entity.Event{}, fmt.Errorf("error on commit on create event: %w", err)
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
	for i := range ansUnits {
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

	go func() {
		log.Println("start goroutine", event)
		// メールの通知を希望するなら
		fmt.Println("aaaaaaaaaaaaaa", event.NotifyByEmail, event.ConfirmationEmail)
		if event.NotifyByEmail && event.ConfirmationEmail != "" {
			// ユーザーの回答数を数える
			userAnswerCount, err := eu.er.FetchUserAnswerCount(ctx, nil, eventId)
			fmt.Println("aaaaaaaaaaaaaa", userAnswerCount, event.NumberOfParticipants)
			if userAnswerCount == event.NumberOfParticipants {
				title := `[マジスケ]「` + event.Name + `」イベント参加者が集まりました！`
				idstr := util.ULIDToString(eventId)
				body := `マジスケをご利用頂き誠にありがとうございます。
回答者数が，予定人数の` + fmt.Sprintf("%d", event.NumberOfParticipants) +
					`人に到達しました。

https://magi-sche.net/detail/` + idstr + `から確認できます。
今後ともマジスケをよろしくお願いいたします。`
				util.SendMail(*eu.cfg, event.ConfirmationEmail, title, body)
				if err != nil {
					log.Printf("failed to send confirmation email: %v", err)
					// return entity.UserEventAnswer{}, apperror.NewUnknownError(fmt.Errorf("failed to send confirmation email: %w", err), nil)
				}
			}
		}
	}()

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
