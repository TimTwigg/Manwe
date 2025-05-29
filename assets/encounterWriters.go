package assets

import (
	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func SetEncounterEntities(creatures []entities.Entity, encounterID int) error {
	// Empty the EncounterEntities table for this encounter
	_, err := asset_utils.ExecSQL(asset_utils.DB, "DELETE FROM EncounterEntities WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error deleting EncounterEntities: " + err.Error())
		return err
	}
	_, err = asset_utils.ExecSQL(asset_utils.DB, "DELETE FROM EncEntConditions WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error deleting EncEntConditions: " + err.Error())
		return err
	}

	// Insert each creature into the EncounterEntities table
	for row, creature := range creatures {
		_, err := asset_utils.ExecSQL(
			asset_utils.DB,
			"INSERT INTO EncounterEntities (EncounterID, RowID, EntityID, Suffix, Initiative, MaxHitPoints, TempHitPoints, CurrentHitPoints, ArmorClassBonus, Concentration, Notes, IsHostile, EncounterLocked, Domain, Published, ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			encounterID,
			row+1,
			creature.DBID,
			creature.Suffix,
			creature.Initiative,
			creature.MaxHitPoints,
			creature.TempHitPoints,
			creature.CurrentHitPoints,
			creature.ArmorClassBonus,
			utils.FormatBool(creature.Concentration),
			creature.Notes,
			utils.FormatBool(creature.IsHostile),
			utils.FormatBool(creature.EncounterLocked),
			"Private", // TODO
			"",        // TODO
			creature.ID,
		)
		if err != nil {
			logger.Error("Error inserting EncounterEntity: " + err.Error())
			return err
		}

		// Insert Conditions
		for condition, rounds := range creature.Conditions {
			_, err := asset_utils.ExecSQL(
				asset_utils.DB,
				"INSERT INTO EncEntConditions (EncounterID, RowID, Condition, Duration) VALUES (?, ?, ?, ?)",
				encounterID,
				row+1,
				condition,
				rounds,
			)
			if err != nil {
				logger.Error("Error inserting EncounterEntityCondition: " + err.Error())
				return err
			}
		}
	}
	return nil
}

func SetEncounter(encounter encounters.Encounter) (encounters.Encounter, error) {
	if encounter.ID == 0 {
		res, err := asset_utils.ExecSQL(
			asset_utils.DB,
			"INSERT INTO Encounter (Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn, HasLair, LairOwnerID, ActiveID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			encounter.Name,
			encounter.Description,
			utils.FormatDate(encounter.Metadata.CreationDate),
			utils.FormatDate(encounter.Metadata.AccessedDate),
			encounter.Metadata.Campaign,
			utils.FormatBool(encounter.Metadata.Started),
			encounter.Metadata.Round,
			encounter.Metadata.Turn,
			utils.FormatBool(encounter.HasLair),
			encounter.LairOwnerID,
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
		err = SetEncounterEntities(encounter.Entities, encounter.ID)
		if err != nil {
			logger.Error("Error setting Encounter entity: " + err.Error())
			return encounters.Encounter{}, err
		}
		return encounter, nil
	} else {
		_, err := asset_utils.ExecSQL(
			asset_utils.DB,
			"UPDATE Encounter SET Name = ?, Description = ?, CreationDate = ?, AccessedDate = ?, Campaign = ?, Started = ?, Round = ?, Turn = ?, HasLair = ?, LairOwnerID = ?, ActiveID = ? WHERE EncounterID = ?",
			encounter.Name,
			encounter.Description,
			utils.FormatDate(encounter.Metadata.CreationDate),
			utils.FormatDate(encounter.Metadata.AccessedDate),
			encounter.Metadata.Campaign,
			utils.FormatBool(encounter.Metadata.Started),
			encounter.Metadata.Round,
			encounter.Metadata.Turn,
			utils.FormatBool(encounter.HasLair),
			encounter.LairOwnerID,
			encounter.ActiveID,
			encounter.ID,
		)
		if err != nil {
			logger.Error("Error updating Encounter: " + err.Error())
			return encounters.Encounter{}, err
		}
		err = SetEncounterEntities(encounter.Entities, encounter.ID)
		if err != nil {
			logger.Error("Error setting Encounter entity: " + err.Error())
			return encounters.Encounter{}, err
		}

		return encounter, nil
	}
}
