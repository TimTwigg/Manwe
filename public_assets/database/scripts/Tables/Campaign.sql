CREATE TABLE
    "Campaign" (
        "Campaign" TEXT NOT NULL UNIQUE CHECK ("Campaign" <> ''),
        "Domain" TEXT NOT NULL,
        "Description" TEXT DEFAULT '',
        PRIMARY KEY ("Campaign", "Domain"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    );