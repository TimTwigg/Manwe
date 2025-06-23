CREATE TABLE
    "ActionDamage" (
        "StatBlockID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "DamageID" INTEGER NOT NULL,
        "Amount" TEXT DEFAULT '',
        "Type" TEXT DEFAULT '',
        "AltDmgActive" TEXT DEFAULT '' CHECK ("AltDmgActive" IN ('X', '')) COLLATE NOCASE,
        "Amount2" TEXT DEFAULT '',
        "Type2" TEXT DEFAULT '',
        "AltDmgNote" TEXT DEFAULT '',
        "SaveDmgActive" TEXT DEFAULT '' CHECK ("SaveDmgActive" IN ('X', '')) COLLATE NOCASE,
        "Ability" TEXT DEFAULT '',
        "DC" INTEGER DEFAULT 0,
        "HalfDamage" TEXT DEFAULT '' CHECK ("HalfDamage" IN ('X', '')) COLLATE NOCASE,
        "SaveDmgNote" TEXT DEFAULT '',
        PRIMARY KEY ("StatBlockID", "ActionID", "DamageID"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE,
        FOREIGN KEY ("StatBlockID", "ActionID") REFERENCES "Action" ("StatBlockID", "ActionID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "DamageType" ("DamageType"),
        FOREIGN KEY ("Type2") REFERENCES "DamageType" ("DamageType"),
        FOREIGN KEY ("Ability") REFERENCES "Ability" ("Ability")
    );