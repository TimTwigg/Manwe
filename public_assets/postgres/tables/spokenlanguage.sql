
CREATE TABLE "public"."spokenlanguage" (
  "statblockid" integer,
  "language" varchar(20),
  "description" text,
  CONSTRAINT "spokenlanguage_p_key" PRIMARY KEY ("statblockid", "language"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_3" FOREIGN KEY ("language") REFERENCES "public"."language" ("language") ON UPDATE CASCADE ON DELETE RESTRICT
)
