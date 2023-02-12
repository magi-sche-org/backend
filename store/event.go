package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/geekcamp-vol11-team30/backend/pb"
	"github.com/oklog/ulid/v2"
)

type EventID string
type Event struct {
	ID                EventID
	Name              string
	Owner             User
	Description       string
	Duration          time.Duration
	UnitSecond        time.Duration
	ProposedStartTime []time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewEventByRequest(r *pb.CreateEventRequest, u User) (Event, error) {
	now := time.Now()
	entropy := rand.New(rand.NewSource(now.UnixNano()))
	ms := ulid.Timestamp(now)
	pst := make([]time.Time, len(r.ProposedStartTime))
	for i, t := range r.ProposedStartTime {
		pst[i] = t.AsTime()
	}
	event := Event{
		ID:                EventID(ulid.MustNew(ms, entropy).String()),
		Name:              r.Name,
		Owner:             u,
		Description:       "",
		Duration:          r.Duration.AsDuration(),
		UnitSecond:        r.TimeUnit.AsDuration(),
		ProposedStartTime: pst,
	}

	return event, nil
}

func (e *Event) Save(ctx context.Context, tx *sql.Tx) error {
	now := time.Now()
	entropy := rand.New(rand.NewSource(now.UnixNano()))
	ms := ulid.Timestamp(now)
	r, err := tx.ExecContext(ctx, "INSERT INTO events (id, name, owner_id, description, duration, unit_second, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", e.ID, e.Name, e.Owner.ID, e.Description, e.Duration, e.UnitSecond, now, now)
	if err != nil {
		return err
	}
	log.Println(r.LastInsertId())
	var values []string
	for _, t := range e.ProposedStartTime {
		values = append(values, fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", ulid.MustNew(ms, entropy).String(), e.ID, t.Format(time.DateTime), now.Format("2006-01-02 15:04:05.000000"), now.Format("2006-01-02 15:04:05.000000")))
	}
	log.Println("INSERT INTO unit_events (id, event_id, start_at, created_at, updated_at) VALUES " + strings.Join(values, ", "))
	_, err = tx.ExecContext(ctx, "INSERT INTO unit_events (id, event_id, start_at, created_at, updated_at) VALUES "+strings.Join(values, ", "))
	if err != nil {
		return err
	}
	return nil
}

func FetchEvent(ctx context.Context, db *sql.DB, id EventID) (Event, error) {
	var event Event
	err := db.QueryRowContext(ctx, "SELECT id, name, owner_id, description, duration, unit_second, created_at, updated_at FROM events WHERE id = ?", id).Scan(&event.ID, &event.Name, &event.Owner.ID, &event.Description, &event.Duration, &event.UnitSecond, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		return Event{}, err
	}

	rows, err := db.QueryContext(ctx, "SELECT start_at FROM unit_events WHERE event_id = ? ORDER BY start_at", id)
	if err != nil {
		return Event{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var start time.Time
		err := rows.Scan(&start)
		if err != nil {
			return Event{}, err
		}
		event.ProposedStartTime = append(event.ProposedStartTime, start)
	}
	return event, nil
}

func (e Event) IsOwner(u User) bool {
	return e.Owner.ID == u.ID
}

type Answer struct {
	ID        string
	UserID    UserID
	EventID   EventID
	Name      string
	Note      string
	Schedules []Schedule
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Schedule struct {
	StartTime    time.Time
	Availability int
}

func NewAnswerByRequest(r *pb.RegisterAnswerRequest, u User, e Event) (Answer, error) {
	now := time.Now()
	entropy := rand.New(rand.NewSource(now.UnixNano()))
	ms := ulid.Timestamp(now)
	schedules := make([]Schedule, len(r.Answer.Schedule))
	for i, s := range r.Answer.Schedule {
		schedules[i] = Schedule{
			StartTime:    s.StartTime.AsTime(),
			Availability: int(s.Availability),
		}
	}
	answer := Answer{
		ID:        ulid.MustNew(ms, entropy).String(),
		UserID:    u.ID,
		EventID:   e.ID,
		Name:      r.Answer.Name,
		Note:      r.Answer.Note,
		Schedules: schedules,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return answer, nil
}

func (a *Answer) Save(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM unit_statuses WHERE answer_id = ?", a.ID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	_, err = tx.ExecContext(ctx, "DELETE FROM answers WHERE user_id = ? AND event_id = ?", a.UserID, a.EventID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	now := time.Now()
	entropy := rand.New(rand.NewSource(now.UnixNano()))
	ms := ulid.Timestamp(now)
	r, err := tx.ExecContext(ctx, "INSERT INTO answers (id, user_id, event_id, name, note, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", a.ID, a.UserID, a.EventID, a.Name, a.Note, now, now)
	if err != nil {
		return err
	}
	log.Println(r.LastInsertId())
	var values []string
	for _, s := range a.Schedules {
		values = append(values, fmt.Sprintf("('%s', '%s', '%s', %d, '%s', '%s')", ulid.MustNew(ms, entropy).String(), a.ID, s.StartTime.Format(time.DateTime), s.Availability, now.Format("2006-01-02 15:04:05.000000"), now.Format("2006-01-02 15:04:05.000000")))
	}
	log.Println("INSERT INTO unit_statuses (id, answer_id, start_at, status, created_at, updated_at) VALUES " + strings.Join(values, ", "))
	_, err = tx.ExecContext(ctx, "INSERT INTO unit_statuses (id, answer_id, start_at, status, created_at, updated_at) VALUES "+strings.Join(values, ", "))
	if err != nil {
		return err
	}
	return nil
}

func FetchAnswersByEventId(ctx context.Context, db *sql.DB, id EventID) ([]Answer, error) {
	rows, err := db.QueryContext(ctx, "SELECT a.id, a.name, a.note, s.start_at, s.status FROM unit_statuses as s JOIN answers as a ON a.id = s.answer_id WHERE event_id = ?", id)
	// rows, err := db.QueryContext(ctx, "SELECT id, user_id, event_id, name, note, created_at, updated_at FROM answers WHERE event_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	answers := map[string]Answer{}
	for rows.Next() {
		var (
			id      string
			name    string
			note    string
			startAt time.Time
			status  int
		)
		err := rows.Scan(&id, &name, &note, &startAt, &status)
		if err != nil {
			return nil, err
		}
		answer, ok := answers[id]
		if !ok {
			answer = Answer{
				ID:        id,
				Name:      name,
				Note:      note,
				Schedules: []Schedule{},
			}
		}
		answer.Schedules = append(answer.Schedules, Schedule{
			StartTime:    startAt,
			Availability: status,
		})
		answers[id] = answer
	}
	lanswers := []Answer{}
	for _, a := range answers {
		lanswers = append(lanswers, a)
	}
	return lanswers, nil
}
