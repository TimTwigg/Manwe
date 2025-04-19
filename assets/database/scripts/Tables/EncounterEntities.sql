CREATE TABLE
    "EncounterEntities" (
        "EncounterID" INTEGER NOT NULL,
        "RowID" INTEGER NOT NULL,
        "EntityID" INTEGER NOT NULL,
        "Suffix" TEXT DEFAULT '',
        "Initiative" INTEGER DEFAULT 0,
        "MaxHitPoints" INTEGER NOT NULL,
        "TempHitPoints" INTEGER DEFAULT 0,
        "CurrentHitPoints" INTEGER NOT NULL,
        "ArmorClassBonus" INTEGER DEFAULT 0,
        "Notes" TEXT DEFAULT '',
        "IsHostile" TEXT DEFAULT 'X' CHECK ("IsHostile" IN ('X', '')) COLLATE NOCASE,
        "EncounterLocked" TEXT DEFAULT '' CHECK ("EncounterLocked" IN ('X', '')) COLLATE NOCASE,
        PRIMARY KEY ("EncounterID", "RowID"),
        FOREIGN KEY ("EncounterID") REFERENCES "Encounter" ("EncounterID") ON DELETE CASCADE,
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID")
    )