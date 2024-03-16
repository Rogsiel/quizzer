package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/rogsiel/quizzer/internal/model"
)


type CreateQuizTxParams struct{
	UserID		int64           `json:"user_id"`
	Title		string          `json:"title"`
	QuestionNo	int32           `json:"question_no"`
	StartAt		time.Time       `json:"start_at"`
	EndAt		sql.NullTime    `json:"end_at"`
	Questions	model.Question	`json:"questions"`
	Answers		[]int32         `json:"answers"`
}

type CreateQuizTxResult struct{
	Quiz	Quiz	`json:"quiz"`
}

func (store *Store) CreateQuizTx(ctx context.Context, arg CreateQuizTxParams) (CreateQuizTxResult, error) {
	var Quiz CreateQuizTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		Questions, err := json.Marshal(arg.Questions)
		if err != nil {
			return err
		}
		Quiz.Quiz, err = q.CreateQuiz(ctx, CreateQuizParams{
			UserID: arg.UserID,
			Title: arg.Title,
			QuestionNo: arg.QuestionNo,
			StartAt: arg.StartAt,
			EndAt: arg.EndAt,
			Questions: Questions,
			Answers: arg.Answers,
		})
		return nil
	})
	return Quiz, err
}

func (store *Store) GetQuizTx(ctx context.Context, id int64) (Quiz, error) {
	var quiz Quiz

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		quiz, err = store.GetQuiz(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return quiz, err
}

func (store *Store) GetCorrectAnswersTx(ctx context.Context, QuizID int64) ([]int32, error) {
	var Answers []int32

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		Answers, err = q.GetCorrectAnswers(ctx, QuizID)
		if err != nil {
			return err
		}
		return nil
	})
	return Answers, err
}


func (store *Store) IncrementAnsweredCountTx(ctx context.Context, QuizID int64) error {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.IncrementAnsweredCount(ctx, QuizID)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
