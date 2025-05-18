CREATE TABLE
    "ModifierType" (
        "ModifierType" TEXT NOT NULL UNIQUE,
        "Description" TEXT NOT NULL,
        "IsProficiencyRelevant" TEXT NOT NULL DEFAULT '' CHECK ("IsProficiencyRelevant" in ('', 'X')),
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("ModifierType"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    );