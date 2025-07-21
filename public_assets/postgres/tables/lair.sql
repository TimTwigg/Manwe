
CREATE TABLE "public"."lair" (
  "statblockid" integer PRIMARY KEY,
  "name" text,
  "description" text,
  "intiative" integer DEFAULT '20',
  CONSTRAINT "constraint_1" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE
)