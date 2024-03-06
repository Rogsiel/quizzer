// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: quiz.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

const createQuiz = `-- name: CreateQuiz :one
INSERT INTO "quiz" (
  user_id, title, question_no, start_at, end_at, questions, answers, answered
) VALUES 
($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id
`

type CreateQuizParams struct {
	UserID     int64           `json:"user_id"`
	Title      string          `json:"title"`
	QuestionNo int32           `json:"question_no"`
	StartAt    time.Time       `json:"start_at"`
	EndAt      sql.NullTime    `json:"end_at"`
	Questions  json.RawMessage `json:"questions"`
	Answers    []int32         `json:"answers"`
	Answered   int32           `json:"answered"`
}

func (q *Queries) CreateQuiz(ctx context.Context, arg CreateQuizParams) (int64, error) {
	row := q.queryRow(ctx, q.createQuizStmt, createQuiz,
		arg.UserID,
		arg.Title,
		arg.QuestionNo,
		arg.StartAt,
		arg.EndAt,
		arg.Questions,
		pq.Array(arg.Answers),
		arg.Answered,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getCorrectAnswers = `-- name: GetCorrectAnswers :one
SELECT answers
FROM "quiz"
WHERE id = $1
`

func (q *Queries) GetCorrectAnswers(ctx context.Context, id int64) ([]int32, error) {
	row := q.queryRow(ctx, q.getCorrectAnswersStmt, getCorrectAnswers, id)
	var answers []int32
	err := row.Scan(pq.Array(&answers))
	return answers, err
}

const incrementAnswerCount = `-- name: IncrementAnswerCount :exec
UPDATE "quiz"
SET answered = answered + 1
WHERE id = $1
`

func (q *Queries) IncrementAnswerCount(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.incrementAnswerCountStmt, incrementAnswerCount, id)
	return err
}
