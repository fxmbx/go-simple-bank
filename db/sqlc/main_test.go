package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/fxmbx/go-simple-bank/utils"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	// var err error
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config files: ", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSoure)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
