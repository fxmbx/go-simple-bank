package api

import (
	"os"
	"testing"
	"time"

	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
// )

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}
