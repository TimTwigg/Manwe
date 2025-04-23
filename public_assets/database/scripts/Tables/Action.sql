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
        PRIMARY KEY ("EntityID", "ActionID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE
    )