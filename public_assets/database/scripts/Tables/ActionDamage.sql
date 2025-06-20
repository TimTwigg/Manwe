CREATE TABLE
    "ActionDamage" (
        "EntityID" INTEGER NOT NULL,
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
        PRIMARY KEY ("EntityID", "ActionID", "DamageID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("EntityID", "ActionID") REFERENCES "Action" ("EntityID", "ActionID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "DamageType" ("DamageType"),
        FOREIGN KEY ("Type2") REFERENCES "DamageType" ("DamageType"),
        FOREIGN KEY ("Ability") REFERENCES "Ability" ("Ability")
    );