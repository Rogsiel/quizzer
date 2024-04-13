CREATE TABLE "session" (
  "id" uuid PRIMARY KEY NOT NULL,
  "user_name" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expires_at" timestamptz NOT NULL
);

ALTER TABLE "session" ADD FOREIGN KEY ("user_name") REFERENCES "user" ("user_name");
