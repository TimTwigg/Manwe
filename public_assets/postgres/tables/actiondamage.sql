
CREATE TABLE "public"."actiondamage" (
  "statblockid" integer,
  "actionid" integer,
  "damageid" serial,
  "amount" text,
  "type" varchar(20),
  "altdamageactive" boolean NOT NULL DEFAULT 'false',
  "amount2" text,
  "type2" varchar(20),
  "altdamagenote" text,
  "savedamageactive" boolean NOT NULL DEFAULT 'false',
  "ability" varchar(20),
  "dc" smallint DEFAULT '0',
  "halfdamage" boolean NOT NULL DEFAULT 'false',
  "savedamagenote" text,
  CONSTRAINT "actiondamage_p_key" PRIMARY KEY ("statblockid", "actionid", "damageid"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_3" FOREIGN KEY ("statblockid", "actionid") REFERENCES "public"."action" ("statblockid", "actionid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_4" FOREIGN KEY ("type") REFERENCES "public"."damagetype" ("damagetype") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "constraint_5" FOREIGN KEY ("type2") REFERENCES "public"."damagetype" ("damagetype") ON UPDATE CASCADE ON DELETE RESTRICT,
  CONSTRAINT "constraint_6" FOREIGN KEY ("ability") REFERENCES "public"."ability" ("ability") ON UPDATE CASCADE ON DELETE RESTRICT
)