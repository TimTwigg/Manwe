CREATE TABLE
    "Lair" (
        "EntityID" INTEGER NOT NULL,
        "Name" TEXT,
        "Description" TEXT,
        "Initiative" INTEGER CHECK (Initiative >= 0),
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )