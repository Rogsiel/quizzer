package model

import (
	"database/sql"
	"time"
)

type Quiz struct {
	ID         int64           `json:"id"`
	UserID     int64           `json:"user_id"`
	Title      string          `json:"title"`
	QuestionNo int32           `json:"question_no"`
	StartAt    time.Time       `json:"start_at"`
	EndAt      sql.NullTime    `json:"end_at"`
	Questions  Question			`json:"questions"`
	Answers    []int32         `json:"answers"`
	Answered   int32           `json:"answered"`
}

type Result struct {
	ID        int64         `json:"id"`
	QuizID    int64         `json:"quiz_id"`
	UserID    int64         `json:"user_id"`
	SentAt    time.Time     `json:"sent_at"`
	Score     sql.NullInt32 `json:"score"`
	Responses []int32       `json:"responses"`
}

type Question map[string][]string
