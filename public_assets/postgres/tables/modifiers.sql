
CREATE TABLE "public"."modifiers" (
  "statblockid" integer,
  "item" serial,
  "type" char(2) NOT NULL,
  "name" text NOT NULL,
  "value" integer NOT NULL,
  "description" text,
  CONSTRAINT "modifiers_constraint_1" PRIMARY KEY ("statblockid", "item"),
  CONSTRAINT "modifiers_constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "modifiers_constraint_3" FOREIGN KEY ("type") REFERENCES "public"."modifiertype" ("modifiertype") ON UPDATE CASCADE ON DELETE RESTRICT
)