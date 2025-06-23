CREATE TABLE
    "Modifiers" (
        "StatBlockID" INTEGER NOT NULL,
        "Item" INTEGER NOT NULL CHECK ("Item" >= 0),
        "Type" TEXT NOT NULL CHECK ("Type" NOT IN ('SK', 'ST')),
        "Name" TEXT NOT NULL,
        "Value" INTEGER DEFAULT 0,
        "Description" TEXT DEFAULT '',
        PRIMARY KEY ("StatBlockID", "Item"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "ModifierType" ("ModifierType") ON DELETE CASCADE
    )