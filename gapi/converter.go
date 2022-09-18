package gapi

import (
	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {

	return &pb.User{
		Username:           user.Username,
		Fullname:           user.FullName,
		Email:              user.Email,
		PasswordChanagedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:          timestamppb.New(user.CreatedAt),
	}
}
