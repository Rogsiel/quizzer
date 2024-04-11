package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rogsiel/quizzer/internal/auth/token"
	db "github.com/rogsiel/quizzer/internal/database"
	"github.com/rogsiel/quizzer/internal/model"
)

type createQuizReq struct {
	UserID		int64           `json:"user_id" binding:"required"`
	Title		string          `json:"title" binding:"required"`
	QuestionNo	int32           `json:"question_no" binding:"required"`
	StartAt		time.Time       `json:"start_at" binding:"required"`
	EndAt		sql.NullTime    `json:"end_at"`
	Questions	model.Question	`json:"questions" binding:"required"`
	Answers		[]int32         `json:"answers" binding:"required"` 
}

func (server *Server) createQuiz(ctx *gin.Context) {
    var req createQuizReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    
    authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
    arg := db.CreateQuizTxParams{
	UserID: req.UserID,
	UserName: authPayload.UserName,
	Title: req.Title,
	QuestionNo: req.QuestionNo,
	StartAt: req.StartAt,
	EndAt: req.EndAt,
	Questions: req.Questions,
	Answers: req.Answers,
    }
    quiz, err := server.store.CreateQuizTx(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    }

    ctx.JSON(http.StatusOK, quiz)
}

type getQuizByIDReq struct {
    ID    int64  `uri:"id" binding:"required"`
}

func (server *Server) getQuizByID(ctx *gin.Context) {
    var req getQuizByIDReq
    if err := ctx.ShouldBindUri(&req); err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    arg := req.ID
    quiz, err := server.store.GetQuizTx(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
    } 

    ctx.JSON(http.StatusOK, quiz)
}

type getUserQuizListReq struct{
    ID    int64  `uri:"user_id" binding:"required"`
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

    userQuizList, err := server.store.GetUserQuiz(ctx,arg.ID)
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
