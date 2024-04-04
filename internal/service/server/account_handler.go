package server

import (
	"database/sql"
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
    arg := db.SignupTxParams{
	Name: req.Name,
	Email: req.Email,
    } 

    account, err := server.store.SignupTx(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    ctx.JSON(http.StatusOK, account)
}

type getUserQuizListReq struct{
    ID    int64  `uri:"id" binding:"required"`
}

func (server *Server) getUserQuizList(ctx *gin.Context) {
    var req getUserQuizListReq
    if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    arg := db.GetUserQuizTxParams{
	ID: req.ID,
    }

    userQuizList, err := server.store.GetUserQuizTx(ctx,arg)
    if err != nil {
	if err == sql.ErrNoRows {
	    ctx.JSON(http.StatusNotFound, errorResponse(err))
	    return
	}

	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    ctx.JSON(http.StatusOK, userQuizList)
}
