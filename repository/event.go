package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/db/models"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type EventRepository interface {
	// イベントを作成する
	CreateEvent(ctx context.Context, tx *sql.Tx, event entity.Event) (entity.Event, error)
	// イベントのタイムスロットを作成する
	CreateEventTimeUnits(ctx context.Context, tx *sql.Tx, event []entity.EventTimeUnit) ([]entity.EventTimeUnit, error)
	// イベントを取得する
	FetchEvent(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) (entity.Event, error)
	// イベントのタイムスロットを取得する
	FetchEventTimeUnits(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) ([]entity.EventTimeUnit, error)
	// イベントの全ユーザー回答(Unit付き)を取得する
	FetchEventAnswersWithUnits(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) ([]entity.UserEventAnswer, error)
	// イベントの指定ユーザー回答(Unit無し)を取得する
	FetchEventAnswer(ctx context.Context, tx *sql.Tx, eventId ulid.ULID, userId ulid.ULID) (entity.UserEventAnswer, error)

	// イベント参加回答更新
	UpdateEventAnswer(ctx context.Context, tx *sql.Tx, answer entity.UserEventAnswer) (entity.UserEventAnswer, error)
	// イベント参加回答時間単位を登録する
	RegisterAnswerUnits(ctx context.Context, tx *sql.Tx, answer []entity.UserEventAnswerUnit) ([]entity.UserEventAnswerUnit, error)

	// // イベントとイベント単位を取得する
	// FetchEventAndUnits(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) (entity.Event, error)
	// // イベントの全情報を取得する
	// FetchEventAllDataByID(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) (entity.Event, error)
	// // イベント参加可否を登録する
	// RegisterAnswerWithUnits(ctx context.Context, tx *sql.Tx, answer entity.UserEventAnswer) (entity.UserEventAnswer, error)
}

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepository{
		db: db,
	}
}

// CreateEvent implements EventRepository.
func (er *eventRepository) CreateEvent(ctx context.Context, tx *sql.Tx, event entity.Event) (entity.Event, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	id := util.GenerateULID(ctx)
	e := &models.Event{
		ID:            util.ULIDToString(id),
		OwnerID:       util.ULIDToString(event.OwnerID),
		Name:          event.Name,
		Description:   event.Description,
		DurationAbout: event.DurationAbout,
		UnitSeconds:   uint64(event.UnitSeconds),
	}
	err := e.Insert(ctx, exc, boil.Infer())
	if err != nil {
		return entity.Event{}, err
	}

	return entity.Event{
		ID:            id,
		OwnerID:       event.OwnerID,
		Name:          e.Name,
		Description:   e.Description,
		DurationAbout: e.DurationAbout,
		UnitSeconds:   int(e.UnitSeconds),
		// Units:         units,
		// UserAnswers:   []entity.UserEventAnswer{},
	}, nil
}

// CreateEventTimeUnits implements EventRepository.
func (er *eventRepository) CreateEventTimeUnits(ctx context.Context, tx *sql.Tx, units []entity.EventTimeUnit) ([]entity.EventTimeUnit, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	valueStrings := make([]string, 0, len(units))
	valueArgs := make([]interface{}, 0, len(units)*4)

	for i, unit := range units {
		unitId := util.GenerateULID(ctx)

		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, util.ULIDToString(unitId))
		valueArgs = append(valueArgs, util.ULIDToString(unit.EventID))
		valueArgs = append(valueArgs, unit.TimeSlot)
		valueArgs = append(valueArgs, uint64(unit.SlotSeconds))
		units[i].ID = unitId
	}

	query := fmt.Sprintf("INSERT INTO event_time_unit (id, event_id, time_slot, slot_seconds) VALUES %s", strings.Join(valueStrings, ","))
	// log.Println(query, valueStrings, valueArgs)

	_, err := exc.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		return []entity.EventTimeUnit{}, err
	}
	return units, nil
}

// FetchEvent implements EventRepository.
func (er *eventRepository) FetchEvent(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) (entity.Event, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	eventM, err := models.FindEvent(ctx, exc, util.ULIDToString(eventId))
	if err != nil {
		return entity.Event{}, err
	}
	oid, err := util.ULIDFromString(eventM.OwnerID)
	if err != nil {
		return entity.Event{}, err
	}
	return entity.Event{
		ID:            eventId,
		OwnerID:       oid,
		Name:          eventM.Name,
		Description:   eventM.Description,
		DurationAbout: eventM.DurationAbout,
		UnitSeconds:   int(eventM.UnitSeconds),
		// Units:         etus,
	}, nil
}

