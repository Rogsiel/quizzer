package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rogsiel/quizzer/internal/auth"
	db "github.com/rogsiel/quizzer/internal/database"
)

type createAccountReq struct {
    UserName	string  `json:"user_name" binding:"required,alphanum"`
    Email	string  `json:"email" binding:"required,email"`
    Password	string	`json:"password" binding:"required,min=8"`
}
type userResponse struct {
    UserName		string	    `json:"user_name"`
    Email		string	    `json:"email"`
    PasswrordChangedAt	time.Time   `json:"password_changed_at"`
    CreatedAt		time.Time   `json:"created_at"`
}

func newUserResponse(user db.User) userResponse{
    return userResponse{
	UserName: user.UserName,
	Email: user.Email,
	PasswrordChangedAt: user.PasswordChangedAt,
	CreatedAt: user.CreatedAt,
    }
}

func (server *Server) createAccount(ctx *gin.Context) {
    var req createAccountReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    
    hashedPassword, err := auth.HashPassword(req.Password)
    if err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err)) 
	return
    }
    arg := db.CreateUserParams{
	UserName: req.UserName,
	Email: req.Email,
	HashedPassword: hashedPassword,	
    } 

    account, err := server.store.CreateUser(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    res := newUserResponse(account)
    ctx.JSON(http.StatusOK, res)
}

type userLoginReq struct {
    UserName	string  `json:"user_name" binding:"required,alphanum"`
    Password	string	`json:"password" binding:"required,min=8"`
}

type userLoginRes struct {
    AccessToken	string		`json:"access_token"`
    User	userResponse    `json:"user"`
}

func (server *Server) UserLogin(ctx *gin.Context) {
    var req userLoginReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
    	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    } 
    
    user, err := server.store.GetUser(ctx, req.UserName)
    if err != nil {
	if err == sql.ErrNoRows {
	    ctx.JSON(http.StatusNotFound, errorResponse(err))
	    return
	}
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    err = auth.CheckPassword(req.Password, user.HashedPassword)
    if err != nil {
	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	return
    }

    accessToken, err := server.tokenMaker.CreateToken(
	user.UserName,
	server.config.AccessTokenDuration,
    )
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    res := userLoginRes{
	AccessToken: accessToken,
	User: newUserResponse(user),
    }

    ctx.JSON(http.StatusOK, res)
}
