// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// EventTimeUnit is an object representing the database table.
type EventTimeUnit struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	EventID   string    `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	TimeSlot  time.Time `boil:"time_slot" json:"time_slot" toml:"time_slot" yaml:"time_slot"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *eventTimeUnitR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L eventTimeUnitL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EventTimeUnitColumns = struct {
	ID        string
	EventID   string
	TimeSlot  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	EventID:   "event_id",
	TimeSlot:  "time_slot",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var EventTimeUnitTableColumns = struct {
	ID        string
	EventID   string
	TimeSlot  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "event_time_unit.id",
	EventID:   "event_time_unit.event_id",
	TimeSlot:  "event_time_unit.time_slot",
	CreatedAt: "event_time_unit.created_at",
	UpdatedAt: "event_time_unit.updated_at",
}

// Generated where

var EventTimeUnitWhere = struct {
	ID        whereHelperstring
	EventID   whereHelperstring
	TimeSlot  whereHelpertime_Time
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperstring{field: "`event_time_unit`.`id`"},
	EventID:   whereHelperstring{field: "`event_time_unit`.`event_id`"},
	TimeSlot:  whereHelpertime_Time{field: "`event_time_unit`.`time_slot`"},
	CreatedAt: whereHelpertime_Time{field: "`event_time_unit`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`event_time_unit`.`updated_at`"},
}

// EventTimeUnitRels is where relationship names are stored.
var EventTimeUnitRels = struct {
	Event                string
	UserEventAnswerUnits string
}{
	Event:                "Event",
	UserEventAnswerUnits: "UserEventAnswerUnits",
}

// eventTimeUnitR is where relationships are stored.
type eventTimeUnitR struct {
	Event                *Event                   `boil:"Event" json:"Event" toml:"Event" yaml:"Event"`
	UserEventAnswerUnits UserEventAnswerUnitSlice `boil:"UserEventAnswerUnits" json:"UserEventAnswerUnits" toml:"UserEventAnswerUnits" yaml:"UserEventAnswerUnits"`
}

// NewStruct creates a new relationship struct
func (*eventTimeUnitR) NewStruct() *eventTimeUnitR {
	return &eventTimeUnitR{}
}

func (r *eventTimeUnitR) GetEvent() *Event {
	if r == nil {
		return nil
	}
	return r.Event
}

func (r *eventTimeUnitR) GetUserEventAnswerUnits() UserEventAnswerUnitSlice {
	if r == nil {
		return nil
	}
	return r.UserEventAnswerUnits
}

// eventTimeUnitL is where Load methods for each relationship are stored.
type eventTimeUnitL struct{}

var (
	eventTimeUnitAllColumns            = []string{"id", "event_id", "time_slot", "created_at", "updated_at"}
	eventTimeUnitColumnsWithoutDefault = []string{"id", "event_id", "time_slot"}
	eventTimeUnitColumnsWithDefault    = []string{"created_at", "updated_at"}
	eventTimeUnitPrimaryKeyColumns     = []string{"id"}
	eventTimeUnitGeneratedColumns      = []string{}
)

type (
	// EventTimeUnitSlice is an alias for a slice of pointers to EventTimeUnit.
	// This should almost always be used instead of []EventTimeUnit.
	EventTimeUnitSlice []*EventTimeUnit
	// EventTimeUnitHook is the signature for custom EventTimeUnit hook methods
	EventTimeUnitHook func(context.Context, boil.ContextExecutor, *EventTimeUnit) error

	eventTimeUnitQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	eventTimeUnitType                 = reflect.TypeOf(&EventTimeUnit{})
	eventTimeUnitMapping              = queries.MakeStructMapping(eventTimeUnitType)
	eventTimeUnitPrimaryKeyMapping, _ = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, eventTimeUnitPrimaryKeyColumns)
	eventTimeUnitInsertCacheMut       sync.RWMutex
	eventTimeUnitInsertCache          = make(map[string]insertCache)
	eventTimeUnitUpdateCacheMut       sync.RWMutex
	eventTimeUnitUpdateCache          = make(map[string]updateCache)
	eventTimeUnitUpsertCacheMut       sync.RWMutex
	eventTimeUnitUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var eventTimeUnitAfterSelectHooks []EventTimeUnitHook

var eventTimeUnitBeforeInsertHooks []EventTimeUnitHook
var eventTimeUnitAfterInsertHooks []EventTimeUnitHook

var eventTimeUnitBeforeUpdateHooks []EventTimeUnitHook
var eventTimeUnitAfterUpdateHooks []EventTimeUnitHook

var eventTimeUnitBeforeDeleteHooks []EventTimeUnitHook
var eventTimeUnitAfterDeleteHooks []EventTimeUnitHook

var eventTimeUnitBeforeUpsertHooks []EventTimeUnitHook
var eventTimeUnitAfterUpsertHooks []EventTimeUnitHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EventTimeUnit) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EventTimeUnit) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EventTimeUnit) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EventTimeUnit) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EventTimeUnit) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EventTimeUnit) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EventTimeUnit) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EventTimeUnit) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EventTimeUnit) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range eventTimeUnitAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEventTimeUnitHook registers your hook function for all future operations.
func AddEventTimeUnitHook(hookPoint boil.HookPoint, eventTimeUnitHook EventTimeUnitHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		eventTimeUnitAfterSelectHooks = append(eventTimeUnitAfterSelectHooks, eventTimeUnitHook)
	case boil.BeforeInsertHook:
		eventTimeUnitBeforeInsertHooks = append(eventTimeUnitBeforeInsertHooks, eventTimeUnitHook)
	case boil.AfterInsertHook:
		eventTimeUnitAfterInsertHooks = append(eventTimeUnitAfterInsertHooks, eventTimeUnitHook)
	case boil.BeforeUpdateHook:
		eventTimeUnitBeforeUpdateHooks = append(eventTimeUnitBeforeUpdateHooks, eventTimeUnitHook)
	case boil.AfterUpdateHook:
		eventTimeUnitAfterUpdateHooks = append(eventTimeUnitAfterUpdateHooks, eventTimeUnitHook)
	case boil.BeforeDeleteHook:
		eventTimeUnitBeforeDeleteHooks = append(eventTimeUnitBeforeDeleteHooks, eventTimeUnitHook)
	case boil.AfterDeleteHook:
		eventTimeUnitAfterDeleteHooks = append(eventTimeUnitAfterDeleteHooks, eventTimeUnitHook)
	case boil.BeforeUpsertHook:
		eventTimeUnitBeforeUpsertHooks = append(eventTimeUnitBeforeUpsertHooks, eventTimeUnitHook)
	case boil.AfterUpsertHook:
		eventTimeUnitAfterUpsertHooks = append(eventTimeUnitAfterUpsertHooks, eventTimeUnitHook)
	}
}

