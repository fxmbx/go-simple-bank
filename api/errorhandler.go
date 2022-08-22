package api

import (
	"database/sql"
	"log"
	"net/http"

	token "github.com/fxmbx/go-simple-bank/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func errroHandler(ctx *gin.Context, err error) {
	if pqErr, ok := err.(*pq.Error); ok {
		log.Println(pqErr.Code)
		switch pqErr.Code.Name() {
		case "foreign_key_violation", "unique_violation":
			ctx.JSON(http.StatusForbidden, errorResponse("Unique key viloation🎣", err))
			return
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse("🎣", err))
			return
		}
	}
	switch err {
	case sql.ErrNoRows:
		ctx.JSON(http.StatusNotFound, errorResponse("🎣", err))
		return
	case sql.ErrConnDone:
		ctx.JSON(http.StatusInternalServerError, errorResponse("🎣", err))
		return
	case token.ErrExpiredToken:
		ctx.JSON(http.StatusUnauthorized, errorResponse("🎣", err))
		return
	case token.ErrInvalidToken:
		ctx.JSON(http.StatusForbidden, errorResponse("🎣", err))
		return
	}

	ctx.JSON(http.StatusBadRequest, errorResponse("🎣", err))

}
