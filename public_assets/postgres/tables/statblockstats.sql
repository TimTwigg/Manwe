CREATE TABLE "public"."statblockstats" (
  "statblockid" integer NOT NULL,
  "ability" varchar(20) NOT NULL,
  "value" integer NOT NULL DEFAULT '10',
  CONSTRAINT "constraint_1" PRIMARY KEY ("statblockid", "ability"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_3" FOREIGN KEY ("ability") REFERENCES "public"."ability" ("ability") ON UPDATE CASCADE ON DELETE CASCADE
)