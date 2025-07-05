package assets

import (
	"strconv"
	"strings"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	actions "github.com/TimTwigg/Manwe/types/actions"
	generics "github.com/TimTwigg/Manwe/types/generics"
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

func ReadStatBlockByID(id int, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlock, error) {
	// Read StatBlock information
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT StatBlockID, Name, ChallengeRating, ProficiencyBonus, Source, Size, Type, Alignment, ArmorClass, HitPoints1, HitPoints2, SWalk, SFly, SClimb, SSwim, SBurrow, ArmorType FROM StatBlock WHERE StatBlockID = ? AND (Domain = 'Public' OR Domain = ? OR Published = 'X')"+asset_utils.StatBlockRestrictionClause(restriction, true), id, userid)
	if err != nil {
		logger.Error("Error querying database for Statblock: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer rows.Close()
	var block stat_blocks.StatBlock
	if rows.Next() {
		if err := rows.Scan(
			&block.ID,
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
			&block.Details.ArmorType,
		); err != nil {
			logger.Error("Error Scanning StatBlock Row: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
	block.Details.Skills = make([]generics.ProficiencyItem, 0)
	block.Details.SavingThrows = make([]generics.ProficiencyItem, 0)
	block.Details.Traits = make([]generics.SimpleItem, 0)
	block.Details.Languages.Languages = make([]string, 0)
	block.Details.Languages.Note = ""
	block.Stats.Abilities = make(map[string]int, 0)
	block.Actions = make([]actions.Action, 0)
	block.BonusActions = make([]generics.SimpleItem, 0)

	// Read Abilities
	ability_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Ability, Value FROM EntityStats WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Stats: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
	}
	defer ability_rows.Close()
	abilities := map[string]int{"Strength": 10, "Dexterity": 10, "Constitution": 10, "Wisdom": 10, "Charisma": 10, "Intelligence": 10}
	for ability_rows.Next() {
		var Ability string
		var Value int
		if err := ability_rows.Scan(
			&Ability,
			&Value,
		); err != nil {
			logger.Error("Error Scanning Ability Row: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}
		abilities[Ability] = Value
	}
	for ability, Value := range abilities {
		// Add to StatBlock
		block.Stats.Abilities[ability] = Value
	}

	// Read Modifiers
	mod_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Type, Name, Value, Description FROM Modifiers WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Modifiers: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
		case "SE":
			block.Details.Senses = append(block.Details.Senses, generics.NumericalItem{Name: Name, Modifier: Value})
		case "TR":
			block.Details.Traits = append(block.Details.Traits, generics.SimpleItem{Name: Name, Description: Description})
		default:
			logger.Error("Unsupported modifier type: " + ModType)
		}
	}

	// Read Proficiencies
	prof_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Type, Name, Level, Override FROM Proficiencies WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Proficiencies: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
	}
	defer prof_rows.Close()
	// Read row from Proficiencies table
	for prof_rows.Next() {
		var Type string
		var Name string
		var Level int
		var Override int
		if err := prof_rows.Scan(
			&Type,
			&Name,
			&Level,
			&Override,
		); err != nil {
			logger.Error("Error Scanning Proficiency Row: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}
		// Add to StatBlock
		switch Type {
		case "SK":
			block.Details.Skills = append(block.Details.Skills, generics.ProficiencyItem{Name: Name, Level: Level, Override: Override})
		case "ST":
			block.Details.SavingThrows = append(block.Details.SavingThrows, generics.ProficiencyItem{Name: Name, Level: Level, Override: Override})
		default:
			logger.Error("Unsupported proficiency type: " + Type)
		}
	}

	// Read Languages
	lang_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Language, Description FROM SpokenLanguage WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Languages: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
	action_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT ActionID, Name, AttackType, HitModifier, Reach, Targets, Description FROM Action WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Actions: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}

		var action = actions.Action{Name: Name, AttackType: AttackType, ToHitModifier: HitModifier, Reach: Reach, Targets: Targets, AdditionalDescription: Description}

		// Create Damage array
		dmg_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Amount, Type, AltDmgActive, Amount2, Type2, AltDmgNote, SaveDmgActive, Ability, DC, HalfDamage, SaveDmgNote FROM ActionDamage WHERE StatBlockID = ? and ActionID = ?", id, ActionID)
		if err != nil {
			logger.Error("Error querying database for Action Damage: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
				return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
	simple_actions_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Type, Name, Description FROM SimpleAction WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Simple Actions: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
	super_hdr_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Type, Description, Points FROM SuperActionHV WHERE StatBlockID = ?", id)
	if err != nil {
		logger.Error("Error querying database for Super Action Headers: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
		super_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Name, Description, Points FROM SuperActionV WHERE StatBlockID = ? and Type = ?", id, HType)
		if err != nil {
			logger.Error("Error querying database for Super Actions: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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
				return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
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

	if block.Lair, err = ReadLairByEntityID(id); err != nil {
		if !strings.HasPrefix(err.Error(), "No Lair found") {
			logger.Error("Error reading lair: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}
	}

	return block, nil
}

func ReadStatBlockByName(name string, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlock, error) {
	// Read StatBlock information
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT StatBlockID FROM StatBlock WHERE name = ? AND (Domain = 'Public' OR Domain = ? OR Published = 'X')"+asset_utils.StatBlockRestrictionClause(restriction, true), name, userid)
	if err != nil {
		logger.Error("Error querying database for StatBlock: " + err.Error())
		return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
	}
	defer rows.Close()
	var id int
	if rows.Next() {
		if err := rows.Scan(
			&id,
		); err != nil {
			logger.Error("Error Scanning Row: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}
	} else {
		logger.Error("No stat block found with name: " + name)
		return stat_blocks.StatBlock{}, errors.New("No stat block found with name: " + name)
	}
	return ReadStatBlockByID(id, userid, restriction)
}

func ReadStatBlockOverviewByID(id int, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlockOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT StatBlockID, Name, Type, Size, ChallengeRating, Source FROM StatBlock WHERE StatBlockID = ? AND (Domain = 'Public' OR Domain = ? OR Published = 'X')"+asset_utils.StatBlockRestrictionClause(restriction, true), id, userid)
	if err != nil {
		logger.Error("Error querying database for StatBlock Overview: " + err.Error())
		return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
	}
	defer rows.Close()
	var block stat_blocks.StatBlockOverview
	if rows.Next() {
		if err := rows.Scan(
			&block.ID,
			&block.Name,
			&block.Type,
			&block.Size,
			&block.ChallengeRating,
			&block.Source,
		); err != nil {
			logger.Error("Error Scanning StatBlock Row: " + err.Error())
			return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
		}
	} else {
		logger.Error("No stat block found with id: " + strconv.Itoa(id))
		return stat_blocks.StatBlockOverview{}, errors.New("No stat block found with id: " + strconv.Itoa(id))
	}

	return block, nil
}

func ReadStatBlockOverviewByName(name string, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlockOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT StatBlockID FROM StatBlock WHERE name = ? AND (Domain = 'Public' OR Domain = ? OR Published = 'X')"+asset_utils.StatBlockRestrictionClause(restriction, true), name, userid)
	if err != nil {
		logger.Error("Error querying database for StatBlock Overview: " + err.Error())
		return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
	}

	defer rows.Close()
	var id int
	if rows.Next() {
		if err := rows.Scan(
			&id,
		); err != nil {
			logger.Error("Error Scanning Row: " + err.Error())
			return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
		}
	} else {
		logger.Error("No stat block found with name: " + name)
		return stat_blocks.StatBlockOverview{}, errors.New("No stat block found with name: " + name)
	}

	return ReadStatBlockOverviewByID(id, userid, restriction)
}

func ReadAllStatBlockOverviews(userid string, restriction asset_utils.EntityTypeRestriction) ([]stat_blocks.StatBlockOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT StatBlockID, Name, Type, Size, ChallengeRating, Source FROM StatBlock WHERE (Domain = 'Public' OR Domain = ? OR Published = 'X')"+asset_utils.StatBlockRestrictionClause(restriction, true), userid)
	if err != nil {
		logger.Error("Error querying database for StatBlock Overviews: " + err.Error())
		return nil, error_utils.ParseError{Message: err.Error()}
	}
	defer rows.Close()
	var statblocks []stat_blocks.StatBlockOverview
	for rows.Next() {
		var ID string
		var Name string
		var Type string
		var Size string
		var CR int
		var Source string
		if err := rows.Scan(
			&ID,
			&Name,
			&Type,
			&Size,
			&CR,
			&Source,
		); err != nil {
			logger.Error("Error Scanning StatBlock Row: " + err.Error())
			return nil, error_utils.ParseError{Message: err.Error()}
		}
		statblocks = append(statblocks, stat_blocks.StatBlockOverview{ID: ID, Name: Name, Type: Type, Size: Size, ChallengeRating: CR, Source: Source})
	}
	return statblocks, nil
}

func ReadStatBlockByAccessType(accessType string, accessor string, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlock, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}
		return ReadStatBlockByID(id, userid, restriction)
	case "name":
		return ReadStatBlockByName(accessor, userid, restriction)
	default:
		logger.Error("Invalid access type: " + accessType)
		return stat_blocks.StatBlock{}, errors.New("Invalid access type: " + accessType)
	}
}

func ReadStatBlockOverviewByAccessType(accessType string, accessor string, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlockOverview, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
		}
		return ReadStatBlockOverviewByID(id, userid, restriction)
	case "name":
		return ReadStatBlockOverviewByName(accessor, userid, restriction)
	default:
		logger.Error("Invalid access type: " + accessType)
		return stat_blocks.StatBlockOverview{}, errors.New("Invalid access type: " + accessType)
	}
}
