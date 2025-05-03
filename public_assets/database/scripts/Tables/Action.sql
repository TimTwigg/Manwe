CREATE TABLE
    "Action" (
        "EntityID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL CHECK ("ActionID" >= 0),
        "Name" TEXT DEFAULT '',
        "AttackType" TEXT DEFAULT '',
        "HitModifier" INTEGER DEFAULT 0,
        "Reach" INTEGER DEFAULT 0,
        "Targets" INTEGER DEFAULT 0,
        "Description" TEXT DEFAULT '',
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "ActionID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )