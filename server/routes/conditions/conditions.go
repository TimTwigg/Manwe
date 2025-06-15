package conditionroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/EncounterManagerBackend/assets"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func AllConditionsHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodOptions:
		logger.OptionsRequest("ConditionHandler: OPTIONS request")
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		logger.GetRequest("ConditionHandler: GET request")
		logger.GetRequest("Requesting all conditions")

		conditions, err := assets.ReadAllConditions(userid)
		if err != nil {
			http.Error(w, "Error reading conditions", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(conditions); err != nil {
			logger.Error("ConditionHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
