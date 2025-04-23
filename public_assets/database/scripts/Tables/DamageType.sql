CREATE TABLE
    "DamageType" (
        "DamageType" TEXT NOT NULL UNIQUE COLLATE NOCASE,
        "Description" TEXT NOT NULL,
        PRIMARY KEY ("DamageType")
    )