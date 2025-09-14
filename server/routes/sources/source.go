package sourceroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/Manwe/assets"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func AllUsedSourcesHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("SourceHandler: GET request")
		logger.GetRequest("Requesting all sources")

		sources, err := assets.ReadAllUsedSources(userid)
		if err != nil {
			http.Error(w, "Error reading used sources", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(sources); err != nil {
			logger.Error("SourceHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
