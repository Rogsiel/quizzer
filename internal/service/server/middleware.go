package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rogsiel/quizzer/internal/auth/token"
)

const (
    authorizationHeaderKey = "authorization"
    authorizationTypeBearer = "bearer"
    authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker	token.Maker) gin.HandlerFunc {
    return  func(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
	    err := errors.New("authorization is not provided")
	    ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
	    return
	}
	
	authFields := strings.Fields(authorizationHeader)
	if len(authFields) < 2 {
	    err := errors.New("invalid authorization format")
	    ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
	    return
	}

	authorizationType := strings.ToLower(authFields[0])
	if authorizationType != authorizationTypeBearer {
	    err := fmt.Errorf("invalid authorization type: %s", authorizationType)
	    ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
	    return
	}

	accessToken := authFields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
	    ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
	    return
	}

	ctx.Set(authorizationPayloadKey, payload)
	ctx.Next()
    }
}
