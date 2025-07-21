
CREATE TABLE "public"."simpleaction" (
  "statblockid" integer,
  "actionid" serial,
  "type" SimpleActionType NOT NULL,
  "name" text,
  "description" text,
  CONSTRAINT "simpleaction_p_key" PRIMARY KEY ("statblockid", "actionid"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE
)