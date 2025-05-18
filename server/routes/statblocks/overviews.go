package statblockroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func StatBlockOverviewHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		logger.OptionsRequest("StatBlockOverviewHandler: OPTIONS request")
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		logger.GetRequest("StatBlockOverviewHandler: GET request")

		statblocks, err := assets.ReadAllStatBlockOverviews()
		if err != nil {
			http.Error(w, "Error reading statblocks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statblocks); err != nil {
			logger.Error("StatBlockOverviewHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
