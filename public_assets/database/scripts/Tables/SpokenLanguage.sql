CREATE TABLE
    "SpokenLanguage" (
        "EntityID" INTEGER NOT NULL,
        "Language" TEXT NOT NULL,
        "Description" TEXT DEFAULT '',
        "Domain" TEXT NOT NULL DEFAULT 'Private',
        "Published" TEXT NOT NULL DEFAULT '' CHECK ("Published" in ('', 'X')),
        PRIMARY KEY ("Language", "EntityID"),
        FOREIGN KEY ("EntityID") REFERENCES "Entity" ("EntityID") ON DELETE CASCADE,
        FOREIGN KEY ("Language") REFERENCES "Language" ("Language") ON DELETE CASCADE,
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    )