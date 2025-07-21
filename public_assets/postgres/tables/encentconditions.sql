
CREATE TABLE "public"."encentconditions" (
  "encounterid" integer,
  "rowid" integer,
  "condition" varchar(20),
  "duration" integer NOT NULL DEFAULT '0',
  CONSTRAINT "encentconditions_p_key" PRIMARY KEY ("encounterid", "rowid", "condition"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("encounterid", "rowid") REFERENCES "public"."encounterentities" ("encounterid", "rowid") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "constraint_3" FOREIGN KEY ("condition") REFERENCES "public"."condition" ("condition") ON UPDATE CASCADE ON DELETE RESTRICT
)