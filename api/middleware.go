package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fxmbx/go-simple-bank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_handler"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		//strings.Fields splits by space you can specify your split with strings.Split(authorizationHeader, " ") but fields is cooler 😃
		fields := strings.Split(authorizationHeader, " ")
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("authorization type not supported %s, use %s", authorizationType, authorizationTypeBearer)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}

}
