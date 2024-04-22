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
  email = @email
  AND otp_code = @otp_code
  AND otp_type = @otp_type;

-- name: UpdateOTP :exec
UPDATE "otp"
SET
  is_used = TRUE
WHERE
  id = @id
  AND otp_code = @otp_code
  AND is_used = FALSE
  AND expired_at > now();
