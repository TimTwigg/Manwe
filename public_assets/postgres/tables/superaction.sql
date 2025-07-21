CREATE TABLE "public"."superaction" (
  "statblockid" integer,
  "actionid" serial,
  "type" superactiontype NOT NULL,
  "name" text DEFAULT 'X',
  "description" text NOT NULL,
  "points" smallint NOT NULL,
  "isregional" boolean NOT NULL,
  CONSTRAINT "superaction_p_key" PRIMARY KEY ("statblockid", "actionid", "type"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE
)