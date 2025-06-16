package encounterroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func EncounterOverviewHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("EncounterOverviewHandler: GET request")
		logger.GetRequest("Requesting all encounter overviews")

		encounters, err := assets.ReadAllEncounterOverviews(userid)
		if err != nil {
			http.Error(w, "Error reading encounters", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(encounters); err != nil {
			logger.Error("EncounterOverviewHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
