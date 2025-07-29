package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	encounters "github.com/TimTwigg/Manwe/types/encounters"
	entities "github.com/TimTwigg/Manwe/types/entities"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	utils "github.com/TimTwigg/Manwe/utils/functions"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

func encounterBelongsToUser(encounterID int, userid string) (bool, error) {
	// Check that the encounter belongs to the user
	var username string
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT username FROM public.encounter WHERE encounterid = $1", encounterID).Scan(&username)
	if err != nil {
		logger.Error("Error checking EncounterID: " + err.Error())
		return false, errors.Wrap(err, "Error checking EncounterID")
	}
	return username == userid, nil
}

func encounterExists(encounterID int) bool {
	var count int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT COUNT(*) FROM public.encounter WHERE encounterid = $1", encounterID).Scan(&count)
	if err != nil {
		logger.Error("Error checking EncounterID: " + err.Error())
		return false
	}
	return count > 0
}

func SetEncounterEntities(creatures []entities.Entity, encounterID int) error {
	// Empty the EncounterEntities table for this encounter
	_, err := asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.encounterentities WHERE encounterid = $1", encounterID)
	if err != nil {
		logger.Error("Error deleting EncounterEntities: " + err.Error())
		return errors.Wrap(err, "Error deleting EncounterEntities")
	}

	// Insert each creature into the EncounterEntities table
	for row, creature := range creatures {
		_, err := asset_utils.DBPool.Exec(
			context.Background(),
			"INSERT INTO public.encounterentities (encounterid, rowid, statblockid, suffix, initiative, maxhitpoints, temphitpoints, currenthitpoints, armorclassbonus, concentration, notes, ishostile, encounterlocked, id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
			encounterID,
			row+1,
			creature.DBID,
			creature.Suffix,
			creature.Initiative,
			creature.MaxHitPoints,
			creature.TempHitPoints,
			creature.CurrentHitPoints,
			creature.ArmorClassBonus,
			creature.Concentration,
			creature.Notes,
			creature.IsHostile,
			creature.EncounterLocked,
			creature.ID,
		)
		if err != nil {
			logger.Error("Error inserting EncounterEntity: " + err.Error())
			return errors.Wrap(err, "Error inserting EncounterEntity")
		}

		// Insert Conditions
		for condition, rounds := range creature.Conditions {
			_, err := asset_utils.DBPool.Exec(
				context.Background(),
				"INSERT INTO public.encentconditions (encounterid, rowid, condition, duration) VALUES ($1, $2, $3, $4)",
				encounterID,
				row+1,
				condition,
				rounds,
			)
			if err != nil {
				logger.Error("Error inserting EncounterEntityCondition: " + err.Error())
				return errors.Wrap(err, "Error inserting EncounterEntityCondition")
			}
		}
	}
	return nil
}

func SetEncounter(encounter encounters.Encounter, userid string) (encounters.Encounter, error) {
	if encounter.ID == 0 || !encounterExists(encounter.ID) {
		//  If the campaign does not exist, create it
		if !campaignExists(encounter.Metadata.Campaign, userid) {
			_, err := asset_utils.DBPool.Exec(
				context.Background(),
				"INSERT INTO public.campaign (campaign, username, description, creationdate, lastmodified) VALUES ($1, $2, $3, $4, $5)",
				encounter.Metadata.Campaign,
				userid,
				"Auto-generated campaign for encounter",
				utils.FormatDate(encounter.Metadata.AccessedDate),
				utils.FormatDate(encounter.Metadata.AccessedDate),
			)
			if err != nil {
				logger.Error("Error inserting Campaign: " + err.Error())
				return encounters.Encounter{}, errors.Wrap(err, "Error inserting Campaign")
			}
		}

		//  If the encounter does not exist, create it
		err := asset_utils.DBPool.QueryRow(
			context.Background(),
			"INSERT INTO public.encounter (name, description, creationdate, accesseddate, campaign, started, round, turn, haslair, lairownerid, activeid, username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING encounterid",
			encounter.Name,
			encounter.Description,
			utils.FormatDate(encounter.Metadata.CreationDate),
			utils.FormatDate(encounter.Metadata.AccessedDate),
			encounter.Metadata.Campaign,
			encounter.Metadata.Started,
			encounter.Metadata.Round,
			encounter.Metadata.Turn,
			encounter.HasLair,
			encounter.LairOwnerID,
			encounter.ActiveID,
			userid,
		).Scan(&encounter.ID)
		if err != nil {
			logger.Error("Error inserting Encounter: " + err.Error())
			return encounters.Encounter{}, errors.Wrap(err, "Error inserting Encounter")
		}
	} else {
		//  Check if the encounter belongs to the user
		auth, err := encounterBelongsToUser(encounter.ID, userid)
		if err != nil {
			logger.Error("Error checking if encounter belongs to user: " + err.Error())
			return encounters.Encounter{}, errors.Wrap(err, "Error checking if encounter belongs to user")
		}
		if !auth {
			logger.Error("Encounter does not belong to user")
			return encounters.Encounter{}, errors.Wrap(error_utils.AuthError{Message: "Encounter does not belong to user"}, "Error checking encounter ownership")
		}

		//  Update the encounter
		_, err = asset_utils.DBPool.Exec(
			context.Background(),
			"UPDATE public.encounter SET name = $1, description = $2, creationdate = $3, accesseddate = $4, campaign = $5, started = $6, round = $7, turn = $8, haslair = $9, lairownerid = $10, activeid = $11 WHERE encounterid = $12",
			encounter.Name,
			encounter.Description,
			utils.FormatDate(encounter.Metadata.CreationDate),
			utils.FormatDate(encounter.Metadata.AccessedDate),
			encounter.Metadata.Campaign,
			encounter.Metadata.Started,
			encounter.Metadata.Round,
			encounter.Metadata.Turn,
			encounter.HasLair,
			encounter.LairOwnerID,
			encounter.ActiveID,
			encounter.ID,
		)
		if err != nil {
			logger.Error("Error updating Encounter: " + err.Error())
			return encounters.Encounter{}, errors.Wrap(err, "Error updating Encounter")
		}
	}

	err := SetEncounterEntities(encounter.Entities, encounter.ID)
	if err != nil {
		logger.Error("Error setting Encounter entity: " + err.Error())
		return encounters.Encounter{}, errors.Wrap(err, "Error setting Encounter entity")
	}
	return encounter, nil
}

func DeleteEncounter(encounterID int, userid string) error {
	auth, err := encounterBelongsToUser(encounterID, userid)
	if err != nil {
		logger.Error("Error checking if encounter belongs to user: " + err.Error())
		return errors.Wrap(err, "Error checking if encounter belongs to user")
	}
	if !auth {
		logger.Error("Encounter does not belong to user")
		return errors.Wrap(error_utils.AuthError{Message: "Encounter does not belong to user"}, "Error checking encounter ownership")
	}

	_, err = asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.encounter WHERE encounterid = $1", encounterID)
	if err != nil {
		logger.Error("Error deleting Encounter: " + err.Error())
		return errors.Wrap(err, "Error deleting Encounter")
	}

	return nil
}
