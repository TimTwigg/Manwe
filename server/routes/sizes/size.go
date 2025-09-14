package sizeroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/Manwe/assets"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func AllSizesHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("SizeHandler: GET request")
		logger.GetRequest("Requesting all sizes")

		sizes, err := assets.ReadAllSizes(userid)
		if err != nil {
			http.Error(w, "Error reading entity sizes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(sizes); err != nil {
			logger.Error("SizeHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
