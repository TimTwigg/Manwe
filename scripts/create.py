from pathlib import Path
import json

damages = [
    "Cold",
	"Lightning",
	"Radiant",
	"Acid",
	"Bludgeoning",
	"Force",
	"Psychic",
	"Necrosis",
	"Piercing",
	"Poison",
	"Fire",
	"Slashing",
	"Thunder",
]

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
    "Gith",
]


############################################################################################################
# CAUTION: This will overwrite the existing files                                                          #
############################################################################################################

# for damage in damages:
#     with Path("./../assets/damage_types").joinpath(f"{damage.lower()}.json").open("w") as file:
#         json.dump({"DamageType": damage, "Description": ""}, file, indent = 4)

# for lang in languages:
#     with Path("./../assets/languages").joinpath(f"{lang.lower().replace(" ", "_").replace("'", "")}.json").open("w") as file:
#         json.dump({"Language": lang, "Description": ""}, file, indent = 4)

