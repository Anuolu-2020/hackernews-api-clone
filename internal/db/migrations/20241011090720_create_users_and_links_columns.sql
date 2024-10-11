-- Create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "username" character varying(127) NOT NULL,
  "password" character varying(127) NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_username" UNIQUE ("username")
);
-- Create "links" table
CREATE TABLE "public"."links" (
  "id" bigserial NOT NULL,
  "title" character varying(255) NULL,
  "address" character varying(255) NULL,
  "user_id" bigint NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_links_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE SET NULL
);
