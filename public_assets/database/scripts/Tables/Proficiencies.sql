CREATE TABLE
    "Proficiencies" (
        "EntityID" INTEGER NOT NULL,
        "Item" INTEGER NOT NULL CHECK ("Item" >= 0),
        "Type" TEXT NOT NULL CHECK ("Type" IN ('SK', 'ST')),
        "Name" TEXT NOT NULL,
        "Level" INTEGER DEFAULT 0,
        "Override" INTEGER DEFAULT 0,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "Item"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "ModifierType" ("ModifierType") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )