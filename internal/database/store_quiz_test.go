package db

import (
	"context"
	"testing"

	"github.com/rogsiel/quizzer/internal/util"
	"github.com/stretchr/testify/require"
)

func TestCreateQuizTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan CreateQuizTxResult)

	for i := 0; i < n; i++ {
		go func() {
			quiz := createRandomQuiz(t)
		
			result, err := store.CreateQuizTx(context.Background(), CreateQuizTxParams{
				UserID: quiz.UserID,
				Title: quiz.Title,
				QuestionNo: quiz.QuestionNo,
				StartAt: quiz.StartAt,
				EndAt: quiz.EndAt,
				Questions: util.RandQuestions(quiz.QuestionNo),
				Answers: quiz.Answers,
			})
			errs <- err
			results <- result
		}()
	}
	
	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)

		quiz := <- results
		require.NotEmpty(t, quiz.Quiz)
		require.NotZero(t, quiz.Quiz.ID)
	}
}

func TestGetQuizTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan Quiz)

	for i := 0; i < n; i++ {
		go func () {
			quiz := createRandomQuiz(t)

			result, err := store.GetQuizTx(context.Background(), quiz.ID)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)

		result := <- results
		require.NotEmpty(t, result)
	}
}

func TestGetCorrectAnswersTx(t *testing.T) {
	store := NewStore(testDB)
	n := 5

	errs := make(chan error)
	results := make(chan []int32)

	for i := 0; i < n; i++ {
		go func () {
			quiz := createRandomQuiz(t)

			answers, err := store.GetCorrectAnswersTx(context.Background(), quiz.ID)

			errs <- err
			results <- answers
		}()
	}

	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)

		answers := <- results
		require.NotEmpty(t, answers)
		require.NotZero(t, answers)
	}
}

func TestIncrementAnsweredCountTx(t *testing.T) {
	store := NewStore(testDB)
	n := 100

	errs := make(chan error)
	
	for i := 0; i < n; i++ {
		go func () {
			quiz := createRandomQuiz(t)
			err := store.IncrementAnsweredCountTx(context.Background(), quiz.ID)

			errs <- err
		}()
	}
	
	for i := 0; i < n; i++ {
		err := <- errs
		require.NoError(t, err)
	}
}
