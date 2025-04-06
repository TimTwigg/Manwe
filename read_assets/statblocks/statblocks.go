package read_asset_statblocks

import (
	asset_utils "github.com/TimTwigg/EncounterManagerBackend/read_assets/utils"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
	dbutils "github.com/TimTwigg/EncounterManagerBackend/utils/database"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

// Read a stat block from database
func ReadStatBlock(name string) (stat_blocks.StatBlock, error) {
	if asset_utils.StatBlockExists(name) {
		data, err := asset_utils.ReadAsset(name, "stat_blocks")
		if err != nil {
			logger.Error("Error reading stat block: " + name + ": " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		statblock, err := stat_blocks.ParseStatBlockData(data)
		if err != nil {
			return stat_blocks.StatBlock{}, err
		}
		return statblock.(stat_blocks.StatBlock), nil
	} else {
		logger.Error("StatBlock " + name + " Not Found")
	}
	return stat_blocks.StatBlock{}, nil
}

func ReadStatBlockFromDB(name string) (stat_blocks.StatBlock, error) {
	rows, err := dbutils.QuerySQL(dbutils.DB, "SELECT * FROM Entity WHERE name = ?", name)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.StatBlock{}, err
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var block stat_blocks.StatBlock
		if err := rows.Scan(
			&id,
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
			&block.Stats.ReactionCount,
			&block.Stats.Strength,
			&block.Stats.Dexterity,
			&block.Stats.Constitution,
			&block.Stats.Intelligence,
			&block.Stats.Wisdom,
			&block.Stats.Charisma,
			&block.Details.ArmorType,
		); err != nil {
			logger.Error("Error scanning row: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
		logger.Info(block)
	}

	return stat_blocks.StatBlock{}, nil
}
