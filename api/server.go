package api

import (
	"fmt"

	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	token "github.com/fxmbx/go-simple-bank/token"
	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//serve all http request
type Server struct {
	config     utils.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

//create a new Server instance and set up al api routes for that server
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker, config: config}
	// router := gin.Default()

	//regitering custome validator with gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
		v.RegisterValidation("customemail", validEmail)
	}

	// //add routes to the router
	// router.POST("/api/accounts", server.createAccount)
	// router.GET("/api/accounts/:id", server.getAccount)
	// router.PUT("/api/accounts/:id", server.updateAccount)
	// router.GET("/accounts", server.getAccounts)

	// //transfers
	// router.POST("/api/transfer", server.createTransfer)

	// //users
	// router.POST("/api/users", server.createUser)
	// router.POST("/api/users/login", server.loginUser)

	// server.router = router
	server.setUpRouter()
	return server, nil
}

//Runs Http server on the input address to start listening for api request
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(s string, err error) gin.H {
	return gin.H{"error": s + "\n" + err.Error()}
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	//users
	router.POST("/api/users", server.createUser)
	router.POST("/api/users/login", server.loginUser)

	//authRoutes to add middleware
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	//add routes to the router
	authRoutes.POST("/api/accounts", server.createAccount)
	authRoutes.GET("/api/accounts/:id", server.getAccount)
	authRoutes.PUT("/api/accounts/:id", server.updateAccount)
	authRoutes.GET("/api/accounts", server.getAccounts)

	//transfers
	authRoutes.POST("/api/transfers", server.createTransfer)

	server.router = router
}
