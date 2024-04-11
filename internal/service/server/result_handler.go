package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rogsiel/quizzer/internal/auth/token"
	db "github.com/rogsiel/quizzer/internal/database"
)

type answerReq struct {
    Responses	[]int32	`json:"responses" binding:"required"`
}

func (server *Server) SendAnswer(ctx *gin.Context) {
    quizID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
    if err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
    }
    userID, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
    if err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
    }

    var req answerReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
 
    authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
    arg := db.AnswerTxParams{
	QuizID: quizID,
	UserID: userID,
	UserName:authPayload.UserName,	
	Responses: req.Responses,
    }
    result, err := server.store.AnswerTx(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    ctx.JSON(http.StatusOK, result)
}
