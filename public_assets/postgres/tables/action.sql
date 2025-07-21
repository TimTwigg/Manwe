
CREATE TABLE "public"."action" (
  "statblockid" integer,
  "actionid" serial NOT NULL,
  "name" text NOT NULL,
  "attacktype" text,
  "hitmodifier" smallint DEFAULT '0',
  "reach" integer DEFAULT '0',
  "targets" integer DEFAULT '0',
  "description" text,
  CONSTRAINT "constraint_1" PRIMARY KEY ("statblockid", "actionid"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE
)