CREATE TABLE "user" (
  "id" SERIAL NOT NULL,
  "full_name" varchar(60) NOT NULL DEFAULT '',
  "phone_number" varchar(15) NOT NULL DEFAULT '',
  "password" varchar(64) NOT NULL DEFAULT '',
  "account_salt" varchar(15) NOT NULL DEFAULT '',
  "successful_login" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NULL DEFAULT NULL,
  PRIMARY KEY ("id"),
  UNIQUE ("phone_number")
);