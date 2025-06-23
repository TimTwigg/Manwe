CREATE TABLE
    "Action" (
        "StatBlockID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL CHECK ("ActionID" >= 0),
        "Name" TEXT DEFAULT '',
        "AttackType" TEXT DEFAULT '',
        "HitModifier" INTEGER DEFAULT 0,
        "Reach" INTEGER DEFAULT 0,
        "Targets" INTEGER DEFAULT 0,
        "Description" TEXT DEFAULT '',
        PRIMARY KEY ("StatBlockID", "ActionID"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE
    )