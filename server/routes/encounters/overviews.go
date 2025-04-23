package encounterroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets/encounters"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	requests_utils "github.com/TimTwigg/EncounterManagerBackend/utils/requests"
)

func EncounterOverviewHandler(w http.ResponseWriter, r *http.Request) {
	logger.GetRequest("EncounterOverviewHandler: GET request")
	detail, err := requests_utils.GetDetailLevel(r)
	if err != nil || detail < 1 || detail > 2 {
		http.Error(w, "Invalid detail level", http.StatusBadRequest)
		return
	}

	encounters, err := assets.ReadAllEncounterOverviews()
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
}
