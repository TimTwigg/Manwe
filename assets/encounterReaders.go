package assets

import (
	"strconv"
	"strings"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	encounters "github.com/TimTwigg/Manwe/types/encounters"
	entities "github.com/TimTwigg/Manwe/types/entities"
	generics "github.com/TimTwigg/Manwe/types/generics"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	utils "github.com/TimTwigg/Manwe/utils/functions"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

func ReadEncounterByID(id int, userid string) (encounters.Encounter, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID, Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn, HasLair, LairOwnerID, ActiveID FROM Encounter WHERE EncounterID = ? AND (Domain = 'Public' OR Domain = ?)", id, userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	defer rows.Close()

	var encounter encounters.Encounter

	encounter.Entities = make([]entities.Entity, 0)
	encounter.Lair.Actions.Items = make([]generics.SimpleItem, 0)
	encounter.Lair.RegionalEffects.Items = make([]generics.SimpleItem, 0)

	if rows.Next() {
		var CreationDate, AccessedDate, Started, HasLair string
		if err := rows.Scan(
			&encounter.ID,
			&encounter.Name,
			&encounter.Description,
			&CreationDate,
			&AccessedDate,
			&encounter.Metadata.Campaign,
			&Started,
			&encounter.Metadata.Round,
			&encounter.Metadata.Turn,
			&HasLair,
			&encounter.LairOwnerID,
			&encounter.ActiveID,
		); err != nil {
			logger.Error("Error Scanning Encounter Row: " + err.Error())
			return encounters.Encounter{}, err
		}
		encounter.Metadata.CreationDate = utils.ParseStringDate(CreationDate)
		encounter.Metadata.AccessedDate = utils.ParseStringDate(AccessedDate)
		encounter.Metadata.Started = Started == "X"
		encounter.HasLair = HasLair == "X"
	} else {
		logger.Error("No Encounter found with id: " + strconv.Itoa(id))
		return encounters.Encounter{}, errors.New("No Encounter found with id: " + strconv.Itoa(id))
	}

	entity_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT RowID, StatBlockID, Suffix, Initiative, MaxHitPoints, TempHitPoints, CurrentHitPoints, ArmorClassBonus, Concentration, Notes, IsHostile, EncounterLocked, ID FROM EncounterEntities WHERE EncounterID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	defer entity_rows.Close()
	for entity_rows.Next() {
		var rowID, entityID, initiative, maxHitPoints, tempHitPoints, currentHitPoints, armorClassBonus int
		var suffix, notes, isHostile, encounterLocked, concentration, ID string

		if err := entity_rows.Scan(
			&rowID,
			&entityID,
			&suffix,
			&initiative,
			&maxHitPoints,
			&tempHitPoints,
			&currentHitPoints,
			&armorClassBonus,
			&concentration,
			&notes,
			&isHostile,
			&encounterLocked,
			&ID,
		); err != nil {
			logger.Error("Error Scanning Encounter Entity Row: " + err.Error())
			return encounters.Encounter{}, err
		}

		statblock, err := ReadStatBlockByID(entityID, userid, asset_utils.ANY)
		if err != nil {
			logger.Error("Error reading statblock: " + err.Error())
			return encounters.Encounter{}, err
		}

		entity := entities.Entity{
			DBID:             entityID,
			ID:               ID,
			Name:             statblock.Name,
			Suffix:           suffix,
			Initiative:       initiative,
			MaxHitPoints:     maxHitPoints,
			TempHitPoints:    tempHitPoints,
			CurrentHitPoints: currentHitPoints,
			ArmorClass:       statblock.Stats.ArmorClass,
			ArmorClassBonus:  armorClassBonus,
			Speed:            statblock.Stats.Speed,
			Conditions:       make(map[string]int, 0),
			SpellSaveDC:      statblock.Details.SpellSaveDC,
			SpellSlots:       make(map[int]entities.SpellSlotLevel, 0),
			Concentration:    concentration == "X",
			Notes:            notes,
			IsHostile:        isHostile == "X",
			EncounterLocked:  encounterLocked == "X",
			Displayable:      statblock,
			EntityType:       entities.StatBlock,
			SavingThrows:     statblock.Details.SavingThrows,
			ChallengeRating:  statblock.ChallengeRating,
		}

		conditions_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Condition, Duration FROM EncEntConditions WHERE EncounterID = ? and RowID = ?", id, rowID)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return encounters.Encounter{}, err
		}
		defer conditions_rows.Close()
		for conditions_rows.Next() {
			var condition string
			var duration int
			if err := conditions_rows.Scan(&condition, &duration); err != nil {
				logger.Error("Error Scanning Encounter Entity Condition Row: " + err.Error())
				return encounters.Encounter{}, err
			}
			entity.Conditions[condition] = duration
		}

		encounter.Entities = append(encounter.Entities, entity)
	}

	if encounter.HasLair && encounter.LairOwnerID > 0 {
		if encounter.Lair, err = ReadLairByEntityID(encounter.LairOwnerID); err != nil {
			if !strings.HasPrefix(err.Error(), "No Lair found") {
				logger.Error("Error reading lair: " + err.Error())
				return encounters.Encounter{}, error_utils.ParseError{Message: err.Error()}
			}
		}
	}

	return encounter, nil
}

