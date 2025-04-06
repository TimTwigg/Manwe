package statblockroutes

import (
	"encoding/json"
	"net/http"

	read_asset_statblocks "github.com/TimTwigg/EncounterManagerBackend/read_assets/statblocks"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func StatBlockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Check if the request has a name parameter
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "StatBlock name is required", http.StatusBadRequest)
			return
		}
		logger.GetRequest("StatBlock: " + name)

		// Read the statblock from the asset file
		statBlock, err := read_asset_statblocks.ReadStatBlock(name)
		if err != nil {
			http.Error(w, "StatBlock not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(statBlock); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		return
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
