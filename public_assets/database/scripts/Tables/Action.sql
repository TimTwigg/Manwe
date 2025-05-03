CREATE TABLE
    "Action" (
        "EntityID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL CHECK ("ActionID" >= 0),
        "Name" TEXT,
        "AttackType" TEXT,
        "HitModifier" INTEGER,
        "Reach" INTEGER,
        "Targets" INTEGER,
        "Description" TEXT,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("EntityID", "ActionID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName")
    )