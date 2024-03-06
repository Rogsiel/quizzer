-- name: CreateQuiz :one
INSERT INTO "quiz" (
  user_id, title, question_no, start_at, end_at, questions, answers, answered
) VALUES 
($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: GetCorrectAnswers :one
SELECT answers
FROM "quiz"
WHERE id = $1;

-- name: IncrementAnswerCount :exec
UPDATE "quiz"
SET answered = answered + 1
WHERE id = $1;
