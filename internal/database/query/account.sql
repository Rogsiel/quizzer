-- name: CreateUser :one
INSERT INTO "user" (
  user_name, 
  email, 
  hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE "user_name" = $1 LIMIT 1;

-- name: GetUsers :many
SELECT "id", "user_name" FROM "user"
ORDER BY "user_name"
LIMIT $1
OFFSET $2;

-- name: UpdatePassword :exec
UPDATE "user"
SET hashed_password = $1, password_changed_at = NOW()
WHERE email = $2;

-- name: VerifyEmail :one
UPDATE "user"
SET is_email_verified = TRUE
WHERE email = @email
RETURNING *;

-- name: IsUserVerified :one
SELECT is_email_verified FROM "user"
WHERE id = @user_id;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_name = $1;
