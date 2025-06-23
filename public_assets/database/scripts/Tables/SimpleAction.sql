CREATE TABLE
    "SimpleAction" (
        "StatBlockID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "Type" TEXT NOT NULL CHECK ("Type" IN ('Bonus', 'Reaction')),
        "Name" TEXT NOT NULL,
        "Description" TEXT NOT NULL,
        PRIMARY KEY ("StatBlockID", "ActionID"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE
    )