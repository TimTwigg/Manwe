CREATE TABLE
    "Proficiencies" (
        "StatBlockID" INTEGER NOT NULL,
        "Item" INTEGER NOT NULL CHECK ("Item" >= 0),
        "Type" TEXT NOT NULL CHECK ("Type" IN ('SK', 'ST')),
        "Name" TEXT NOT NULL,
        "Level" INTEGER DEFAULT 0,
        "Override" INTEGER DEFAULT 0,
        PRIMARY KEY ("StatBlockID", "Item"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "ModifierType" ("ModifierType") ON DELETE CASCADE
    )