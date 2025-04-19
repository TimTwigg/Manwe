DROP TABLE IF EXISTS "ActionDamage";

CREATE TABLE
    "ActionDamage" (
        "EntityID" INTEGER NOT NULL,
        "ActionID" INTEGER NOT NULL,
        "DamageID" INTEGER NOT NULL,
        "Amount" TEXT,
        "Type" TEXT,
        "AltDmgActive" TEXT CHECK ("AltDmgActive" IN ('X', '')) COLLATE NOCASE,
        "Amount2" TEXT,
        "Type2" TEXT,
        "AltDmgNote" TEXT,
        "SaveDmgActive" TEXT CHECK ("SaveDmgActive" IN ('X', '')) COLLATE NOCASE,
        "Ability" TEXT,
        "DC" INTEGER,
        "HalfDamage" TEXT CHECK ("HalfDamage" IN ('X', '')) COLLATE NOCASE,
        "SaveDmgNote" TEXT,
        PRIMARY KEY ("EntityID", "ActionID", "DamageID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("EntityID", "ActionID") REFERENCES "Action" ("EntityID", "ActionID") ON DELETE CASCADE,
        FOREIGN KEY ("Type") REFERENCES "DamageType" ("DamageType"),
        FOREIGN KEY ("Type2") REFERENCES "DamageType" ("DamageType"),
        FOREIGN KEY ("Ability") REFERENCES "Ability" ("Ability")
    );