CREATE TABLE "SpokenLanguage" (
	"EntityID"	INTEGER NOT NULL,
	"Language"	TEXT NOT NULL,
	"Description"	TEXT,
	PRIMARY KEY("Language","EntityID"),
	FOREIGN KEY("EntityID") REFERENCES "Entity"("EntityID") ON DELETE CASCADE,
	FOREIGN KEY("Language") REFERENCES "Language"("Language") ON DELETE CASCADE
)