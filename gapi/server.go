package gapi

import (
	"fmt"

	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/pb"
	token "github.com/fxmbx/go-simple-bank/token"
	"github.com/fxmbx/go-simple-bank/utils"
)

//serve all grpc request
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

//create a new Server instance and set up api routes for that server
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	return server, nil
}
