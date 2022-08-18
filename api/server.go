package api

import (
	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

//serve all http request
type Server struct {
	store  db.Store
	router *gin.Engine
}

//create a new Server instance and set up al api routes for that server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to the router
	router.POST("/api/accounts", server.createAccount)
	router.GET("/api/accounts/:id", server.getAccount)
	router.PUT("/api/accounts/:id", server.updateAccount)
	router.GET("/accounts", server.getAccounts)
	server.router = router
	return server
}

//Runs Http server on the input address to start listening for api request
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(s string, err error) gin.H {
	return gin.H{"error": s + "\n" + err.Error()}
}
