package gapi

import (
	"context"
	"database/sql"
	"log"
	"time"

	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/pb"
	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashed, err := utils.HashedPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashed,
		FullName:       req.GetFullname(),
		Email:          req.GetEmail(),
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code)
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "Unique key viloation : %s", err)

			default:
				return nil, status.Errorf(codes.Internal, "failed to create user : %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	response := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return response, nil
}

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "no user found: %s", err)

		}
		return nil, status.Errorf(codes.Internal, "error processing request: %s", err)
	}
	if err = utils.MatchPassword(req.GetPassword(), user.HashedPassword); err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "incorrect password: %s", err)
	}
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(req.GetUsername(), time.Minute*15)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "incorrect password: %s", err)
	}
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(req.GetUsername(), time.Hour*24)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "incorrect password: %s", err)
	}

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}
	session, err := server.store.CreateSession(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "incorrect password: %s", err)
	}

	response := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return response, nil
}
