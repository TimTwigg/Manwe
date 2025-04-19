package encounterroutes

import (
	"encoding/json"
	"net/http"

	read_asset_encounters "github.com/TimTwigg/EncounterManagerBackend/read_assets/encounters"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func EncounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("EncounterHandler: GET request")
		defer r.Body.Close()

		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Encounter name is required", http.StatusBadRequest)
			return
		}
		logger.Info("Requesting Encounter: (" + name + ")")

		encounter, err := read_asset_encounters.ReadEncounterFromDB(name)
		if err != nil {
			http.Error(w, "Encounter not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(encounter); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		logger.PostRequest("EncounterHandler: POST request")
		defer r.Body.Close()

		enc := encounters.Encounter{}
		json.NewDecoder(r.Body).Decode(&enc)
		logger.Info("EncounterHandler: Received encounter: " + enc.Name)

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
