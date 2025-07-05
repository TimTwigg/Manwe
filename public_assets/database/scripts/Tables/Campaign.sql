CREATE TABLE
    "Campaign" (
        "Campaign" TEXT NOT NULL CHECK ("Campaign" <> ''),
        "Domain" TEXT NOT NULL,
        "Description" TEXT DEFAULT '',
        "CreationDate" TEXT NOT NULL CHECK (length ("CreationDate") = 8),
        "LastModified" TEXT NOT NULL CHECK (length ("LastModified") = 8),
        PRIMARY KEY ("Campaign", "Domain"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE
    );