CREATE TABLE
    "SuperAction" (
        "StatBlockID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "Type" TEXT NOT NULL CHECK ("Type" IN ('Legendary', 'Mythic', 'Lair')),
        "Name" TEXT NOT NULL DEFAULT 'X',
        "Description" TEXT NOT NULL,
        "Points" INTEGER NOT NULL,
        "IsRegional" TEXT CHECK ("IsRegional" in ('X', '')) COLLATE NOCASE,
        PRIMARY KEY ("StatBlockID", "ActionID", "Type"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE
    )