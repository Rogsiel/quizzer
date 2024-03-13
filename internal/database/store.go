package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:	db,
		Queries: New(db),
	}
}

func (store *Store) CalculateScore(responses []int32, correctAnswers []int32) int32 {
    var score int32
    for i := 0; i < len(responses); i++ {
        if responses[i] == correctAnswers[i] {
            score++
        }
    }
    return score
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil{
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Tx err: %v, Rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type AnswerTxParams struct{
	QuizID int64 `json:"quiz_id"`
	UserID int64 `json:"user_id"`
	Responses []int32 `json:"responses"`
}

type AnswerTxResult struct{
	Result Result	`json:"result"`
	Score int32 `json:"score"`
}

func (store *Store) AnswerTx(ctx context.Context, arg AnswerTxParams) (AnswerTxResult, error) {
	var result AnswerTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		
		result.Result, err = q.SendAnswers(ctx, SendAnswersParams{
			QuizID: arg.QuizID,
			UserID: arg.UserID,
			Responses: arg.Responses,
		})
		if err != nil {
			return err
		}
		
		var correctAnswers []int32
		correctAnswers, err = q.GetCorrectAnswers(ctx, result.Result.QuizID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		
		score := store.CalculateScore(arg.Responses, correctAnswers)
		result.Score, err = q.UpdateScore(ctx, UpdateScoreParams{
			ID: result.Result.ID,
			Score: score,
		})
		if err != nil {
			return err
		}
		
		err = q.IncrementAnsweredCount(ctx, arg.QuizID)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
