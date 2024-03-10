package db

import (
	"testing"
	"time"

	"github.com/rogsiel/quizzer/internal/util"
)

func createRandomQuiz(t *testing.T) Quiz{
	user := createRandomUser(t)
	arg := CreateQuizParams{
		UserID: user.ID,
		Title:	util.RandString(100),
		QuestionNo: int32(util.RandInt(3, 10)),
		StartAt: time.Now(),
		Questions: util.RandQuestions(QuestionNo),
		
	}
}
