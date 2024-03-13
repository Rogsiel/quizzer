-- name: SendAnswers :one
INSERT INTO "result" (
  quiz_id, user_id, sent_at, score, responses
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateScore :one
UPDATE "result"
  SET score = $2
  WHERE id = $1
  RETURNING score;
