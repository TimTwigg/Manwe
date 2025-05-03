CREATE TABLE
    "EncEntConditions" (
        "EncounterID" INTEGER NOT NULL,
        "RowID" INTEGER NOT NULL,
        "Condition" TEXT NOT NULL,
        "Duration" INTEGER NOT NULL DEFAULT 0,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EncounterID", "RowID", "Condition"),
        FOREIGN KEY ("EncounterID", "RowID") REFERENCES "EncounterEntities" ("EncounterID", "RowID") ON DELETE CASCADE,
        FOREIGN KEY ("Condition") REFERENCES "Condition" ("Condition"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )