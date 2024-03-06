-- name: AnswerQuiz :one
INSERT INTO "result" (
  quiz_id, user_id, score, responses
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: InsertScore :one
UPDATE "result"
  SET score = $2
  WHERE id = $1
  RETURNING score;
