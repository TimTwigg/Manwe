CREATE TABLE
    "CampaignEntities" (
        "Campaign" TEXT NOT NULL,
        "Domain" TEXT NOT NULL,
        "RowID" INTEGER NOT NULL,
        "StatBlockID" INTEGER NOT NULL,
        "Notes" TEXT DEFAULT '',
        PRIMARY KEY ("Campaign", "Domain", "RowID"),
        FOREIGN KEY ("Domain") REFERENCES "User" ("UserName") ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY ("Campaign", "Domain") REFERENCES "Campaign" ("Campaign", "Domain") ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY ("StatBlockID") REFERENCES "StatBlock" ("StatBlockID") ON DELETE CASCADE
    );