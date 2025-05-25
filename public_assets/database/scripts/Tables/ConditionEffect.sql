CREATE TABLE
    "ConditionEffect" (
        "Condition" TEXT NOT NULL COLLATE NOCASE,
        "EffectID" INTEGER NOT NULL,
        "Description" TEXT NOT NULL,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EffectID", "Condition"),
        FOREIGN KEY ("Condition") REFERENCES "Condition" ("Condition"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )