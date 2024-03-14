package db

import (
	"context"
	"time"
)

func calculateScore(responses []int32, correctAnswers []int32) int32 {
    var score int32
    for i := 0; i < len(responses); i++ {
        if responses[i] == correctAnswers[i] {
            score++
        }
    }
    return score
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
	var Result AnswerTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		quizAnswers, err := store.GetCorrectAnswersTx(ctx, arg.QuizID)
		if err != nil {
			return err
		}
		userAnswers := arg.Responses
		score := calculateScore(userAnswers, quizAnswers)
		
		Result.Result, err = q.SendAnswers(ctx, SendAnswersParams{
			QuizID: arg.QuizID,
			UserID: arg.UserID,
			SentAt: time.Now(),
			Score: score,
			Responses: arg.Responses,
		})
		if err != nil {
			return err
		}
		
		err = store.IncrementAnsweredCount(ctx, arg.QuizID)
		if err != nil {
			return err
		}
		return nil
	})
	return Result, err
}
