package assets

import (
	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	encounters "github.com/TimTwigg/Manwe/types/encounters"
	entities "github.com/TimTwigg/Manwe/types/entities"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	utils "github.com/TimTwigg/Manwe/utils/functions"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func encounterBelongsToUser(encounterID int, userid string) (bool, error) {
	// Check that the encounter belongs to the user
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Domain FROM Encounter WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error checking EncounterID: " + err.Error())
		return false, err
	}
	var domain string
	if rows.Next() {
		if err := rows.Scan(&domain); err != nil {
			logger.Error("Error scanning domain from Encounter: " + err.Error())
			return false, err
		}
	}
	if domain != userid {
		logger.Error("EncounterID does not belong to user")
		return false, error_utils.AuthError{Message: "EncounterID does not belong to user"}
	}
	return true, nil
}

func encounterExists(encounterID int) bool {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT COUNT(*) FROM Encounter WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error checking EncounterID: " + err.Error())
		return false
	}
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			logger.Error("Error scanning count from Encounter: " + err.Error())
			return false
		}
	}
	return count > 0
}

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
			"INSERT INTO EncounterEntities (EncounterID, RowID, StatBlockID, Suffix, Initiative, MaxHitPoints, TempHitPoints, CurrentHitPoints, ArmorClassBonus, Concentration, Notes, IsHostile, EncounterLocked, ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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

func SetEncounter(encounter encounters.Encounter, userid string) (encounters.Encounter, error) {
	if encounter.ID == 0 || !encounterExists(encounter.ID) {
		//  If the campaign does not exist, create it
		if !campaignExists(encounter.Metadata.Campaign, userid) {
			_, err := asset_utils.ExecSQL(
				asset_utils.DB,
				"INSERT INTO Campaign (Campaign, Domain, Description, CreationDate, LastModified) VALUES (?, ?, ?, ?, ?)",
				encounter.Metadata.Campaign,
				userid,
				"Auto-generated campaign for encounter",
				utils.FormatDate(encounter.Metadata.AccessedDate),
				utils.FormatDate(encounter.Metadata.AccessedDate),
			)
			if err != nil {
				logger.Error("Error inserting Campaign: " + err.Error())
				return encounters.Encounter{}, err
			}
		}

		//  If the encounter does not exist, create it
		res, err := asset_utils.ExecSQL(
			asset_utils.DB,
			"INSERT INTO Encounter (Name, Description, CreationDate, AccessedDate, Campaign, Started, Round, Turn, HasLair, LairOwnerID, ActiveID, Domain) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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
			userid,
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
	} else {
		//  Check if the encounter belongs to the user
		auth, err := encounterBelongsToUser(encounter.ID, userid)
		if err != nil {
			logger.Error("Error checking if encounter belongs to user: " + err.Error())
			return encounters.Encounter{}, err
		}
		if !auth {
			logger.Error("Encounter does not belong to user")
			return encounters.Encounter{}, error_utils.AuthError{Message: "Encounter does not belong to user"}
		}

		//  Update the encounter
		_, err = asset_utils.ExecSQL(
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
	}

	err := SetEncounterEntities(encounter.Entities, encounter.ID)
	if err != nil {
		logger.Error("Error setting Encounter entity: " + err.Error())
		return encounters.Encounter{}, err
	}
	return encounter, nil
}

func DeleteEncounter(encounterID int, userid string) error {
	auth, err := encounterBelongsToUser(encounterID, userid)
	if err != nil {
		logger.Error("Error checking if encounter belongs to user: " + err.Error())
		return err
	}
	if !auth {
		logger.Error("Encounter does not belong to user")
		return error_utils.AuthError{Message: "Encounter does not belong to user"}
	}

	_, err = asset_utils.ExecSQL(asset_utils.DB, "DELETE FROM EncEntConditions WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error deleting EncEntConditions: " + err.Error())
		return err
	}

	_, err = asset_utils.ExecSQL(asset_utils.DB, "DELETE FROM EncounterEntities WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error deleting EncounterEntities: " + err.Error())
		return err
	}

	_, err = asset_utils.ExecSQL(asset_utils.DB, "DELETE FROM Encounter WHERE EncounterID = ?", encounterID)
	if err != nil {
		logger.Error("Error deleting Encounter: " + err.Error())
		return err
	}

	return nil
}
