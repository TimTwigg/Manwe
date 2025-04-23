package assets

import (
	"strconv"

	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	actions "github.com/TimTwigg/EncounterManagerBackend/types/actions"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
	dbutils "github.com/TimTwigg/EncounterManagerBackend/utils/database"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	errors "github.com/pkg/errors"
)

// Read a stat block from database
func ReadStatBlock(name string) (stat_blocks.StatBlock, error) {
	if asset_utils.StatBlockExists(name) {
		data, err := asset_utils.ReadAsset(name, "stat_blocks")
		if err != nil {
			logger.Error("Error reading stat block: " + name + ": " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		statblock, err := stat_blocks.ParseStatBlockData(data)
		if err != nil {
			return stat_blocks.StatBlock{}, err
		}
		return statblock.(stat_blocks.StatBlock), nil
	} else {
		logger.Error("StatBlock " + name + " Not Found")
	}
	return stat_blocks.StatBlock{}, nil
}

func ReadStatBlockByID(id int) (stat_blocks.StatBlock, error) {
	// Read Entity information
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM Entity WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer rows.Close()
	var t int
	var block stat_blocks.StatBlock
	if rows.Next() {
		if err := rows.Scan(
			&t,
			&block.Name,
			&block.ChallengeRating,
			&block.ProficiencyBonus,
			&block.Source,
			&block.Description.Size,
			&block.Description.Type,
			&block.Description.Alignment,
			&block.Stats.ArmorClass,
			&block.Stats.HitPoints.Average,
			&block.Stats.HitPoints.Dice,
			&block.Stats.Speed.Walk,
			&block.Stats.Speed.Fly,
			&block.Stats.Speed.Climb,
			&block.Stats.Speed.Swim,
			&block.Stats.Speed.Burrow,
			&block.Stats.ReactionCount,
			&block.Stats.Strength,
			&block.Stats.Dexterity,
			&block.Stats.Constitution,
			&block.Stats.Intelligence,
			&block.Stats.Wisdom,
			&block.Stats.Charisma,
			&block.Details.ArmorType,
		); err != nil {
			logger.Error("Error Scanning Entity Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
	} else {
		logger.Error("No stat block found with id: " + strconv.Itoa(id))
		return stat_blocks.StatBlock{}, errors.New("No stat block found with id: " + strconv.Itoa(id))
	}

	// Set default values
	block.DamageModifiers.Immunities = make([]string, 0)
	block.DamageModifiers.Resistances = make([]string, 0)
	block.DamageModifiers.Vulnerabilities = make([]string, 0)
	block.ConditionImmunities = make([]string, 0)
	block.Details.Senses = make([]generics.NumericalItem, 0)
	block.Details.Skills = make([]generics.NumericalItem, 0)
	block.Details.SavingThrows = make([]generics.NumericalItem, 0)
	block.Details.Traits = make([]generics.SimpleItem, 0)
	block.Details.Languages.Languages = make([]string, 0)
	block.Details.Languages.Note = ""
	block.Actions = make([]actions.Action, 0)
	block.BonusActions = make([]generics.SimpleItem, 0)
	block.Reactions = make([]generics.SimpleItem, 0)

	// Read Modifiers
	mod_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Type, Name, Value, Description FROM ModifierV WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer mod_rows.Close()
	// Read row from Modifiers table
	for mod_rows.Next() {
		var ModType string
		var Name string
		var Value int
		var Description string
		if err := mod_rows.Scan(
			&ModType,
			&Name,
			&Value,
			&Description,
		); err != nil {
			logger.Error("Error Scanning Modifier Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}

		// Add to StatBlock
		switch ModType {
		case "DI":
			block.DamageModifiers.Immunities = append(block.DamageModifiers.Immunities, Name)
		case "DR":
			block.DamageModifiers.Resistances = append(block.DamageModifiers.Resistances, Name)
		case "DV":
			block.DamageModifiers.Vulnerabilities = append(block.DamageModifiers.Vulnerabilities, Name)
		case "CI":
			block.ConditionImmunities = append(block.ConditionImmunities, Name)
		case "SK":
			block.Details.Skills = append(block.Details.Skills, generics.NumericalItem{Name: Name, Modifier: Value})
		case "ST":
			block.Details.SavingThrows = append(block.Details.SavingThrows, generics.NumericalItem{Name: Name, Modifier: Value})
		case "SE":
			block.Details.Senses = append(block.Details.Senses, generics.NumericalItem{Name: Name, Modifier: Value})
		case "TR":
			block.Details.Traits = append(block.Details.Traits, generics.SimpleItem{Name: Name, Description: Description})
		default:
			logger.Error("Unknown modifier type: " + ModType)
		}
	}

	// Read Languages
	lang_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Language, Description FROM SpokenLanguageV WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer lang_rows.Close()
	// Read row from Languages table
	for lang_rows.Next() {
		var Language string
		var Note string
		if err := lang_rows.Scan(
			&Language,
			&Note,
		); err != nil {
			logger.Error("Error Scanning Language Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}

		// Add to StatBlock
		switch Language {
		case "Note":
			block.Details.Languages.Note = Note
		default:
			block.Details.Languages.Languages = append(block.Details.Languages.Languages, Language)
		}
	}

	// Read Actions
	action_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT ActionID, Name, AttackType, HitModifier, Reach, Targets, Description FROM ActionV WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer action_rows.Close()
	// Read row from Actions table
	for action_rows.Next() {
		var ActionID int
		var Name string
		var AttackType string
		var HitModifier int
		var Reach int
		var Targets int
		var Description string
		if err := action_rows.Scan(
			&ActionID,
			&Name,
			&AttackType,
			&HitModifier,
			&Reach,
			&Targets,
			&Description,
		); err != nil {
			logger.Error("Error Scanning Action Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}

		var action = actions.Action{Name: Name, AttackType: AttackType, ToHitModifier: HitModifier, Reach: Reach, Targets: Targets, AdditionalDescription: Description}

		// Create Damage array
		dmg_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Amount, Type, AltDmgActive, Amount2, Type2, AltDmgNote, SaveDmgActive, Ability, DC, HalfDamage, SaveDmgNote FROM ActionDamageV WHERE EntityID = ? and ActionID = ?", id, ActionID)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		defer dmg_rows.Close()
		for dmg_rows.Next() {
			var Amount string
			var Type string
			var AltDmgActive string
			var Amount2 string
			var Type2 string
			var AltDmgNote string
			var SaveDmgActive string
			var Ability string
			var DC int
			var HalfDamage string
			var SaveDmgNote string
			if err := dmg_rows.Scan(
				&Amount,
				&Type,
				&AltDmgActive,
				&Amount2,
				&Type2,
				&AltDmgNote,
				&SaveDmgActive,
				&Ability,
				&DC,
				&HalfDamage,
				&SaveDmgNote,
			); err != nil {
				logger.Error("Error Scanning Action Damage Row: " + err.Error())
				return stat_blocks.StatBlock{}, err
			}

			// Check if AltDmgActive is true
			if AltDmgActive != "X" {
				Amount2 = ""
				Type2 = ""
				AltDmgNote = ""
			}
			// Check if SaveDmgActive is true
			if SaveDmgActive != "X" {
				Ability = ""
				DC = 0
				HalfDamage = ""
				SaveDmgNote = ""
			}

			// Add to Action
			action.Damage = append(action.Damage, actions.DamageT{
				Amount: Amount,
				Type:   Type,
				AlternativeDamage: actions.AltDamageT{
					Amount: Amount2,
					Type:   Type2,
					Note:   AltDmgNote,
				},
				SavingThrow: actions.SavingThrowDamageT{
					Ability:    Ability,
					DC:         DC,
					HalfDamage: HalfDamage == "X",
					Note:       SaveDmgNote,
				},
			})
		}

		// Add to StatBlock
		block.Actions = append(block.Actions, action)
	}

	// Read Bonus Actions and Reactions
	simple_actions_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Type, Name, Description FROM SimpleAction WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer simple_actions_rows.Close()
	// Read row from SimpleActions table
	for simple_actions_rows.Next() {
		var Type string
		var Name string
		var Description string
		if err := simple_actions_rows.Scan(
			&Type,
			&Name,
			&Description,
		); err != nil {
			logger.Error("Error Scanning Simple Action Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		// Add to StatBlock
		switch Type {
		case "Bonus":
			block.BonusActions = append(block.BonusActions, generics.SimpleItem{Name: Name, Description: Description})
		case "Reaction":
			block.Reactions = append(block.Reactions, generics.SimpleItem{Name: Name, Description: Description})
		default:
			logger.Error("Unknown simple action type: " + Type)
		}
	}

	// Read Legendary Actions
	super_hdr_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Type, Description, Points FROM SuperActionHV WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer super_hdr_rows.Close()
	// Read row from LegendaryActions table
	for super_hdr_rows.Next() {
		var HType string
		var HDescription string
		var HPoints int
		if err := super_hdr_rows.Scan(
			&HType,
			&HDescription,
			&HPoints,
		); err != nil {
			logger.Error("Error Scanning Super Action Header Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}

		// Add to StatBlock
		switch HType {
		case "Legendary":
			block.LegendaryActions = actions.Legendary{Description: HDescription, Points: HPoints, Actions: make([]actions.LegendaryAction, 0)}
		case "Mythic":
			block.MythicActions = actions.Mythic{Description: HDescription, Actions: make([]actions.MythicAction, 0)}
		default:
			logger.Error("Unknown legendary action header type: " + HType)
		}

		// Read Super Actions
		super_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Name, Description, Points FROM SuperActionV WHERE EntityID = ? and Type = ?", id, HType)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		defer super_rows.Close()
		// Read row from SuperActions table
		for super_rows.Next() {
			var Name string
			var Description string
			var Points int
			if err := super_rows.Scan(
				&Name,
				&Description,
				&Points,
			); err != nil {
				logger.Error("Error Scanning Super Action Row: " + err.Error())
				return stat_blocks.StatBlock{}, err
			}
			// Add to StatBlock
			switch HType {
			case "Legendary":
				block.LegendaryActions.Actions = append(block.LegendaryActions.Actions, actions.LegendaryAction{Name: Name, Description: Description, Cost: Points})
			case "Mythic":
				block.MythicActions.Actions = append(block.MythicActions.Actions, actions.MythicAction{Name: Name, Description: Description, Cost: Points})
			default:
				logger.Error("Unknown legendary action type: " + HType)
			}
		}
	}

	// Read Lair
	lair_row, err := dbutils.QuerySQL(dbutils.DB, "SELECT Description, Initiative FROM Lair WHERE EntityID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer lair_row.Close()
	// Read row from Lair table
	if lair_row.Next() {
		var Description string
		var Initiative int
		if err := lair_row.Scan(
			&Description,
			&Initiative,
		); err != nil {
			logger.Error("Error Scanning Lair Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		block.Lair = stat_blocks.Lair{Name: block.Name, Description: Description, Initiative: Initiative, Actions: generics.ItemList{Description: "", Items: make([]generics.SimpleItem, 0)}, RegionalEffects: generics.ItemList{Description: "", Items: make([]generics.SimpleItem, 0)}}

		// Read Lair Actions
		lair_actions_row, err := dbutils.QuerySQL(dbutils.DB, "SELECT Name, Description, IsRegional FROM LairActionV WHERE EntityID = ?", id)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		defer lair_actions_row.Close()
		// Read row from LairActions table
		for lair_actions_row.Next() {
			var Name string
			var Description string
			var IsRegional string
			if err := lair_actions_row.Scan(
				&Name,
				&Description,
				&IsRegional,
			); err != nil {
				logger.Error("Error Scanning Lair Action Row: " + err.Error())
				return stat_blocks.StatBlock{}, err
			}
			// Add to StatBlock
			if Name == "X" {
				if IsRegional == "X" {
					block.Lair.RegionalEffects.Description = Description
				} else {
					block.Lair.Actions.Description = Description
				}
			} else {
				if IsRegional == "X" {
					block.Lair.RegionalEffects.Items = append(block.Lair.RegionalEffects.Items, generics.SimpleItem{Name: Name, Description: Description})
				} else {
					block.Lair.Actions.Items = append(block.Lair.Actions.Items, generics.SimpleItem{Name: Name, Description: Description})
				}
			}
		}
	}

	return block, nil
}

func ReadStatBlockByName(name string) (stat_blocks.StatBlock, error) {
	// Read Entity information
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT EntityID FROM Entity WHERE name = ?", name)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer rows.Close()
	var id int
	if rows.Next() {
		if err := rows.Scan(
			&id,
		); err != nil {
			logger.Error("Error Scanning Row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
	} else {
		logger.Error("No stat block found with name: " + name)
		return stat_blocks.StatBlock{}, errors.New("No stat block found with name: " + name)
	}
	return ReadStatBlockByID(id)
}

func ReadStatBlockOverviewFromDB(name string) (stat_blocks.StatBlockOverview, error) {
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM EntityOverviewV WHERE name = ?", name)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlockOverview{}, err
	}
	defer rows.Close()
	var Name string
	var Type string
	var Size string
	var CR int
	var Source string
	if rows.Next() {
		if err := rows.Scan(
			&Name,
			&Type,
			&Size,
			&CR,
			&Source,
		); err != nil {
			logger.Error("Error Scanning Entity Row: " + err.Error())
			return stat_blocks.StatBlockOverview{}, err
		}
	} else {
		logger.Error("No stat block found with name: " + name)
		return stat_blocks.StatBlockOverview{}, errors.New("No stat block found with name: " + name)
	}

	return stat_blocks.StatBlockOverview{
		Name:            Name,
		Type:            Type,
		Size:            Size,
		ChallengeRating: CR,
		Source:          Source,
	}, nil
}
