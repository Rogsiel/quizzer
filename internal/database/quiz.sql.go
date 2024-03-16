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
  user_id, title, question_no, start_at, end_at, questions, answers) VALUES 
($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, title, question_no, start_at, end_at, questions, answered, answers
`

type CreateQuizParams struct {
	UserID     int64           `json:"user_id"`
	Title      string          `json:"title"`
	QuestionNo int32           `json:"question_no"`
	StartAt    time.Time       `json:"start_at"`
	EndAt      sql.NullTime    `json:"end_at"`
	Questions  json.RawMessage `json:"questions"`
	Answers    []int32         `json:"answers"`
}

func (q *Queries) CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error) {
	row := q.queryRow(ctx, q.createQuizStmt, createQuiz,
		arg.UserID,
		arg.Title,
		arg.QuestionNo,
		arg.StartAt,
		arg.EndAt,
		arg.Questions,
		pq.Array(arg.Answers),
	)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.QuestionNo,
		&i.StartAt,
		&i.EndAt,
		&i.Questions,
		&i.Answered,
		pq.Array(&i.Answers),
	)
	return i, err
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

const getQuiz = `-- name: GetQuiz :one
SELECT id, user_id, title, question_no, start_at, end_at, questions, answered, answers
FROM "quiz"
WHERE id = $1
`

func (q *Queries) GetQuiz(ctx context.Context, id int64) (Quiz, error) {
	row := q.queryRow(ctx, q.getQuizStmt, getQuiz, id)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.QuestionNo,
		&i.StartAt,
		&i.EndAt,
		&i.Questions,
		&i.Answered,
		pq.Array(&i.Answers),
	)
	return i, err
}

const incrementAnsweredCount = `-- name: IncrementAnsweredCount :exec
UPDATE "quiz"
SET answered = answered + 1
WHERE id = $1
`

func (q *Queries) IncrementAnsweredCount(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.incrementAnsweredCountStmt, incrementAnsweredCount, id)
	return err
}
