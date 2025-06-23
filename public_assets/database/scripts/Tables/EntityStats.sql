CREATE TABLE
    "EntityStats" (
        "StatBlockID" INTEGER NOT NULL,
        "Ability" TEXT NOT NULL CHECK (length ("Ability") > 0),
        "Value" INTEGER NOT NULL DEFAULT 10 CHECK ("Value" >= 0),
        PRIMARY KEY ("StatBlockID", "Ability"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE,
        FOREIGN KEY ("Ability") REFERENCES "Ability" ("Ability") ON DELETE CASCADE
    )