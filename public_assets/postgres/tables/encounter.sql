
CREATE TABLE "public"."encounter" (
  "encounterid" serial PRIMARY KEY,
  "name" varchar(20),
  "description" text,
  "creationdate" date NOT NULL,
  "accesseddate" date NOT NULL,
  "campaign" varchar(20) NOT NULL,
  "started" boolean NOT NULL DEFAULT 'false',
  "round" integer DEFAULT '0',
  "turn" integer DEFAULT '0',
  "haslair" boolean NOT NULL DEFAULT 'false',
  "lairownerid" integer DEFAULT '-1',
  "activeid" text,
  "user" text NOT NULL,
  "published" boolean NOT NULL DEFAULT 'false',
  CONSTRAINT "constraint_1" FOREIGN KEY ("campaign", "user") REFERENCES "public"."campaign" ("campaign", "user") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_2" FOREIGN KEY ("user") REFERENCES "public"."user" ("username") ON UPDATE CASCADE ON DELETE CASCADE
)