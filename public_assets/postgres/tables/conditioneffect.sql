
CREATE TABLE "public"."conditioneffect" (
  "condition" text,
  "effectid" serial,
  "description" text NOT NULL,
  CONSTRAINT "constraint_1" PRIMARY KEY ("condition", "effectid"),
  CONSTRAINT "constraint_2" FOREIGN KEY ("condition") REFERENCES "public"."condition" ("condition")
)