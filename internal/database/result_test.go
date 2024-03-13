package db

import (
	"context"
	"testing"

	"github.com/rogsiel/quizzer/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomResult(t *testing.T) Result {
	user := createRandomUser(t)
	quiz := createRandomQuiz(t)
	
	var responses []int32
	for i := 0; i < int(quiz.QuestionNo); i++ {
		r := util.RandInt(int32(0), int32(2))
		responses = append(responses, r)
	}
	arg := SendAnswersParams{
		QuizID: quiz.ID,
		UserID: user.ID,
		Responses: responses,
	}
	result, err := testQueries.SendAnswers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, arg.QuizID, result.QuizID)
	require.Equal(t, arg.UserID, result.UserID)
	require.Equal(t, arg.Responses, result.Responses)

	require.NotZero(t, result.ID)

	return result
}

func TestSendAnswers(t *testing.T) {
	createRandomResult(t)
}

func TestUpdateScore(t *testing.T) {
	result := createRandomResult(t)
	score := util.RandInt(0, 10)
	arg := UpdateScoreParams{
		ID: result.ID,
		Score: score,
	}
	newScore, err := testQueries.UpdateScore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newScore)

	require.Equal(t, arg.Score, newScore)
}
