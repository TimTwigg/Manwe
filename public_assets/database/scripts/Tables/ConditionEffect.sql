CREATE TABLE
    "ConditionEffect" (
        "Condition" TEXT NOT NULL COLLATE NOCASE,
        "EffectID" INTEGER NOT NULL,
        "Description" TEXT NOT NULL,
        PRIMARY KEY ("EffectID", "Condition"),
        FOREIGN KEY ("Condition") REFERENCES "Condition" ("Condition") ON DELETE CASCADE
    )