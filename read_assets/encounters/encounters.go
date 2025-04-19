package read_asset_encounters

import (
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
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
	// encounter.InitiativeOrder = make([]chan struct{}, 0)

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

	return encounter, nil
}
