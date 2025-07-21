CREATE TABLE "public"."entitytype" (
  "entitytype" varchar(20) PRIMARY KEY,
  "user" text NOT NULL DEFAULT 'public',
  "published" boolean NOT NULL DEFAULT 'false',
  CONSTRAINT "constraint_1" FOREIGN KEY ("user") REFERENCES "public"."user" ("username") ON UPDATE CASCADE ON DELETE CASCADE
)