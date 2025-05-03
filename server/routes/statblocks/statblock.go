package statblockroutes

import (
	"encoding/json"
	"net/http"
	"strconv"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets"
	server_utils "github.com/TimTwigg/EncounterManagerBackend/server/utils"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func StatBlockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("StatBlockHandler: GET request")
		// Check if the request has a name parameter
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "StatBlock name is required", http.StatusBadRequest)
			return
		}
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

		logger.Info("Requesting StatBlock: (" + name + ") with Detail Level: (" + strconv.Itoa(detail) + ")")

		switch detail {
		case 1:
			// Read the statblock overview from the database
			statBlockOverview, err := assets.ReadStatBlockOverview(name)
			if err != nil {
				http.Error(w, "StatBlock not found", server_utils.ErrorStatus(err))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(statBlockOverview); err != nil {
				logger.Error("StatBlockHandler: Error encoding JSON: " + err.Error())
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		case 2:
			// Read the statblock from the database
			statBlock, err := assets.ReadStatBlockByName(name)
			if err != nil {
				http.Error(w, "StatBlock not found", server_utils.ErrorStatus(err))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(statBlock); err != nil {
				logger.Error("StatBlockHandler: Error encoding JSON: " + err.Error())
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Invalid detail level", http.StatusBadRequest)
		}
		return

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