func ReadEncounterByName(name string, userid string) (encounters.Encounter, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID FROM Encounter WHERE name = ? AND (Domain = 'Public' OR Domain = ?)", name, userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	defer rows.Close()

	var id int
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			logger.Error("Error Scanning Encounter Row: " + err.Error())
			return encounters.Encounter{}, err
		}
	} else {
		logger.Error("No Encounter found with name: " + name)
		return encounters.Encounter{}, errors.New("No Encounter found with name: " + name)
	}

	return ReadEncounterByID(id, userid)
}

func ReadEncounterOverviewByID(id int, userid string) (encounters.EncounterOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn FROM Encounter WHERE EncounterID = ? AND (Domain = 'Public' OR Domain = ?)", id, userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.EncounterOverview{}, err
	}
	defer rows.Close()

	var encounter encounters.EncounterOverview

	if rows.Next() {
		var CreationDate, AccessedDate, Started string
		if err := rows.Scan(
			&encounter.Name,
			&encounter.Description,
			&CreationDate,
			&AccessedDate,
			&encounter.Metadata.Campaign,
			&Started,
			&encounter.Metadata.Round,
			&encounter.Metadata.Turn,
		); err != nil {
			logger.Error("Error Scanning Encounter Row: " + err.Error())
			return encounters.EncounterOverview{}, err
		}
		encounter.Metadata.CreationDate = utils.ParseStringDate(CreationDate)
		encounter.Metadata.AccessedDate = utils.ParseStringDate(AccessedDate)
		encounter.Metadata.Started = Started == "X"
	} else {
		logger.Error("No Encounter found with id: " + strconv.Itoa(id))
		return encounters.EncounterOverview{}, errors.New("No Encounter found with id: " + strconv.Itoa(id))
	}

	return encounter, nil
}

func ReadEncounterOverviewByName(name string, userid string) (encounters.EncounterOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID FROM Encounter WHERE Name = ? AND (Domain = 'Public' OR Domain = ?)", name, userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.EncounterOverview{}, err
	}
	defer rows.Close()
	var id int

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			logger.Error("Error Scanning Encounter Row: " + err.Error())
			return encounters.EncounterOverview{}, err
		}
	} else {
		logger.Error("No Encounter found with name: " + name)
		return encounters.EncounterOverview{}, errors.New("No Encounter found with name: " + name)
	}

	return ReadEncounterOverviewByID(id, userid)
}

func ReadAllEncounterOverviews(userid string) ([]encounters.EncounterOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID, Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn FROM Encounter WHERE (Domain = 'Public' OR Domain = ?)", userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var encs []encounters.EncounterOverview = make([]encounters.EncounterOverview, 0)

	for rows.Next() {
		var Name, Description, Campaign, CreationDate, AccessedDate, Started string
		var Round, Turn, id int
		if err := rows.Scan(
			&id,
			&Name,
			&Description,
			&CreationDate,
			&AccessedDate,
			&Campaign,
			&Started,
			&Round,
			&Turn,
		); err != nil {
			logger.Error("Error Scanning Encounter Row: " + err.Error())
			return nil, err
		}
		encounter := encounters.EncounterOverview{
			ID:          id,
			Name:        Name,
			Description: Description,
			Metadata: encounters.EncounterMetadata{
				Campaign: Campaign,
				Started:  Started == "X",
				Round:    Round,
				Turn:     Turn,
			},
		}
		encounter.Metadata.CreationDate = utils.ParseStringDate(CreationDate)
		encounter.Metadata.AccessedDate = utils.ParseStringDate(AccessedDate)
		encs = append(encs, encounter)
	}

	return encs, nil
}

func ReadEncounterByAccessType(accessType string, accessor string, userid string) (encounters.Encounter, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return encounters.Encounter{}, err
		}
		return ReadEncounterByID(id, userid)
	case "name":
		return ReadEncounterByName(accessor, userid)
	default:
		logger.Error("Invalid access type: " + accessType)
		return encounters.Encounter{}, errors.New("Invalid access type: " + accessType)
	}
}

func ReadEncounterOverviewByAccessType(accessType string, accessor string, userid string) (encounters.EncounterOverview, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return encounters.EncounterOverview{}, err
		}
		return ReadEncounterOverviewByID(id, userid)
	case "name":
		return ReadEncounterOverviewByName(accessor, userid)
	default:
		logger.Error("Invalid access type: " + accessType)
		return encounters.EncounterOverview{}, errors.New("Invalid access type: " + accessType)
	}
}
