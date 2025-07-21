
CREATE TABLE "public"."proficiencies" (
  "statblockid" integer,
  "item" integer,
  "type" char(2) NOT NULL,
  "name" text NOT NULL,
  "level" smallint NOT NULL DEFAULT '0',
  "override" smallint NOT NULL DEFAULT '0',
  CONSTRAINT "proficiencies_p_key" PRIMARY KEY ("statblockid", "item"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_3" FOREIGN KEY ("type") REFERENCES "public"."modifiertype" ("modifiertype") ON UPDATE CASCADE ON DELETE RESTRICT
)