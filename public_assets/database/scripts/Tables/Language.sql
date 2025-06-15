CREATE TABLE
    "Language" (
        "Language" TEXT NOT NULL UNIQUE,
        "Description" INTEGER NOT NULL,
        "Domain" TEXT NOT NULL DEFAULT 'Public',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("Language"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    );