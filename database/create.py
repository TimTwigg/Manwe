import sqlite3

con = sqlite3.connect("database/database.sqlite3")
cur = con.cursor()

languages = [
    "Common",
	"Dwarven",
	"Elvish",
    "Giant",
    "Gnomish",
    "Goblin",
    "Halfling",
    "Orc",
    "Abyssal",
    "Celestial",
    "Draconic",
    "Kraul",
    "Lodoxon",
    "Merfolk",
    "Minotaur",
    "Sphinx",
    "Sylvan",
    "Vedalken",
    "Deep Speech",
    "Thieves' Cant",
    "Primordial",
    "Undercommon",
    "Infernal",
    "Aquan",
    "Ignan",
    "Auran",
    "Terran",
    "Aarakocra",
    "Druidic",
    "Gith"
]

tuple_langs = [(lang,"") for lang in languages]

cur.executemany("INSERT INTO Languages (Language, Description) VALUES (?, ?)", tuple_langs)
con.commit()



