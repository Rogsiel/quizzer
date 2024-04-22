CREATE TABLE "quiz" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "user_name" varchar NOT NULL,
  "title" varchar NOT NULL,
  "question_no" int NOT NULL,
  "start_at" timestamp NOT NULL DEFAULT 'now()',
  "end_at" timestamp,
  "questions" JSONB NOT NULL,
  "answered" int NOT NULL DEFAULT 0,
  "answers" int[]
);

ALTER TABLE "quiz" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
ALTER TABLE "quiz" ADD FOREIGN KEY ("user_name") REFERENCES "user" ("user_name");
