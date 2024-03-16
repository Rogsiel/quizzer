package db

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/rogsiel/quizzer/internal/util"
	"github.com/stretchr/testify/require"
)

func randAnswers(n int) []int32{
	var answers []int32
	for i := 0; i < n; i++ {
		a := util.RandInt(int32(0), int32(2))
		answers = append(answers, a)
	}
	return answers
}
func createRandomQuiz(t *testing.T) Quiz{
	user := createRandomUser(t)
	qno :=  util.RandInt(3, 10)
	q, _ := json.Marshal(util.RandQuestions(qno))
	arg := CreateQuizParams{
		UserID: user.ID,
		Title:	util.RandString(100),
		QuestionNo:	qno,
		StartAt: time.Now(),
		Questions: json.RawMessage(q),
		Answers: randAnswers(int(qno)),	
	}
	quiz, err := testQueries.CreateQuiz(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, quiz)

	require.Equal(t, arg.UserID, quiz.UserID)
	require.Equal(t, arg.Title, quiz.Title)
	require.Equal(t, arg.QuestionNo, quiz.QuestionNo)
	require.JSONEq(t, string(arg.Questions), string(quiz.Questions))
	require.Equal(t, arg.Answers, quiz.Answers)

	require.NotZero(t, quiz.ID)
	require.NotZero(t, quiz.QuestionNo)

	return quiz
}

func TestCreateQuiz(t *testing.T) {
	createRandomQuiz(t)
}

func TestGetCorrectAnswers(t *testing.T) {
	quiz := createRandomQuiz(t)
	answers, err := testQueries.GetCorrectAnswers(context.Background(), quiz.ID)
	require.NoError(t, err)
	require.NotEmpty(t, answers)

	require.Equal(t, quiz.Answers, answers)
}

func TestIncrementAnsweredCount(t *testing.T) {
	quiz := createRandomQuiz(t)
	err := testQueries.IncrementAnsweredCount(context.Background(), quiz.ID)
	require.NoError(t, err)
}

func TestGetQuiz(t *testing.T) {
	quiz := createRandomQuiz(t)
	result, err := testQueries.GetQuiz(context.Background(), quiz.ID)
	require.NoError(t, err)
	
	require.Equal(t, quiz.Title, result.Title)
	require.Equal(t, quiz.Questions, result.Questions)
}
