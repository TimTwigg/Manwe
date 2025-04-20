package encounterroutes

import (
	"encoding/json"
	"net/http"
	"strconv"

	read_asset_encounters "github.com/TimTwigg/EncounterManagerBackend/read_assets/encounters"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func EncounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("EncounterHandler: GET request")
		defer r.Body.Close()

		detail_level := r.URL.Query().Get("detail_level")
		var detail int = 1
		if detail_level != "" {
			d, err := strconv.Atoi(detail_level)
			if err != nil || (d < 1 || d > 2) {
				http.Error(w, "Invalid detail level", http.StatusBadRequest)
				return
			}
			detail = d
		}

		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Encounter name is required", http.StatusBadRequest)
			return
		}
		logger.Info("Requesting Encounter: (" + name + ") with detail level: " + strconv.Itoa(detail))

		switch detail {
		case 1:
			encounter, err := read_asset_encounters.ReadEncounterOverview(name)
			if err != nil {
				http.Error(w, "Encounter not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(encounter); err != nil {
				logger.Error("EncounterHandler: Error encoding JSON: " + err.Error())
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		case 2:
			encounter, err := read_asset_encounters.ReadEncounterFromDB(name)
			if err != nil {
				http.Error(w, "Encounter not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(encounter); err != nil {
				logger.Error("EncounterHandler: Error encoding JSON: " + err.Error())
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Invalid detail level", http.StatusBadRequest)
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
