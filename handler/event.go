package handler

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/geekcamp-vol11-team30/backend/pb"
	"github.com/geekcamp-vol11-team30/backend/store"
	"github.com/oklog/ulid/v2"
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

	return &pb.GetEventResponse{
		Id:                ulid.Make().String(),
		Name:              "aaaa",
		Owner:             false,
		TimeUnit:          &durationpb.Duration{},
		Duration:          &durationpb.Duration{},
		Answers:           []*pb.Answer{},
		ProposedStartTime: []*timestamppb.Timestamp{},
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
	log.Println(user)
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	defer tx.Rollback()
	// _, err = tx.ExecContext(ctx, "INSERT INTO `users` (`id`, `name`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)", "01G65Z755AFWAKHE12NY0CQ9FH", "aaaa", "2021-01-01 00:00:00", "2021-01-01 00:00:00")
	// if
	// _, err := db.ExecContext(ctx, "INSERT INTO `users` (`id`, `name`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)", ulid.Make().String(), "aaaa", "2021-01-01 00:00:00", "2021-01-01 00:00:00")
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	return &pb.CreateEventResponse{}, nil
}

func (s *EventServer) RegisterAnswer(ctx context.Context, r *pb.RegisterAnswerRequest) (*pb.RegisterAnswerResponse, error) {
	log.Println("RegisterAnswer")
	return &pb.RegisterAnswerResponse{}, nil
}
