CREATE TABLE
    "Modifiers" (
        "EntityID" INTEGER NOT NULL,
        "Item" INTEGER NOT NULL CHECK ("Item" >= 0),
        "Type" TEXT NOT NULL,
        "Name" TEXT NOT NULL,
        "Value" INTEGER,
        "Description" TEXT,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "Item"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "ModifierType" ("ModifierType") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName")
    )