// FetchEventTimeUnits implements EventRepository.
func (er *eventRepository) FetchEventTimeUnits(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) ([]entity.EventTimeUnit, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}
	etusm, err := models.EventTimeUnits(
		models.EventTimeUnitWhere.EventID.EQ(util.ULIDToString(eventId)),
		qm.OrderBy("time_slot"),
	).All(ctx, exc)
	if err != nil {
		return []entity.EventTimeUnit{}, err
	}
	etus := make([]entity.EventTimeUnit, len(etusm))
	for i, etu := range etusm {
		etuid, err := util.ULIDFromString(etu.ID)
		if err != nil {
			return []entity.EventTimeUnit{}, err
		}
		etus[i] = entity.EventTimeUnit{
			ID:       etuid,
			EventID:  eventId,
			TimeSlot: etu.TimeSlot,
		}
	}
	return etus, nil
}

// FetchEventAnswersWithUnits implements EventRepository.
func (er *eventRepository) FetchEventAnswersWithUnits(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) ([]entity.UserEventAnswer, error) {
	// panic("unimplemented")
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	answersm, err := models.UserEventAnswers(
		qm.Load(
			models.UserEventAnswerRels.UserEventAnswerUnits,
			// join EventTimeUnit
		),
		models.UserEventAnswerWhere.EventID.EQ(util.ULIDToString(eventId)),
		// order by created at
		qm.OrderBy("created_at"),
	).All(ctx, exc)
	if err != nil {
		return []entity.UserEventAnswer{}, err
	}
	answers := make([]entity.UserEventAnswer, len(answersm))
	for i, answerm := range answersm {
		id, _ := util.ULIDFromString(answerm.ID)
		userId, _ := util.ULIDFromString(answerm.UserID)
		units := make([]entity.UserEventAnswerUnit, len(answerm.R.UserEventAnswerUnits))
		for i, unitm := range answerm.R.UserEventAnswerUnits {
			unitId, _ := util.ULIDFromString(unitm.ID)
			etuId, _ := util.ULIDFromString(unitm.EventTimeUnitID)
			units[i] = entity.UserEventAnswerUnit{
				ID:                unitId,
				UserEventAnswerID: userId,
				EventTimeUnitID:   etuId,
				Availability:      entity.Availability(unitm.Availability),
			}
		}
		answer := entity.UserEventAnswer{
			ID:           id,
			UserID:       userId,
			EventID:      eventId,
			UserNickname: answerm.UserNickname,
			Note:         answerm.Note,
			Units:        units,
			// Units:        []entity.UserEventAnswerUnit{},
		}
		answers[i] = answer
	}
	log.Println("!!!!!!!!!!!!!!!!!!!!!!", answers, answersm)
	return answers, nil
}

// FetchEventAnswer implements EventRepository.
func (er *eventRepository) FetchEventAnswer(ctx context.Context, tx *sql.Tx, eventId ulid.ULID, userId ulid.ULID) (entity.UserEventAnswer, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	ueam, err := models.UserEventAnswers(
		qm.Where("event_id = ? AND user_id = ?", util.ULIDToString(eventId), util.ULIDToString(userId)),
		qm.Load(
			models.UserEventAnswerRels.UserEventAnswerUnits,
			qm.OrderBy("created_at"),
		),
		models.UserEventAnswerWhere.EventID.EQ(util.ULIDToString(eventId)),
		qm.OrderBy("created_at"),
	).One(ctx, exc)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.UserEventAnswer{}, apperror.NewNotFoundError(err, nil)
		}
		return entity.UserEventAnswer{}, err
	}
	ueaId, _ := util.ULIDFromString(ueam.ID)
	units := make([]entity.UserEventAnswerUnit, len(ueam.R.UserEventAnswerUnits))
	for i, unitm := range ueam.R.UserEventAnswerUnits {
		unitId, _ := util.ULIDFromString(unitm.ID)
		etuId, _ := util.ULIDFromString(unitm.EventTimeUnitID)
		units[i] = entity.UserEventAnswerUnit{
			ID:                unitId,
			UserEventAnswerID: userId,
			EventTimeUnitID:   etuId,
			Availability:      entity.Availability(unitm.Availability),
		}
	}
	return entity.UserEventAnswer{
		ID:           ueaId,
		UserID:       userId,
		EventID:      eventId,
		UserNickname: ueam.UserNickname,
		Note:         ueam.Note,
		Units:        units,
		// Units:        []entity.UserEventAnswerUnit{},
	}, nil
}

