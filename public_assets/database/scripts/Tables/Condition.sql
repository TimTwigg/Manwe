CREATE TABLE
    "Condition" (
        "Condition" TEXT NOT NULL UNIQUE COLLATE NOCASE,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("Condition"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName")
    )