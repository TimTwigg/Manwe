CREATE TABLE
    "Ability" (
        "Ability" TEXT NOT NULL UNIQUE COLLATE NOCASE,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("Ability"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )