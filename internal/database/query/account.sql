-- name: CreateUser :one
INSERT INTO "user" (
  name, 
  email, 
  hashed_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE "name" = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM "user"
ORDER BY "name"
LIMIT $1
OFFSET $2;

-- name: UpdatePassword :one
UPDATE "user"
SET hashed_password = $2
WHERE name = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE name = $1;
