package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/fxmbx/go-simple-bank/db/sqlc"
	"github.com/fxmbx/go-simple-bank/utils"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {

	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errroHandler(ctx, err)
		return
	}
	hashed, err := utils.HashedPassword(req.Password)
	if err != nil {
		errroHandler(ctx, err)
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashed,
		FullName:       req.FullName,
		Email:          req.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		errroHandler(ctx, err)
		return
	}
	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}
func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errroHandler(ctx, err)
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse("🎣", err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("🎣", err))
		return
	}

	if err := utils.MatchPassword(req.Password, user.HashedPassword); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("🎣", err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(req.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("🎣", err))
		return
	}

	response := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, response)

}