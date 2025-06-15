CREATE TABLE
    "Modifiers" (
        "EntityID" INTEGER NOT NULL,
        "Item" INTEGER NOT NULL CHECK ("Item" >= 0),
        "Type" TEXT NOT NULL CHECK ("Type" NOT IN ('SK', 'ST')),
        "Name" TEXT NOT NULL,
        "Value" INTEGER DEFAULT 0,
        "Description" TEXT DEFAULT '',
        PRIMARY KEY ("EntityID", "Item"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "ModifierType" ("ModifierType") ON DELETE CASCADE
    )