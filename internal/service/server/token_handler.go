package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenReq struct {
    RefreshToken	string  `json:"refresh_token" binding:"required"`
}

type renewAccessTokenRes struct {
    AccessToken		    string		`json:"access_token"`
    AccessTokenExpiresAt    time.Time		`json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
    var req renewAccessTokenReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
    	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    } 
    
    refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
    if err != nil {
	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	return
    }

    session, err := server.store.GetSession(ctx, refreshPayload.ID)
    if err != nil {
	if err == sql.ErrNoRows {
	    ctx.JSON(http.StatusNotFound, errorResponse(err))
	    return
	}
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    if session.IsBlocked {
	err := fmt.Errorf("your session is blocked")
	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	return
    }
    
    if session.UserName != refreshPayload.UserName {
	err := fmt.Errorf("Unauthorized user session")
	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	return
    }

    if session.RefreshToken != req.RefreshToken {
	err := fmt.Errorf("mismatched session token")
	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	return
    }

    accessToken, accessPayload, err := server.tokenMaker.CreateToken(
	refreshPayload.UserName,
	server.config.AccessTokenDuration,
    )
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    res := renewAccessTokenRes{
	AccessToken:		accessToken,
	AccessTokenExpiresAt:	accessPayload.ExpiredAt,
    }

    ctx.JSON(http.StatusOK, res)
}
