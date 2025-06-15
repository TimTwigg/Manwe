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
        PRIMARY KEY ("EntityID", "ActionID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE
    )