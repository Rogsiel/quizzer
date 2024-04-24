-- name: CreateOTP :one
INSERT INTO "otp" (
  email,
  otp_code,
  otp_type
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetOTP :one
SELECT * FROM "otp"
WHERE
  otp_code = @otp_code;

-- name: UpdateOTP :exec
UPDATE "otp"
SET
  is_used = TRUE
WHERE
  id = @id
  AND is_used = FALSE
  AND expired_at > now();
