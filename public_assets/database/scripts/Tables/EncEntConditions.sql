CREATE TABLE
    "EncEntConditions" (
        "EncounterID" INTEGER NOT NULL,
        "RowID" INTEGER NOT NULL,
        "Condition" TEXT NOT NULL,
        "Duration" INTEGER NOT NULL DEFAULT 0,
        PRIMARY KEY ("EncounterID", "RowID", "Condition"),
        FOREIGN KEY ("EncounterID", "RowID") REFERENCES "EncounterEntities" ("EncounterID", "RowID") ON DELETE CASCADE,
        FOREIGN KEY ("Condition") REFERENCES "Condition" ("Condition")
    )