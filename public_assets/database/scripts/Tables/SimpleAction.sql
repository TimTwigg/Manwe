CREATE TABLE
    "SimpleAction" (
        "EntityID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "Type" TEXT NOT NULL CHECK ("Type" IN ('Bonus', 'Reaction')),
        "Name" TEXT NOT NULL,
        "Description" TEXT NOT NULL,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "ActionID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )