CREATE TABLE
    "Lair" (
        "StatBlockID" INTEGER NOT NULL,
        "Name" TEXT,
        "Description" TEXT,
        "Initiative" INTEGER CHECK (Initiative >= 0),
        PRIMARY KEY ("StatBlockID"),
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE
    )