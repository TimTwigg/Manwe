package asset_utils

import (
	io "github.com/TimTwigg/EncounterManagerBackend/utils/io"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

// Read an asset from the asset file, located in the given module subfolder
func ReadAsset(name string, module string) (map[string]any, error) {
	data, err := io.ReadJSON("assets\\" + module + "\\" + name + ".json")
	if err != nil {
		logger.Error("Error reading asset: " + err.Error())
		return nil, err
	}
	return data, nil
}
