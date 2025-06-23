package campaignroutes

import (
	"encoding/json"
	"net/http"

	assets "github.com/TimTwigg/Manwe/assets"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func CampaignOverviewHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("CampaignOverviewHandler: GET request")
		logger.GetRequest("Requesting all campaign overviews")

		campaigns, err := assets.ReadAllCampaignOverviews(userid)
		if err != nil {
			http.Error(w, "Error reading campaigns", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(campaigns); err != nil {
			logger.Error("CampaignOverviewHandler: Error encoding JSON: " + err.Error())

			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
