CREATE TABLE
    "SimpleAction" (
        "EntityID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "Type" TEXT NOT NULL CHECK ("Type" IN ('Bonus', 'Reaction')),
        "Name" TEXT NOT NULL,
        "Description" TEXT NOT NULL,
        PRIMARY KEY ("EntityID", "ActionID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE
    )