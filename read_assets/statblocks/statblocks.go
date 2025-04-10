package read_asset_statblocks

import (
	asset_utils "github.com/TimTwigg/EncounterManagerBackend/read_assets/utils"
	actions "github.com/TimTwigg/EncounterManagerBackend/types/actions"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
	dbutils "github.com/TimTwigg/EncounterManagerBackend/utils/database"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
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

func ReadStatBlockFromDB(name string) (stat_blocks.StatBlock, error) {
	// Read Entity information
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM Entity WHERE name = ?", name)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer rows.Close()
	var id int
	var block stat_blocks.StatBlock
	if rows.Next() {
		if err := rows.Scan(
			&id,
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
		logger.Error("No stat block found with name: " + name)
		return stat_blocks.StatBlock{}, nil
	}

	// Read Modifiers
	mod_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Type, Name, Value, Description FROM EntityModifiers WHERE EntityID = ?", id)
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
	lang_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Language, Description FROM EntityLanguages WHERE EntityID = ?", id)
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
	action_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT ActionID, Name, AttackType, HitModifier, Reach, Targets, Description FROM EntityActions WHERE EntityID = ?", id)
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

		// Create Damage array
		dmg_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Amount, Type, AltDmgActive, Amount2, Type2, AltDmgNote, AltDmgActive, Ability, DC, HalfDamage, SaveDmgNote FROM EntityActionDamage WHERE EntityID = ? and ActionID = ?", id, ActionID)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		defer dmg_rows.Close()
		// TODO

		// Add to StatBlock
		block.Actions = append(block.Actions, actions.Action{Name: Name, AttackType: AttackType, ToHitModifier: HitModifier, Reach: Reach, Targets: Targets, AdditionalDescription: Description})
	}

	return block, nil
}
