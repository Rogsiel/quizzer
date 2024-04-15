-- name: CreateVerifyEmail :one
INSERT INTO "verify_email" (
  user_name,
  email,
  secret_code
) VALUES (
  $1, $2, $3
) RETURNING *;