// One returns a single eventTimeUnit record from the query.
func (q eventTimeUnitQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EventTimeUnit, error) {
	o := &EventTimeUnit{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for event_time_unit")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EventTimeUnit records from the query.
func (q eventTimeUnitQuery) All(ctx context.Context, exec boil.ContextExecutor) (EventTimeUnitSlice, error) {
	var o []*EventTimeUnit

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EventTimeUnit slice")
	}

	if len(eventTimeUnitAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EventTimeUnit records in the query.
func (q eventTimeUnitQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count event_time_unit rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q eventTimeUnitQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if event_time_unit exists")
	}

	return count > 0, nil
}

// Event pointed to by the foreign key.
func (o *EventTimeUnit) Event(mods ...qm.QueryMod) eventQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.EventID),
	}

	queryMods = append(queryMods, mods...)

	return Events(queryMods...)
}

// UserEventAnswerUnits retrieves all the user_event_answer_unit's UserEventAnswerUnits with an executor.
func (o *EventTimeUnit) UserEventAnswerUnits(mods ...qm.QueryMod) userEventAnswerUnitQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`user_event_answer_unit`.`event_time_unit_id`=?", o.ID),
	)

	return UserEventAnswerUnits(queryMods...)
}

// LoadEvent allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (eventTimeUnitL) LoadEvent(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventTimeUnit interface{}, mods queries.Applicator) error {
	var slice []*EventTimeUnit
	var object *EventTimeUnit

	if singular {
		var ok bool
		object, ok = maybeEventTimeUnit.(*EventTimeUnit)
		if !ok {
			object = new(EventTimeUnit)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEventTimeUnit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEventTimeUnit))
			}
		}
	} else {
		s, ok := maybeEventTimeUnit.(*[]*EventTimeUnit)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEventTimeUnit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEventTimeUnit))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventTimeUnitR{}
		}
		args = append(args, object.EventID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventTimeUnitR{}
			}

			for _, a := range args {
				if a == obj.EventID {
					continue Outer
				}
			}

			args = append(args, obj.EventID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`event`),
		qm.WhereIn(`event.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Event")
	}

	var resultSlice []*Event
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Event")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for event")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for event")
	}

	if len(eventAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Event = foreign
		if foreign.R == nil {
			foreign.R = &eventR{}
		}
		foreign.R.EventTimeUnits = append(foreign.R.EventTimeUnits, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EventID == foreign.ID {
				local.R.Event = foreign
				if foreign.R == nil {
					foreign.R = &eventR{}
				}
				foreign.R.EventTimeUnits = append(foreign.R.EventTimeUnits, local)
				break
			}
		}
	}

	return nil
}

// LoadUserEventAnswerUnits allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (eventTimeUnitL) LoadUserEventAnswerUnits(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEventTimeUnit interface{}, mods queries.Applicator) error {
	var slice []*EventTimeUnit
	var object *EventTimeUnit

	if singular {
		var ok bool
		object, ok = maybeEventTimeUnit.(*EventTimeUnit)
		if !ok {
			object = new(EventTimeUnit)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEventTimeUnit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEventTimeUnit))
			}
		}
	} else {
		s, ok := maybeEventTimeUnit.(*[]*EventTimeUnit)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEventTimeUnit)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEventTimeUnit))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &eventTimeUnitR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &eventTimeUnitR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_event_answer_unit`),
		qm.WhereIn(`user_event_answer_unit.event_time_unit_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_event_answer_unit")
	}

	var resultSlice []*UserEventAnswerUnit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_event_answer_unit")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_event_answer_unit")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_event_answer_unit")
	}

	if len(userEventAnswerUnitAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserEventAnswerUnits = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userEventAnswerUnitR{}
			}
			foreign.R.EventTimeUnit = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.EventTimeUnitID {
				local.R.UserEventAnswerUnits = append(local.R.UserEventAnswerUnits, foreign)
				if foreign.R == nil {
					foreign.R = &userEventAnswerUnitR{}
				}
				foreign.R.EventTimeUnit = local
				break
			}
		}
	}

	return nil
}

// SetEvent of the eventTimeUnit to the related item.
// Sets o.R.Event to related.
// Adds o to related.R.EventTimeUnits.
func (o *EventTimeUnit) SetEvent(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Event) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `event_time_unit` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"event_id"}),
		strmangle.WhereClause("`", "`", 0, eventTimeUnitPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EventID = related.ID
	if o.R == nil {
		o.R = &eventTimeUnitR{
			Event: related,
		}
	} else {
		o.R.Event = related
	}

	if related.R == nil {
		related.R = &eventR{
			EventTimeUnits: EventTimeUnitSlice{o},
		}
	} else {
		related.R.EventTimeUnits = append(related.R.EventTimeUnits, o)
	}

	return nil
}

// AddUserEventAnswerUnits adds the given related objects to the existing relationships
// of the event_time_unit, optionally inserting them as new records.
// Appends related to o.R.UserEventAnswerUnits.
// Sets related.R.EventTimeUnit appropriately.
func (o *EventTimeUnit) AddUserEventAnswerUnits(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserEventAnswerUnit) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EventTimeUnitID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `user_event_answer_unit` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"event_time_unit_id"}),
				strmangle.WhereClause("`", "`", 0, userEventAnswerUnitPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EventTimeUnitID = o.ID
		}
	}

	if o.R == nil {
		o.R = &eventTimeUnitR{
			UserEventAnswerUnits: related,
		}
	} else {
		o.R.UserEventAnswerUnits = append(o.R.UserEventAnswerUnits, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userEventAnswerUnitR{
				EventTimeUnit: o,
			}
		} else {
			rel.R.EventTimeUnit = o
		}
	}
	return nil
}

// EventTimeUnits retrieves all the records using an executor.
func EventTimeUnits(mods ...qm.QueryMod) eventTimeUnitQuery {
	mods = append(mods, qm.From("`event_time_unit`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`event_time_unit`.*"})
	}

	return eventTimeUnitQuery{q}
}

// FindEventTimeUnit retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEventTimeUnit(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*EventTimeUnit, error) {
	eventTimeUnitObj := &EventTimeUnit{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `event_time_unit` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, eventTimeUnitObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from event_time_unit")
	}

	if err = eventTimeUnitObj.doAfterSelectHooks(ctx, exec); err != nil {
		return eventTimeUnitObj, err
	}

	return eventTimeUnitObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EventTimeUnit) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_time_unit provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventTimeUnitColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	eventTimeUnitInsertCacheMut.RLock()
	cache, cached := eventTimeUnitInsertCache[key]
	eventTimeUnitInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			eventTimeUnitAllColumns,
			eventTimeUnitColumnsWithDefault,
			eventTimeUnitColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `event_time_unit` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `event_time_unit` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `event_time_unit` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, eventTimeUnitPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into event_time_unit")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for event_time_unit")
	}

CacheNoHooks:
	if !cached {
		eventTimeUnitInsertCacheMut.Lock()
		eventTimeUnitInsertCache[key] = cache
		eventTimeUnitInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EventTimeUnit.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EventTimeUnit) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	eventTimeUnitUpdateCacheMut.RLock()
	cache, cached := eventTimeUnitUpdateCache[key]
	eventTimeUnitUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			eventTimeUnitAllColumns,
			eventTimeUnitPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update event_time_unit, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `event_time_unit` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, eventTimeUnitPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, append(wl, eventTimeUnitPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update event_time_unit row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for event_time_unit")
	}

	if !cached {
		eventTimeUnitUpdateCacheMut.Lock()
		eventTimeUnitUpdateCache[key] = cache
		eventTimeUnitUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q eventTimeUnitQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for event_time_unit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for event_time_unit")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EventTimeUnitSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventTimeUnitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `event_time_unit` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventTimeUnitPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in eventTimeUnit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all eventTimeUnit")
	}
	return rowsAff, nil
}

var mySQLEventTimeUnitUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EventTimeUnit) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no event_time_unit provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(eventTimeUnitColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEventTimeUnitUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	eventTimeUnitUpsertCacheMut.RLock()
	cache, cached := eventTimeUnitUpsertCache[key]
	eventTimeUnitUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			eventTimeUnitAllColumns,
			eventTimeUnitColumnsWithDefault,
			eventTimeUnitColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			eventTimeUnitAllColumns,
			eventTimeUnitPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert event_time_unit, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`event_time_unit`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `event_time_unit` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for event_time_unit")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(eventTimeUnitType, eventTimeUnitMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for event_time_unit")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for event_time_unit")
	}

CacheNoHooks:
	if !cached {
		eventTimeUnitUpsertCacheMut.Lock()
		eventTimeUnitUpsertCache[key] = cache
		eventTimeUnitUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EventTimeUnit record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EventTimeUnit) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EventTimeUnit provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), eventTimeUnitPrimaryKeyMapping)
	sql := "DELETE FROM `event_time_unit` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from event_time_unit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for event_time_unit")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q eventTimeUnitQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no eventTimeUnitQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from event_time_unit")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_time_unit")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EventTimeUnitSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(eventTimeUnitBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventTimeUnitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `event_time_unit` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventTimeUnitPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from eventTimeUnit slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for event_time_unit")
	}

	if len(eventTimeUnitAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EventTimeUnit) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEventTimeUnit(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EventTimeUnitSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EventTimeUnitSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), eventTimeUnitPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `event_time_unit`.* FROM `event_time_unit` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, eventTimeUnitPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EventTimeUnitSlice")
	}

	*o = slice

	return nil
}

// EventTimeUnitExists checks if the EventTimeUnit row exists.
func EventTimeUnitExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `event_time_unit` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if event_time_unit exists")
	}

	return exists, nil
}

// Exists checks if the EventTimeUnit row exists.
func (o *EventTimeUnit) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EventTimeUnitExists(ctx, exec, o.ID)
}
