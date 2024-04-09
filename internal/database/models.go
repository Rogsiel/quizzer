// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Quiz struct {
	ID         int64           `json:"id"`
	UserID     int64           `json:"user_id"`
	Title      string          `json:"title"`
	QuestionNo int32           `json:"question_no"`
	StartAt    time.Time       `json:"start_at"`
	EndAt      sql.NullTime    `json:"end_at"`
	Questions  json.RawMessage `json:"questions"`
	Answered   int32           `json:"answered"`
	Answers    []int32         `json:"answers"`
}

type Result struct {
	ID        int64     `json:"id"`
	QuizID    int64     `json:"quiz_id"`
	UserID    int64     `json:"user_id"`
	SentAt    time.Time `json:"sent_at"`
	Score     int32     `json:"score"`
	Responses []int32   `json:"responses"`
}

type User struct {
	ID                int64     `json:"id"`
	UserName          string    `json:"user_name"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
