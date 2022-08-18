package main

import (
	"database/sql"
	"log"

	"github.com/fxmbx/go-simple-bank/api"
	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println("cannot load configurations 😞")
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSoure)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}

}
