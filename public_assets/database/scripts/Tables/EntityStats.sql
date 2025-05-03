CREATE TABLE
    "EntityStats" (
        "EntityID" INTEGER NOT NULL,
        "Ability" TEXT NOT NULL CHECK (length ("Ability") > 0),
        "Value" INTEGER NOT NULL DEFAULT 10 CHECK ("Value" >= 0),
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "Ability"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Ability") REFERENCES "Ability" ("Ability") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )