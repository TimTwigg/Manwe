package assets

import (
	"context"
	"maps"
	"strconv"
	"strings"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	actions "github.com/TimTwigg/Manwe/types/actions"
	generics "github.com/TimTwigg/Manwe/types/generics"
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	logger "github.com/TimTwigg/Manwe/utils/log"
	pgx "github.com/jackc/pgx/v5"
	errors "github.com/pkg/errors"
)

func ReadStatBlockByID(id int, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlock, error) {
	// ################################################################################
	// Read StatBlock information
	// ################################################################################
	var block stat_blocks.StatBlock
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT statblockid, name, challengerating, proficiencybonus, source, size, type, alignment, armorclass, hitpoints1, hitpoints2, walkspeed, flyspeed, climbspeed, swimspeed, burrowspeed, armortype FROM public.statblock WHERE statblockid = $1 AND (username = 'public' OR username = $2 OR published = true)"+asset_utils.StatBlockRestrictionClause(restriction, true), id, userid).Scan(
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
	)
	if pgx.ErrNoRows == err {
		logger.Error("No Statblock found with ID: " + strconv.Itoa(id))
		return stat_blocks.StatBlock{}, errors.New("No Statblock found with ID: " + strconv.Itoa(id))
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error querying database for statblock")
	}

	// ################################################################################
	// Set default values
	// ################################################################################
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

	// ################################################################################
	// Read Abilities
	// ################################################################################
	ability_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT ability, value FROM public.statblockstats WHERE statblockid = $1", id)
	abilities := map[string]int{"Strength": 10, "Dexterity": 10, "Constitution": 10, "Wisdom": 10, "Charisma": 10, "Intelligence": 10}
	var Ability string
	var Value int
	_, err = pgx.ForEachRow(ability_rows, []any{&Ability, &Value}, func() error {
		abilities[Ability] = Value
		return nil
	})
	err = ability_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading abilities from database: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error reading abilities from database")
	}
	maps.Copy(block.Stats.Abilities, abilities)

	// ################################################################################
	// Read Modifiers
	// ################################################################################
	mod_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT type, name, value, description FROM public.modifiers WHERE statblockid = $1", id)
	var ModType, Name, Description string
	_, err = pgx.ForEachRow(mod_rows, []any{&ModType, &Name, &Value, &Description}, func() error {
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
		return nil
	})
	err = mod_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading modifiers from database: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error reading modifiers from database")
	}

	// ################################################################################
	// Read Proficiencies
	// ################################################################################
	prof_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT type, name, level, override FROM public.proficiencies WHERE statblockid = $1", id)
	var Type string
	var Level, Override int
	_, err = pgx.ForEachRow(prof_rows, []any{&Type, &Name, &Level, &Override}, func() error {
		switch Type {
		case "SK":
			block.Details.Skills = append(block.Details.Skills, generics.ProficiencyItem{Name: Name, Level: Level, Override: Override})
		case "ST":
			block.Details.SavingThrows = append(block.Details.SavingThrows, generics.ProficiencyItem{Name: Name, Level: Level, Override: Override})
		default:
			logger.Error("Unsupported proficiency type: " + Type)
		}
		return nil
	})
	err = prof_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading proficiencies from database: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error reading proficiencies from database")
	}

	// ################################################################################
	// Read Languages
	// ################################################################################
	lang_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT language, coalesce(description, '') FROM public.spokenlanguage WHERE statblockid = $1", id)
	var Language, Note string
	_, err = pgx.ForEachRow(lang_rows, []any{&Language, &Note}, func() error {
		switch Language {
		case "Note":
			block.Details.Languages.Note = Note
		default:
			block.Details.Languages.Languages = append(block.Details.Languages.Languages, Language)
		}
		return nil
	})
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error querying database for Languages: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error querying database for Languages")
	}

	// ################################################################################
	// Read Actions
	// ################################################################################
	action_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT actionid, name, attacktype, hitmodifier, reach, targets, coalesce(description, '') FROM public.action WHERE statblockid = $1", id)
	var ActionID, HitModifier, Reach, Targets int
	var AttackType string
	_, err = pgx.ForEachRow(action_rows, []any{&ActionID, &Name, &AttackType, &HitModifier, &Reach, &Targets, &Description}, func() error {
		var action = actions.Action{Name: Name, AttackType: AttackType, ToHitModifier: HitModifier, Reach: Reach, Targets: Targets, AdditionalDescription: Description}
		dmg_rows, _ := asset_utils.DBPool.Query(context.Background(), "SELECT amount, type, altdamageactive, amount2, type2, altdamagenote, savedamageactive, ability, dc, halfdamage, savedamagenote FROM public.actiondamagev WHERE statblockid = $1 and actionid = $2", id, ActionID)
		var Amount, Type, Amount2, Type2, AltDmgNote, Ability, SaveDmgNote string
		var AltDmgActive, SaveDmgActive, HalfDamage bool
		var DC int
		_, err = pgx.ForEachRow(dmg_rows, []any{&Amount, &Type, &AltDmgActive, &Amount2, &Type2, &AltDmgNote, &SaveDmgActive, &Ability, &DC, &HalfDamage, &SaveDmgNote}, func() error {
			// Check if AltDmgActive is true
			if !AltDmgActive {
				Amount2 = ""
				Type2 = ""
				AltDmgNote = ""
			}
			// Check if SaveDmgActive is true
			if !SaveDmgActive {
				Ability = ""
				DC = 0
				HalfDamage = false
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
					HalfDamage: HalfDamage,
					Note:       SaveDmgNote,
				},
			})
			block.Actions = append(block.Actions, action)
			return nil
		})
		err = dmg_rows.Err()
		if err != nil && err != pgx.ErrNoRows {
			logger.Error("Error reading action damage from database: " + err.Error())
			return errors.Wrap(err, "Error reading action damage from database")
		}
		return nil
	})
	err = action_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading actions from database: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error reading actions from database")
	}

	// ################################################################################
	// Read Bonus Actions and Reactions
	// ################################################################################
	simple_actions_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT type, name, description FROM public.simpleaction WHERE statblockid = $1", id)
	_, err = pgx.ForEachRow(simple_actions_rows, []any{&Type, &Name, &Description}, func() error {
		switch Type {
		case "Bonus":
			block.BonusActions = append(block.BonusActions, generics.SimpleItem{Name: Name, Description: Description})
		case "Reaction":
			block.Reactions = append(block.Reactions, generics.SimpleItem{Name: Name, Description: Description})
		default:
			logger.Error("Unknown simple action type: " + Type)
		}
		return nil
	})
	err = simple_actions_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading simple actions from database: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error reading simple actions from database")
	}

	// ################################################################################
	// Read Legendary Actions
	// ################################################################################
	super_hdr_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT type, description, points FROM public.superactionhv WHERE statblockid = $1", id)
	var HType, HDescription string
	var HPoints int
	_, err = pgx.ForEachRow(super_hdr_rows, []any{&HType, &HDescription, &HPoints}, func() error {
		switch HType {
		case "Legendary":
			block.LegendaryActions = actions.Legendary{Description: HDescription, Points: HPoints, Actions: make([]actions.LegendaryAction, 0)}
		case "Mythic":
			block.MythicActions = actions.Mythic{Description: HDescription, Actions: make([]actions.MythicAction, 0)}
		default:
			logger.Error("Unknown legendary action header type: " + HType)
		}

		super_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT name, description, points FROM public.superactionv WHERE statblockid = $1 and type = $2", id, HType)
		var Name, Description string
		var Points int
		_, err = pgx.ForEachRow(super_rows, []any{&Name, &Description, &Points}, func() error {
			switch HType {
			case "Legendary":
				block.LegendaryActions.Actions = append(block.LegendaryActions.Actions, actions.LegendaryAction{Name: Name, Description: Description, Cost: Points})
			case "Mythic":
				block.MythicActions.Actions = append(block.MythicActions.Actions, actions.MythicAction{Name: Name, Description: Description, Cost: Points})
			default:
				logger.Error("Unknown legendary action type: " + HType)
			}
			return nil
		})
		err = super_rows.Err()
		if err != nil && err != pgx.ErrNoRows {
			logger.Error("Error reading super actions from database: " + err.Error())
			return errors.Wrap(err, "Error reading super actions from database")
		}
		return nil
	})
	err = super_hdr_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error querying database for Super Action Headers: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error querying database for Super Action Headers")
	}

	// ################################################################################
	// Read Lair
	// ################################################################################
	if block.Lair, err = ReadLairByEntityID(id, true); err != nil {
		if !strings.HasPrefix(err.Error(), "No Lair found") {
			logger.Error("Error reading lair: " + err.Error())
			return stat_blocks.StatBlock{}, error_utils.ParseError{Message: err.Error()}
		}
	}

	return block, nil
}

func ReadStatBlockByName(name string, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlock, error) {
	var id int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT statblockid FROM public.statblock WHERE name = $1 AND (username = 'public' OR username = $2 OR published = true)"+asset_utils.StatBlockRestrictionClause(restriction, true), name, userid).Scan(&id)
	if pgx.ErrNoRows == err {
		logger.Error("No StatBlock found with name: " + name)
		return stat_blocks.StatBlock{}, errors.New("No StatBlock found with name: " + name)
	} else if err != nil {
		logger.Error("Error querying database for StatBlock: " + err.Error())
		return stat_blocks.StatBlock{}, errors.Wrap(err, "Error querying database for StatBlock")
	}
	return ReadStatBlockByID(id, userid, restriction)
}

func ReadStatBlockOverviewByID(id int, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlockOverview, error) {
	var block stat_blocks.StatBlockOverview
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT statblockid, name, type, size, challengerating, source FROM public.statblock WHERE statblockid = $1 AND (username = 'public' OR username = $2 OR published = true)"+asset_utils.StatBlockRestrictionClause(restriction, true), id, userid).Scan(
		&block.ID,
		&block.Name,
		&block.Type,
		&block.Size,
		&block.ChallengeRating,
		&block.Source,
	)
	if pgx.ErrNoRows == err {
		logger.Error("No StatBlock Overview found with ID: " + strconv.Itoa(id))
		return stat_blocks.StatBlockOverview{}, errors.New("No StatBlock Overview found with ID: " + strconv.Itoa(id))
	} else if err != nil {
		logger.Error("Error querying database for StatBlock Overview: " + err.Error())
		return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
	}
	return block, nil
}

func ReadStatBlockOverviewByName(name string, userid string, restriction asset_utils.EntityTypeRestriction) (stat_blocks.StatBlockOverview, error) {
	var id int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT statblockid FROM public.statblock WHERE name = $1 AND (username = 'public' OR username = $2 OR published = true)"+asset_utils.StatBlockRestrictionClause(restriction, true), name, userid).Scan(&id)
	if pgx.ErrNoRows == err {
		logger.Error("No StatBlock Overview found with Name: " + name)
		return stat_blocks.StatBlockOverview{}, errors.New("No StatBlock Overview found with Name: " + name)
	} else if err != nil {
		logger.Error("Error querying database for StatBlock Overview: " + err.Error())
		return stat_blocks.StatBlockOverview{}, error_utils.ParseError{Message: err.Error()}
	}
	return ReadStatBlockOverviewByID(id, userid, restriction)
}

func ReadAllStatBlockOverviews(userid string, restriction asset_utils.EntityTypeRestriction) ([]stat_blocks.StatBlockOverview, error) {
	rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT statblockid, name, type, size, challengerating, source FROM public.statblock WHERE (username = 'public' OR username = $1 OR published = true)"+asset_utils.StatBlockRestrictionClause(restriction, true), userid)
	statblocks, _ := pgx.CollectRows(rows, func(row pgx.CollectableRow) (stat_blocks.StatBlockOverview, error) {
		var block stat_blocks.StatBlockOverview
		if err := row.Scan(
			&block.ID,
			&block.Name,
			&block.Type,
			&block.Size,
			&block.ChallengeRating,
			&block.Source,
		); err != nil {
			logger.Error("Error Scanning Row: " + err.Error())
			return stat_blocks.StatBlockOverview{}, errors.Wrap(err, "Error scanning StatBlock Overview row")
		}
		return block, nil
	})
	err = rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading StatBlock Overviews from database: " + err.Error())
		return nil, errors.Wrap(err, "Error reading StatBlock Overviews from database")
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
