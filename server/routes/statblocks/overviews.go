package statblockroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/Manwe/assets"
	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func StatBlockOverviewHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("StatBlockOverviewHandler: GET request")
		logger.GetRequest("Requesting all statblock overviews")

		statblocks, err := assets.ReadAllStatBlockOverviews(userid, asset_utils.STATBLOCK)
		if err != nil {
			http.Error(w, "Error reading statblocks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statblocks); err != nil {
			logger.Error("StatBlockOverviewHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
