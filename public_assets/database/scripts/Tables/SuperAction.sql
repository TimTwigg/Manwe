CREATE TABLE
    "SuperAction" (
        "EntityID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "Type" TEXT NOT NULL CHECK ("Type" IN ('Legendary', 'Mythic', 'Lair')),
        "Name" TEXT NOT NULL DEFAULT 'X',
        "Description" TEXT NOT NULL,
        "Points" INTEGER NOT NULL,
        "IsRegional" TEXT CHECK ("IsRegional" in ('X', '')) COLLATE NOCASE,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "ActionID", "Type"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName")
    )