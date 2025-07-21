CREATE TABLE "public"."encounterentities" (
  "encounterid" integer,
  "rowid" serial,
  "statblockid" integer NOT NULL,
  "suffix" varchar(3),
  "intiative" integer DEFAULT '0',
  "maxhitpoints" integer NOT NULL,
  "temphitpoints" integer DEFAULT '0',
  "currenthitpoints" integer NOT NULL,
  "armorclassbonus" smallint DEFAULT '0',
  "concentration" boolean NOT NULL DEFAULT 'false',
  "notes" text,
  "ishostile" boolean NOT NULL DEFAULT 'true',
  "encounterlocked" boolean NOT NULL DEFAULT 'false',
  "id" text NOT NULL,
  CONSTRAINT "encounterentities_p_key" PRIMARY KEY ("encounterid", "rowid"),
  CONSTRAINT "constraint_1" FOREIGN KEY ("encounterid") REFERENCES "public"."encounter" ("encounterid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_2" FOREIGN KEY ("statblockid") REFERENCES "public"."statblock" ("statblockid") ON UPDATE CASCADE ON DELETE CASCADE
)