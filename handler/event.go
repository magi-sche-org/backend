package handler

import (
	"context"
	"database/sql"
	"log"

	"github.com/oklog/ulid/v2"

	"github.com/geekcamp-vol11-team30/backend/pb"
	"google.golang.org/protobuf/types/known/durationpb"
)

type EventServer struct {
	db *sql.DB
	pb.UnimplementedEventServer
}

func NewEventServer(db *sql.DB) *EventServer {
	return &EventServer{db: db}
}

func (s *EventServer) GetEvent(context.Context, *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	log.Println("GetEvent")

	return &pb.GetEventResponse{
		Id:       "01G65Z755AFWAKHE12NY0CQ9FH",
		Name:     "aaaa",
		Owner:    false,
		TimeUnit: &durationpb.Duration{},
		Duration: &durationpb.Duration{},
		Answers:  []*pb.Answer{},
	}, nil
}

func (s *EventServer) CreateEvent(ctx context.Context, r *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	log.Println("CreateEvent")
	db := s.db
	_, err := db.ExecContext(ctx, "INSERT INTO `users` (`id`, `name`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)", ulid.Make().String(), "aaaa", "2021-01-01 00:00:00", "2021-01-01 00:00:00")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.CreateEventResponse{}, nil
}

func (s *EventServer) RegisterAnswer(context.Context, *pb.RegisterAnswerRequest) (*pb.RegisterAnswerResponse, error) {
	log.Println("RegisterAnswer")
	return &pb.RegisterAnswerResponse{}, nil
}