// UpdateEventAnswer implements EventRepository.
func (er *eventRepository) UpdateEventAnswer(ctx context.Context, tx *sql.Tx, answer entity.UserEventAnswer) (entity.UserEventAnswer, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	_, err := models.UserEventAnswers(
		// qm.Where("user_id = ?", util.ULIDToString(answer.UserID)),
		models.UserEventAnswerWhere.UserID.EQ(util.ULIDToString(answer.UserID)),
		// qm.And("event_id = ?", util.ULIDToString(answer.EventID)),
		models.UserEventAnswerWhere.EventID.EQ(util.ULIDToString(answer.EventID)),
	).DeleteAll(ctx, exc)
	if err != nil {
		return entity.UserEventAnswer{}, err
	}

	aid := util.GenerateULID(ctx)
	eam := &models.UserEventAnswer{
		ID:           util.ULIDToString(aid),
		UserID:       util.ULIDToString(answer.UserID),
		EventID:      util.ULIDToString(answer.EventID),
		UserNickname: answer.UserNickname,
		Note:         answer.Note,
	}
	err = eam.Insert(ctx, exc, boil.Infer())
	if err != nil {
		return entity.UserEventAnswer{}, err
	}
	return entity.UserEventAnswer{
		ID:           aid,
		UserID:       answer.UserID,
		EventID:      answer.EventID,
		UserNickname: answer.UserNickname,
		Note:         answer.Note,
	}, nil
}

// RegisterAnswerUnits implements EventRepository.
func (er *eventRepository) RegisterAnswerUnits(ctx context.Context, tx *sql.Tx, units []entity.UserEventAnswerUnit) ([]entity.UserEventAnswerUnit, error) {
	var exc boil.ContextExecutor = tx
	if tx == nil {
		exc = er.db
	}

	valueStrings := make([]string, 0, len(units))
	valueArgs := make([]interface{}, 0, len(units)*4)
	// units := answer.Units
	// newUnits := units

	for i, unit := range units {
		unitId := util.GenerateULID(ctx)
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
		valueArgs = append(valueArgs, util.ULIDToString(unitId))
		valueArgs = append(valueArgs, util.ULIDToString(unit.UserEventAnswerID))
		valueArgs = append(valueArgs, util.ULIDToString(unit.EventTimeUnitID))
		valueArgs = append(valueArgs, unit.Availability)
		units[i].ID = unitId
	}

	query := fmt.Sprintf("INSERT INTO user_event_answer_unit (id, user_event_answer_id, event_time_unit_id, availability) VALUES %s", strings.Join(valueStrings, ","))
	// log.Println(query, valueStrings, valueArgs)

	_, err := exc.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		return []entity.UserEventAnswerUnit{}, err
	}
	return units, nil
}

// // RegisterAnswerWithUnits implements EventRepository.
// func (er *eventRepository) RegisterAnswerWithUnits(ctx context.Context, tx *sql.Tx, answer entity.UserEventAnswer) (entity.UserEventAnswer, error) {
// 	var exc boil.ContextExecutor = tx
// 	if tx == nil {
// 		exc = er.db
// 	}

// 	_, err := models.UserEventAnswers(
// 		qm.Where("user_id = ?", util.ULIDToString(answer.UserID)),
// 		qm.And("event_id = ?", util.ULIDToString(answer.EventID)),
// 	).DeleteAll(ctx, exc)
// 	if err != nil {
// 		return entity.UserEventAnswer{}, err
// 	}

// 	aid:= util.GenerateULID(ctx)
// 	eam := &models.UserEventAnswer{
// 		ID:           util.ULIDToString(aid),
// 		UserID:       util.ULIDToString(answer.UserID),
// 		EventID:      util.ULIDToString(answer.EventID),
// 		UserNickname: answer.UserNickname,
// 		Note:         answer.Note,
// 	}
// 	err = eam.Insert(ctx, exc, boil.Infer())
// 	if err != nil {
// 		return entity.UserEventAnswer{}, err
// 	}

// 	models.UserEventAnswerUnits(
// 		qm.Where("user_event_answer_id = ?", eam.ID),
// 	).DeleteAll(ctx, exc)

// 	valueStrings := make([]string, 0, len(answer.Units))
// 	valueArgs := make([]interface{}, 0, len(answer.Units)*4)
// 	units := answer.Units

// 	for i, unit := range answer.Units {
// 		unitId := util.GenerateULID(ctx)
// 		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
// 		valueArgs = append(valueArgs, util.ULIDToString(unitId))
// 		valueArgs = append(valueArgs, util.ULIDToString(aid))
// 		valueArgs = append(valueArgs, util.ULIDToString(unit.EventTimeUnitID))
// 		valueArgs = append(valueArgs, unit.Availability)
// 		units[i].ID = unitId
// 		units[i].UserEventAnswerID = aid
// 	}

// 	query := fmt.Sprintf("INSERT INTO user_event_answer_unit (id, user_event_answer_id, event_time_unit_id, availability) VALUES %s", strings.Join(valueStrings, ","))
// 	// log.Println(query, valueStrings, valueArgs)

// 	_, err = exc.ExecContext(ctx, query, valueArgs...)
// 	if err != nil {
// 		return entity.UserEventAnswer{}, err
// 	}

