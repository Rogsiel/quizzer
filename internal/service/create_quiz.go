package service

import(
    "context"
    "database/sql"
    "errors"
    "time"

    _ "github.com/lib/pq"

    "github.com/rogsiel/quizzer/internal/model"
)

type CreateQuizService struct{
    db *sql.DB
}

func NewCreateQuizService(db *sql.DB) *CreateQuizService{
    return &CreateQuizService{db: db}
}
func (q *CreateQuizService) CreateQuiz(ctx *context.Context) error{
    var quiz model.Quiz

}
