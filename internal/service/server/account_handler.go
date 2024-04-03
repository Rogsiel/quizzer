package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rogsiel/quizzer/internal/database"
)

type createAccountReq struct {
    Name    string  `json:"name" binding:"required"`
    Email   string  `json:"email" binding:"required,email"`
}

func (server *Server) createAccount(ctx *gin.Context) {
    var req createAccountReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    arg := db.CreateUserParams{
	Name: req.Name,
	Email: req.Email,
    } 

    account, err := server.store.CreateUser(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    ctx.JSON(http.StatusOK, account)
}

type getAccountReq struct{
    ID	int64	`json:"name" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
    var req getAccountReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    arg := db.GetUserTxParams{
	ID: req.ID,
    }

    account, err := server.store.GetUser(ctx,arg.ID)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    ctx.JSON(http.StatusOK, account)
}
