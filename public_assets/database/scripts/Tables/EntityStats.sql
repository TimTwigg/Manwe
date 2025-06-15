CREATE TABLE
    "EntityStats" (
        "EntityID" INTEGER NOT NULL,
        "Ability" TEXT NOT NULL CHECK (length ("Ability") > 0),
        "Value" INTEGER NOT NULL DEFAULT 10 CHECK ("Value" >= 0),
        PRIMARY KEY ("EntityID", "Ability"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Ability") REFERENCES "Ability" ("Ability") ON DELETE CASCADE
    )