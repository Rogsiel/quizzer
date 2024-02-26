-- name: CreateUser :one
INSERT INTO "user" (
  name, email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE "id" = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM "user"
ORDER BY "id"
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "user"
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
