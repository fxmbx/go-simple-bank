package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/fxmbx/go-simple-bank/api"
	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/gapi"
	"github.com/fxmbx/go-simple-bank/pb"
	"github.com/fxmbx/go-simple-bank/utils"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig(".")

	// doc
	if err != nil {
		log.Println("cannot load configurations 😞")
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSoure)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := db.NewStore(conn)
	// runGinServer(config, store) //
	runGrpcServer(config, store)

}

func runGrpcServer(config utils.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("starting grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server: ", err)
	}

}
func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
