CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TABLE "quiz" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "title" varchar NOT NULL,
  "question_no" int NOT NULL,
  "start_at" timestamp NOT NULL DEFAULT 'now()',
  "end_at" timestamp,
  "questions" JSONB NOT NULL,
  "answered" int NOT NULL DEFAULT 0,
  "answers" int[]
);

CREATE TABLE "result" (
  "id" bigserial NOT NULL,
  "quiz_id" bigserial NOT NULL,
  "user_id" bigserial NOT NULL,
  "sent_at" timestamp NOT NULL,
  "score" int NOT NULL DEFAULT 0,
  "responses" int[] NOT NULL
);

ALTER TABLE "quiz" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "result" ADD FOREIGN KEY ("quiz_id") REFERENCES "quiz" ("id");
