package handler

import (
	"context"
	"database/sql"
	"log"

	"github.com/geekcamp-vol11-team30/backend/pb"
	"github.com/geekcamp-vol11-team30/backend/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthorizationServer struct {
	db *sql.DB
	pb.UnimplementedAuthorizeServer
}

func NewAuthorizationServer(db *sql.DB) *AuthorizationServer {
	return &AuthorizationServer{db: db}
}

func (s *AuthorizationServer) GetToken(ctx context.Context, r *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	log.Println("GetToken called")
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	defer tx.Rollback()
	u, err := store.NewUser()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	t, err := store.NewToken()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	err = u.Save(ctx, tx)
	if err != nil {
		log.Fatal("saveuser: ", err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	err = u.RegisterToken(ctx, tx, t)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	log.Println(u, t)
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}
	// s.db.ExecContext(ctx, "")
	return &pb.GetTokenReply{
		Token: string(t),
	}, nil
}

func (s *AuthorizationServer) LinkToken(context.Context, *pb.LinkTokenRequest) (*pb.LinkTokenReply, error) {
	log.Println("LinkToken called")
	return &pb.LinkTokenReply{}, nil
}
