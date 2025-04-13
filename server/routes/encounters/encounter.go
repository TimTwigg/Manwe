package encounterroutes

import (
	"encoding/json"
	"net/http"

	encounters "github.com/TimTwigg/EncounterManagerBackend/types/encounters"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func EncounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		logger.PostRequest("EncounterHandler: POST request")
		defer r.Body.Close()

		enc := encounters.Encounter{}
		json.NewDecoder(r.Body).Decode(&enc)
		logger.Info("EncounterHandler: Received encounter: " + enc.Name)

		w.WriteHeader(http.StatusOK)
	}
}
