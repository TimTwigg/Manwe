package assets

import (
	"strconv"

	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	errors "github.com/pkg/errors"
)

func ReadEncounterByID(id int) (encounters.Encounter, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID, Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn, HasLair, LairEntityName, ActiveID FROM Encounter WHERE EncounterID = ?", id)
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
			&encounter.LairEntityName,
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

	entity_rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT RowID, EntityID, Suffix, Initiative, MaxHitPoints, TempHitPoints, CurrentHitPoints, ArmorClassBonus, Notes, IsHostile, EncounterLocked, ID FROM EncounterEntities WHERE EncounterID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	defer entity_rows.Close()
	for entity_rows.Next() {
		var rowID, entityID, initiative, maxHitPoints, tempHitPoints, currentHitPoints, armorClassBonus int
		var suffix, notes, isHostile, encounterLocked, ID string

		if err := entity_rows.Scan(
			&rowID,
			&entityID,
			&suffix,
			&initiative,
			&maxHitPoints,
			&tempHitPoints,
			&currentHitPoints,
			&armorClassBonus,
			&notes,
			&isHostile,
			&encounterLocked,
			&ID,
		); err != nil {
			logger.Error("Error Scanning Encounter Entity Row: " + err.Error())
			return encounters.Encounter{}, err
		}

		statblock, err := ReadStatBlockByID(entityID)
		if err != nil {
			logger.Error("Error reading statblock: " + err.Error())
			return encounters.Encounter{}, err
		}

		entity := entities.Entity{
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
			Concentration:    false,
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

	return encounter, nil
}

func ReadEncounterByName(name string) (encounters.Encounter, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID FROM Encounter WHERE name = ?", name)
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

	return ReadEncounterByID(id)
}

func ReadEncounterOverviewByID(id int) (encounters.EncounterOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn FROM Encounter WHERE EncounterID = ?", id)
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

func ReadEncounterOverviewByName(name string) (encounters.EncounterOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID FROM Encounter WHERE Name = ?", name)
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

	return ReadEncounterOverviewByID(id)
}

func ReadAllEncounterOverviews() ([]encounters.EncounterOverview, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT EncounterID, Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn FROM Encounter")
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

func SetEncounter(encounter encounters.Encounter) (encounters.Encounter, error) {
	if encounter.ID == 0 {
		res, err := asset_utils.ExecSQL(
			asset_utils.DB,
			"INSERT INTO Encounter (Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn, HasLair, LairEntityName, ActiveID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			encounter.Name,
			encounter.Description,
			utils.FormatDate(encounter.Metadata.CreationDate),
			utils.FormatDate(encounter.Metadata.AccessedDate),
			encounter.Metadata.Campaign,
			utils.FormatBool(encounter.Metadata.Started),
			encounter.Metadata.Round,
			encounter.Metadata.Turn,
			utils.FormatBool(encounter.HasLair),
			encounter.LairEntityName,
			encounter.ActiveID,
		)
		if err != nil {
			logger.Error("Error inserting Encounter: " + err.Error())
			return encounters.Encounter{}, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			logger.Error("Error getting last insert ID: " + err.Error())
			return encounters.Encounter{}, err
		}
		encounter.ID = int(id)
		return encounter, nil
	} else {
		_, err := asset_utils.ExecSQL(
			asset_utils.DB,
			"UPDATE Encounter SET Name = ?, Description = ?, CreationDate = ?, AccessedDate = ?, Campaign = ?, Started = ?, Round = ?, Turn = ?, HasLair = ?, LairEntityName = ?, ActiveID = ? WHERE EncounterID = ?",
			encounter.Name,
			encounter.Description,
			utils.FormatDate(encounter.Metadata.CreationDate),
			utils.FormatDate(encounter.Metadata.AccessedDate),
			encounter.Metadata.Campaign,
			utils.FormatBool(encounter.Metadata.Started),
			encounter.Metadata.Round,
			encounter.Metadata.Turn,
			utils.FormatBool(encounter.HasLair),
			encounter.LairEntityName,
			encounter.ActiveID,
			encounter.ID,
		)
		if err != nil {
			logger.Error("Error updating Encounter: " + err.Error())
			return encounters.Encounter{}, err
		}
		return encounter, nil
	}
}

func ReadEncounterByAccessType(accessType string, accessor string) (encounters.Encounter, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return encounters.Encounter{}, err
		}
		return ReadEncounterByID(id)
	case "name":
		return ReadEncounterByName(accessor)
	default:
		logger.Error("Invalid access type: " + accessType)
		return encounters.Encounter{}, errors.New("Invalid access type: " + accessType)
	}
}

func ReadEncounterOverviewByAccessType(accessType string, accessor string) (encounters.EncounterOverview, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return encounters.EncounterOverview{}, err
		}
		return ReadEncounterOverviewByID(id)
	case "name":
		return ReadEncounterOverviewByName(accessor)
	default:
		logger.Error("Invalid access type: " + accessType)
		return encounters.EncounterOverview{}, errors.New("Invalid access type: " + accessType)
	}
}
