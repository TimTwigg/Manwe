CREATE TABLE "SuperAction" (
	"EntityID"	INTEGER NOT NULL,
	"ActionID"	INTEGER NOT NULL,
	"Type"	TEXT NOT NULL CHECK("Type" IN ('Legendary', 'Mythic', 'Lair')),
	"Name"	TEXT NOT NULL DEFAULT 'X',
	"Description"	TEXT NOT NULL,
	"Points"	INTEGER NOT NULL,
	"IsRegional"	TEXT CHECK(Length("IsRegional") <= 1),
	PRIMARY KEY("EntityID","ActionID","Type"),
	FOREIGN KEY("EntityID") REFERENCES "Entity"("EntityID") ON DELETE CASCADE
)