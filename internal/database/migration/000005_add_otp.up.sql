CREATE TABLE "otp" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "otp_code" varchar NOT NULL,
  "otp_type" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "otp" ADD FOREIGN KEY ("email") REFERENCES "user" ("email");
