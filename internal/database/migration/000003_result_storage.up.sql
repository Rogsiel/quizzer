CREATE TABLE "result" (
  "id" bigserial NOT NULL,
  "quiz_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "user_name" varchar NOT NULL,
  "sent_at" timestamp NOT NULL DEFAULT 'now()',
  "score" int NOT NULL DEFAULT 0,
  "responses" int[] NOT NULL
);

ALTER TABLE "result" ADD FOREIGN KEY ("quiz_id") REFERENCES "quiz" ("id");
