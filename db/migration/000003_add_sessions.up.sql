CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar  NOT NULL,
  "client_ip" varchar NOT NULL ,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL,
  "expires_at" timestamptz NOT NULL DEFAULT (now())
);

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
