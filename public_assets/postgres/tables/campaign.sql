
CREATE TABLE "public"."campaign" (
  "campaign" varchar(20),
  "user" text,
  "description" text,
  "creationdate" date NOT NULL,
  "lastmodified" date NOT NULL,
  CONSTRAINT "campaign_p_key" PRIMARY KEY ("campaign", "user"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("user") REFERENCES "public"."user" ("username") ON UPDATE CASCADE ON DELETE CASCADE
)