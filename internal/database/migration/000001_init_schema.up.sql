CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "quiz" (
  "id" incremental PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "title" varchar NOT NULL,
  "question_no" int NOT NULL,
  "start_at" timestamptz NOT NULL DEFAULT 'now()',
  "end_at" timestamptz,
  "questions" question[] NOT NULL
);

CREATE TABLE "result" (
  "id" incremental NOT NULL,
  "quiz_id" bigserial NOT NULL,
  "score" int NOT NULL,
  "responses" int[][] NOT NULL
);

ALTER TABLE "quiz" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "result" ADD FOREIGN KEY ("quiz_id") REFERENCES "quiz" ("id");
