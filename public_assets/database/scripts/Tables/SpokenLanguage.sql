CREATE TABLE
    "SpokenLanguage" (
        "StatBlockID" INTEGER NOT NULL,
        "Language" TEXT NOT NULL,
        "Description" TEXT DEFAULT '',
        PRIMARY KEY ("Language", "StatBlockID"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE,
        FOREIGN KEY ("Language") REFERENCES "Language" ("Language") ON DELETE CASCADE
    )