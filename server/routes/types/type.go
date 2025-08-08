package typeroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/Manwe/assets"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func AllTypesHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("TypeHandler: GET request")
		logger.GetRequest("Requesting all types")

		types, err := assets.ReadAllTypes(userid)
		if err != nil {
			http.Error(w, "Error reading entity types", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(types); err != nil {
			logger.Error("TypeHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
