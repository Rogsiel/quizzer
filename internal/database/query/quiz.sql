-- name: GetQuiz :one
SELECT *
FROM "quiz"
WHERE id = $1;

-- name: CreateQuiz :one
INSERT INTO "quiz" (
  user_id, user_name, title, question_no, start_at, end_at, questions, answers) VALUES 
($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetCorrectAnswers :one
SELECT answers
FROM "quiz"
WHERE id = $1;

-- name: IncrementAnsweredCount :exec
UPDATE "quiz"
SET answered = answered + 1
WHERE id = $1;

-- name: GetUserQuiz :many
SELECT title, question_no, start_at, end_at
FROM "quiz"
WHERE "user_id" = $1;
