CREATE TABLE
    "Encounter" (
        "EncounterID" INTEGER NOT NULL UNIQUE CHECK ("EncounterID" > 0),
        "Name" TEXT DEFAULT '',
        "Description" TEXT DEFAULT '',
        "CreationDate" TEXT NOT NULL CHECK (length ("CreationDate") = 8),
        "AccessedDate" TEXT NOT NULL CHECK (length ("AccessedDate") = 8),
        "Campaign" TEXT DEFAULT '',
        "Started" TEXT DEFAULT '' CHECK ("Started" IN ('X', '')) COLLATE NOCASE,
        "Round" INTEGER DEFAULT 0,
        "Turn" INTEGER DEFAULT 0,
        "HasLair" TEXT DEFAULT '' CHECK ("HasLair" IN ('X', '')) COLLATE NOCASE,
        "LairEntityName" TEXT DEFAULT '',
        "ActiveID" TEXT DEFAULT '',
        PRIMARY KEY ("EncounterID" AUTOINCREMENT)
    )