CREATE TABLE "public"."modifiertype" (
  "modifiertype" char(2) PRIMARY KEY,
  "description" text NOT NULL,
  "isproficiencyrelevant" boolean NOT NULL,
  "user" text NOT NULL DEFAULT 'public',
  "published" boolean NOT NULL DEFAULT 'false',
  CONSTRAINT "constraint_1" FOREIGN KEY ("user") REFERENCES "public"."user" ("username") ON UPDATE CASCADE ON DELETE CASCADE
)