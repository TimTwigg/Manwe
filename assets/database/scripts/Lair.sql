CREATE TABLE "Lair" (
	"EntityID"	INTEGER NOT NULL,
	"Name"	TEXT,
	"Description"	TEXT,
	"Initiative"	INTEGER CHECK(Initiative >= 0),
	PRIMARY KEY("EntityID"),
	FOREIGN KEY("EntityID") REFERENCES "Entity"("EntityID") ON DELETE CASCADE
)