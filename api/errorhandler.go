package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func errroHandler(ctx *gin.Context, err error) {
	if pqErr, ok := err.(*pq.Error); ok {
		log.Println(pqErr.Code)
		switch pqErr.Code.Name() {
		case "foreign_key_violation", "unique_violation":
			ctx.JSON(http.StatusForbidden, errorResponse("Unique key viloation🎣", err))
		}
	}
}
