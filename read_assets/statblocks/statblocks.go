package read_asset_statblocks

import (
	asset_utils "github.com/TimTwigg/EncounterManagerBackend/read_assets/utils"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"

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
