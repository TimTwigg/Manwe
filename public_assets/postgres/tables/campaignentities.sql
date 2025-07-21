CREATE TABLE "public"."campaignentities" (
  "campaign" varchar(20),
  "user" text,
  "rowid" serial,
  "statblockid" integer NOT NULL,
  "notes" text,
  CONSTRAINT "campaignentities_p_key" PRIMARY KEY ("campaign", "user", "rowid"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("campaign", "user") REFERENCES "public"."campaign" ("campaign", "user") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_3" FOREIGN KEY ("user") REFERENCES "public"."user" ("username") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_4" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE
)