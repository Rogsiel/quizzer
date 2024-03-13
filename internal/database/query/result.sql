-- name: SendAnswers :one
INSERT INTO "result" (
  quiz_id, user_id, responses
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateScore :one
UPDATE "result"
  SET score = $2
  WHERE id = $1
  RETURNING score;
