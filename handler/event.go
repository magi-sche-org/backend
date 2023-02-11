package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/geekcamp-vol11-team30/backend/pb"
	"github.com/geekcamp-vol11-team30/backend/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventServer struct {
	db *sql.DB
	pb.UnimplementedEventServer
}

func NewEventServer(db *sql.DB) *EventServer {
	return &EventServer{db: db}
}

func (s *EventServer) GetEvent(ctx context.Context, r *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	log.Println("GetEvent")
	token := store.Token(r.Token)
	user, err := token.FetchUser(ctx, s.db)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(1, err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}
	if err != nil {
		log.Println(2, err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	event, err := store.FetchEvent(ctx, s.db, store.EventID(r.Id))
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(3, err)
		return nil, status.Errorf(codes.NotFound, "Event not found")
	}
	if err != nil {
		log.Println(4, err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	isOwner := event.IsOwner(user)
	var starts []*timestamppb.Timestamp
	for _, start := range event.ProposedStartTime {
		starts = append(starts, timestamppb.New(start))
	}
	answers, err := store.FetchAnswersByEventId(ctx, s.db, event.ID)
	if err != nil {
		log.Println(5, err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	var ans []*pb.Answer
	for _, answer := range answers {
		var schedules []*pb.Answer_ProposedSchedule
		for _, schedule := range answer.Schedules {
			schedules = append(schedules, &pb.Answer_ProposedSchedule{
				StartTime:    timestamppb.New(schedule.StartTime),
				Availability: pb.Answer_ProposedSchedule_Availability(schedule.Availability),
			})
		}
		ans = append(ans, &pb.Answer{
			Name:     answer.Name,
			Note:     answer.Note,
			Schedule: schedules,
		})
	}

	return &pb.GetEventResponse{
		Id:                string(event.ID),
		Name:              event.Name,
		Owner:             isOwner,
		Duration:          durationpb.New(event.Duration),
		TimeUnit:          durationpb.New(event.UnitSecond),
		Answers:           ans,
		ProposedStartTime: starts,
	}, nil
}

func (s *EventServer) CreateEvent(ctx context.Context, r *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	log.Println("CreateEvent")
	token := store.Token(r.Token)
	user, err := token.FetchUser(ctx, s.db)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	event, err := store.NewEventByRequest(r, user)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}
	log.Println(user, event)

	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	defer tx.Rollback()
	err = event.Save(ctx, tx)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	tx.Commit()
	return &pb.CreateEventResponse{
		EventId: string(event.ID),
	}, nil
}

func (s *EventServer) RegisterAnswer(ctx context.Context, r *pb.RegisterAnswerRequest) (*pb.RegisterAnswerResponse, error) {
	log.Println("RegisterAnswer")
	token := store.Token(r.Token)
	user, err := token.FetchUser(ctx, s.db)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	event, err := store.FetchEvent(ctx, s.db, store.EventID(r.EventId))
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, "Event not found")
	}
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	log.Println(user, event)
	answer, err := store.NewAnswerByRequest(r, user, event)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}
	tx, err := s.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	defer tx.Rollback()
	log.Println(answer)
	err = answer.Save(ctx, tx)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	tx.Commit()

	return &pb.RegisterAnswerResponse{}, nil
}
