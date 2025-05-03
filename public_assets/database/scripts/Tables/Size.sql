CREATE TABLE
    "Size" (
        "Size" TEXT NOT NULL UNIQUE,
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("Size"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName")
    )