// 	return entity.UserEventAnswer{
// 		ID:           aid,
// 		UserID:       answer.UserID,
// 		EventID:      answer.EventID,
// 		UserNickname: answer.UserNickname,
// 		Note:         answer.Note,
// 		Units:        units,
// 	}, nil
// }

// // CreateEventWithSlots implements EventRepository.
// func (er *eventRepository) CreateEventWithSlots(ctx context.Context, tx *sql.Tx, event entity.Event) (entity.Event, error) {
// 	var exc boil.ContextExecutor = tx
// 	if tx == nil {
// 		exc = er.db
// 	}
// 	// tx, err := boil.BeginTx(ctx, nil)
// 	// if err != nil {
// 	// 	return entity.Event{}, err
// 	// }
// 	// defer tx.Rollback()
// 	id, err := util.GenerateULID(ctx)
// 	if err != nil {
// 		return entity.Event{}, err
// 	}
// 	e := &models.Event{
// 		ID:            util.ULIDToString(id),
// 		OwnerID:       util.ULIDToString(event.OwnerID),
// 		Name:          event.Name,
// 		Description:   event.Description,
// 		DurationAbout: event.DurationAbout,
// 		UnitSeconds:   uint64(event.UnitSeconds),
// 	}
// 	err = e.Insert(ctx, exc, boil.Infer())
// 	if err != nil {
// 		return entity.Event{}, err
// 	}

// 	valueStrings := make([]string, 0, len(event.Units))
// 	valueArgs := make([]interface{}, 0, len(event.Units)*4)
// 	units := event.Units

// 	for i, unit := range event.Units {
// 		unitId, err := util.GenerateULID(ctx)
// 		if err != nil {
// 			return entity.Event{}, err
// 		}
// 		valueStrings = append(valueStrings, "(?, ?, ?, ?)")
// 		valueArgs = append(valueArgs, util.ULIDToString(unitId))
// 		valueArgs = append(valueArgs, e.ID)
// 		valueArgs = append(valueArgs, unit.TimeSlot)
// 		valueArgs = append(valueArgs, uint64(unit.SlotSeconds))
// 		units[i].ID = unitId
// 		units[i].EventID = id
// 	}

// 	query := fmt.Sprintf("INSERT INTO event_time_unit (id, event_id, time_slot, slot_seconds) VALUES %s", strings.Join(valueStrings, ","))
// 	// log.Println(query, valueStrings, valueArgs)

// 	_, err = exc.ExecContext(ctx, query, valueArgs...)
// 	if err != nil {
// 		return entity.Event{}, err
// 	}

// 	return entity.Event{
// 		ID:            id,
// 		OwnerID:       event.OwnerID,
// 		Name:          e.Name,
// 		Description:   e.Description,
// 		DurationAbout: e.DurationAbout,
// 		UnitSeconds:   int(e.UnitSeconds),
// 		Units:         units,
// 		UserAnswers:   []entity.UserEventAnswer{},
// 	}, nil
// }

// // FetchEventAndUnits implements EventRepository.
// func (er *eventRepository) FetchEventAndUnits(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) (entity.Event, error) {
// 	var exc boil.ContextExecutor = tx
// 	if tx == nil {
// 		exc = er.db
// 	}

// 	eventM, err := models.FindEvent(ctx, exc, util.ULIDToString(eventId))
// 	if err != nil {
// 		return entity.Event{}, err
// 	}
// 	// order by time_slot
// 	etusm, err := eventM.EventTimeUnits(
// 		qm.OrderBy("time_slot"),
// 	).All(ctx, exc)
// 	if err != nil {
// 		return entity.Event{}, err
// 	}
// 	etus := make([]entity.EventTimeUnit, len(etusm))
// 	for i, etu := range etusm {
// 		etuid, err := util.ULIDFromString(etu.ID)
// 		if err != nil {
// 			return entity.Event{}, err
// 		}
// 		etus[i] = entity.EventTimeUnit{
// 			ID:          etuid,
// 			EventID:     eventId,
// 			TimeSlot:    etu.TimeSlot,
// 			SlotSeconds: int(etu.SlotSeconds),
// 		}
// 	}

// 	oid, err := util.ULIDFromString(eventM.OwnerID)
// 	if err != nil {
// 		return entity.Event{}, err
// 	}
// 	return entity.Event{
// 		ID:            eventId,
// 		OwnerID:       oid,
// 		Name:          eventM.Name,
// 		Description:   eventM.Description,
// 		DurationAbout: eventM.DurationAbout,
// 		UnitSeconds:   int(eventM.UnitSeconds),
// 		Units:         etus,
// 	}, nil
// }
// // FetchEventAllDataByID implements EventRepository.
// func (*eventRepository) FetchEventAllDataByID(ctx context.Context, tx *sql.Tx, eventId ulid.ULID) (entity.Event, error) {
// 	panic("unimplemented")
// }
