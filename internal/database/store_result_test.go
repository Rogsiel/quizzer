package db

import (
	"context"
	"testing"

	"github.com/rogsiel/quizzer/internal/util"
	"github.com/stretchr/testify/require"
)

func TestAnswerTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan AnswerTxResult)
	
	for i := 0; i < n; i++ {
		go func () {
			quiz := createRandomQuiz(t)
			answers := util.RandAnswers(int(quiz.QuestionNo))		
			user := createRandomUser(t)
			result, err := store.AnswerTx(context.Background(), AnswerTxParams{
				QuizID: quiz.ID,
				UserID: user.ID,
				Responses: answers,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)

		result := <- results
		require.NotEmpty(t, result.Result.ID)
	}
}
