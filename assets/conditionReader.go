package assets

import (
	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	condition "github.com/TimTwigg/EncounterManagerBackend/types/conditions"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func ReadAllConditions() ([]condition.Condition, error) {
	rows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Condition FROM Condition")
	if err != nil {
		logger.Error("Error reading conditions from database" + err.Error())
		return nil, err
	}

	conditions := make([]condition.Condition, 0, 0)
	for rows.Next() {
		var cond condition.Condition
		if err := rows.Scan(&cond.Name); err != nil {
			logger.Error("Error scanning condition row: " + err.Error())
			return nil, err
		}
		conditions = append(conditions, cond)

		// Read effects
		effectRows, err := asset_utils.QuerySQL(asset_utils.DB, "SELECT Description FROM ConditionEffect WHERE Condition = ?", cond.Name)
		if err != nil {
			logger.Error("Error reading effects for condition " + cond.Name + ": " + err.Error())
			return nil, err
		}
		for effectRows.Next() {
			var effect string
			if err := effectRows.Scan(&effect); err != nil {
				logger.Error("Error scanning effect row for condition " + cond.Name + ": " + err.Error())
				return nil, err
			}
			cond.Effects = append(cond.Effects, effect)
		}

	}

	// Return the conditions
	return conditions, nil
}
