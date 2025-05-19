package assets

import (
	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func SetEncounterEntity(creature entities.Entity, encounterID int) error {
	// rows, err := asset_utils.QuerySQL(asset_utils.DB, "") // Check if the row exists
	return nil
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
		for _, entity := range encounter.Entities {
			err := SetEncounterEntity(entity, encounter.ID)
			if err != nil {
				logger.Error("Error setting Encounter entity: " + err.Error())
				return encounters.Encounter{}, err
			}
		}
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
		for _, entity := range encounter.Entities {
			err := SetEncounterEntity(entity, encounter.ID)
			if err != nil {
				logger.Error("Error setting Encounter entity: " + err.Error())
				return encounters.Encounter{}, err
			}
		}
		return encounter, nil
	}
}
