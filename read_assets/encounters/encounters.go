package read_asset_encounters

import (
	read_asset_statblocks "github.com/TimTwigg/EncounterManagerBackend/read_assets/statblocks"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	dbutils "github.com/TimTwigg/EncounterManagerBackend/utils/database"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	errors "github.com/pkg/errors"
)

func ReadEncounterFromDB(name string) (encounters.Encounter, error) {
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM Encounter WHERE name = ?", name)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	defer rows.Close()

	var id int
	var encounter encounters.Encounter

	encounter.Entities = make([]entities.Entity, 0)
	encounter.Lair.Actions.Items = make([]generics.SimpleItem, 0)
	encounter.Lair.RegionalEffects.Items = make([]generics.SimpleItem, 0)

	if rows.Next() {
		var CreationDate, AccessedDate, Started, HasLair string
		if err := rows.Scan(
			&id,
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
		logger.Error("No Encounter found with name: " + name)
		return encounters.Encounter{}, errors.New("No Encounter found with name: " + name)
	}

	entity_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT RowID, EntityID, Suffix, Initiative, MaxHitPoints, TempHitPoints, CurrentHitPoints, ArmorClassBonus, Notes, IsHostile, EncounterLocked FROM EncounterEntities WHERE EncounterID = ?", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	defer entity_rows.Close()
	for entity_rows.Next() {
		var rowID, entityID, initiative, maxHitPoints, tempHitPoints, currentHitPoints, armorClassBonus int
		var suffix, notes, isHostile, encounterLocked string

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
		); err != nil {
			logger.Error("Error Scanning Encounter Entity Row: " + err.Error())
			return encounters.Encounter{}, err
		}

		statblock, err := read_asset_statblocks.ReadStatBlockByID(entityID)
		if err != nil {
			logger.Error("Error reading statblock: " + err.Error())
			return encounters.Encounter{}, err
		}

		entity := entities.Entity{
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

		conditions_rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT Condition, Duration FROM EncEntConditions WHERE EncounterID = ? and RowID = ?", id, rowID)
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

func ReadEncounterOverview(name string) (encounters.EncounterOverview, error) {
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM Encounter WHERE name = ?", name)
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
	} else {
		logger.Error("No Encounter found with name: " + name)
		return encounters.EncounterOverview{}, errors.New("No Encounter found with name: " + name)
	}

	return encounter, nil
}

func ReadAllEncounterOverviews() ([]encounters.EncounterOverview, error) {
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM Encounter")
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var encounters []encounters.EncounterOverview = make([]encounters.EncounterOverview, 0)

	// for rows.Next() {
	// 	var Name, Description, Campaign, CreationDate, AccessedDate string
	// 	var Round, Turn int
	// 	if err := rows.Scan(
	// 		&Name,
	// 		&Description,
	// 		&CreationDate,
	// 		&AccessedDate,
	// 		&Campaign,
	// 		&Round,
	// 		&Turn,
	// 	); err != nil {
	// 		logger.Error("Error Scanning Encounter Row: " + err.Error())
	// 		return nil, err
	// 	}
	// 	encounter = encounters.EncounterOverview{
	// 		Name:        Name,
	// 		Description: Description,
	// 		Metadata: encounters.EncounterMetadata{
	// 			Campaign: Campaign,
	// 			Round:    Round,
	// 			Turn:     Turn,
	// 		},
	// 	}
	// 	encounter.Metadata.CreationDate = utils.ParseStringDate(CreationDate)
	// 	encounter.Metadata.AccessedDate = utils.ParseStringDate(AccessedDate)
	// 	encounters = append(encounters, encounter)
	// }

	return encounters, nil
}
