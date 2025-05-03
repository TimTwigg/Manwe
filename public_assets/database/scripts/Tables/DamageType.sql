CREATE TABLE
    "DamageType" (
        "DamageType" TEXT NOT NULL UNIQUE COLLATE NOCASE,
        "Description" TEXT NOT NULL,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("DamageType"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName")
    )