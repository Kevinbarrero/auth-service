CREATE TABLE "users" (
  "email" varchar NOT NULL PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp NOT NULL default (now())
);
