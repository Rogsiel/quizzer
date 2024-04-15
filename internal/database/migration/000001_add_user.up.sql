CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_name" varchar UNIQUE NOT NULL,
  "email" varchar NOT NULL,
  "is_email_verified" boolean NOT NULL DEFAULT false,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);


