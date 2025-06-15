CREATE TABLE
    "Alignment" (
        "Alignment" TEXT NOT NULL UNIQUE,
        "Domain" TEXT NOT NULL DEFAULT 'Public',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("Alignment"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    );