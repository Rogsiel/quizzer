-- name: CreateQuiz :one
INSERT INTO "quiz" (
  user_id, title, question_no, start_at, end_at, questions
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
