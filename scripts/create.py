

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

for language in languages:
    cleaned_language = ''.join(e for e in language if e.isalnum() or e in ' ').replace(' ', '_')
    with open(f'C:/Users/egan/Desktop/EncounterManagerBackend/assets/languages/{cleaned_language.lower()}.json', 'w') as file:
        file.write(f"{{\"language\": \"{language}\", \"Description\": \"\"}}")



