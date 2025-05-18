package encounterroutes

import (
	"encoding/json"
	"net/http"
	"strconv"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets"
	server_utils "github.com/TimTwigg/EncounterManagerBackend/server/utils"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func EncounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		logger.OptionsRequest("EncounterHandler: OPTIONS request")
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		logger.GetRequest("EncounterHandler: GET request")
		defer r.Body.Close()

		detail, err := server_utils.GetDetailLevel(r)
		if err != nil || detail < 1 || detail > 2 {
			logger.Error("EncounterHandler: Error getting detail level: " + err.Error())
			http.Error(w, "Error getting detail level", http.StatusBadRequest)
			return
		}

		var accessType string = "id"
		var accessor string = ""

		accessor = r.URL.Query().Get("id")
		if accessor == "" {
			accessType = "name"
			accessor = r.URL.Query().Get("name")
			if accessor == "" {
				http.Error(w, "Encounter id or name is required", http.StatusBadRequest)
				return
			}
		}

		logger.Info("Requesting Encounter: (" + accessType + ": " + accessor + ") with detail level: " + strconv.Itoa(detail))

		switch detail {
		case 1:
			encounter, err := assets.ReadEncounterOverviewByAccessType(accessType, accessor)
			if err != nil {
				http.Error(w, "Encounter not found", server_utils.ErrorStatus(err))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(encounter); err != nil {
				logger.Error("EncounterHandler: Error encoding JSON: " + err.Error())
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		case 2:
			encounter, err := assets.ReadEncounterByAccessType(accessType, accessor)
			if err != nil {
				http.Error(w, "Encounter not found", server_utils.ErrorStatus(err))
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
		enc, err := assets.SetEncounter(enc)
		if err != nil {
			logger.Error("EncounterHandler: Error setting encounter: " + err.Error())
			http.Error(w, "Error setting encounter", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(enc); err != nil {
			logger.Error("EncounterHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
