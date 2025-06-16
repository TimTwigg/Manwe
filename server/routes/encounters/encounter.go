package encounterroutes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets"
	server_utils "github.com/TimTwigg/EncounterManagerBackend/server/utils"
	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func EncounterHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
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

		logger.GetRequest("Requesting Encounter: (" + accessType + ": " + accessor + ") with detail level: " + strconv.Itoa(detail))

		switch detail {
		case 1:
			encounter, err := assets.ReadEncounterOverviewByAccessType(accessType, accessor, userid)
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
			encounter, err := assets.ReadEncounterByAccessType(accessType, accessor, userid)
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
		if enc.Name == "" {
			logger.Error("EncounterHandler: Encounter name is required")
			http.Error(w, "Encounter name is required", http.StatusBadRequest)
			return
		}

		logger.PostRequest("Setting Encounter: " + enc.Name)

		enc, err := assets.SetEncounter(enc, userid)
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

	case http.MethodDelete:
		logger.DeleteRequest("EncounterHandler: DELETE request")
		defer r.Body.Close()
		encounterID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/encounter/"))
		if err != nil || encounterID <= 0 {
			logger.Error("EncounterHandler: Encounter ID is required for deletion")
			http.Error(w, "Encounter ID is required for deletion", http.StatusBadRequest)
			return
		}
		logger.DeleteRequest("Deleting Encounter: " + strconv.Itoa((encounterID)))

		err = assets.DeleteEncounter(encounterID, userid)
		if err != nil {
			logger.Error("EncounterHandler: Error deleting encounter: " + err.Error())
			http.Error(w, "Error deleting encounter", server_utils.ErrorStatus(err))
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
