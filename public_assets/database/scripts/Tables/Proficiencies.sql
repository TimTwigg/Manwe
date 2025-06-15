CREATE TABLE
    "Proficiencies" (
        "EntityID" INTEGER NOT NULL,
        "Item" INTEGER NOT NULL CHECK ("Item" >= 0),
        "Type" TEXT NOT NULL CHECK ("Type" IN ('SK', 'ST')),
        "Name" TEXT NOT NULL,
        "Level" INTEGER DEFAULT 0,
        "Override" INTEGER DEFAULT 0,
        PRIMARY KEY ("EntityID", "Item"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "ModifierType" ("ModifierType") ON DELETE CASCADE
    )