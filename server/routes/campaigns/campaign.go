package campaignroutes

import (
	"encoding/json"
	"net/http"
	"strings"

	assets "github.com/TimTwigg/Manwe/assets"
	server_utils "github.com/TimTwigg/Manwe/server/utils"
	campaign "github.com/TimTwigg/Manwe/types/campaign"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func CampaignHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("CampaignHandler: GET request")

		var name string = r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "Campaign name is required", http.StatusBadRequest)
			return
		}

		logger.GetRequest("Requesting Campaign: (" + name + ", " + userid[:10] + "...)")

		campaign, err := assets.ReadCampaign(name, userid)
		if err != nil {
			http.Error(w, "Campaign not found", server_utils.ErrorStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(campaign); err != nil {
			logger.Error("CampaignHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		logger.PostRequest("CampaignHandler: POST request")
		defer r.Body.Close()

		camp := campaign.Campaign{}
		json.NewDecoder(r.Body).Decode(&camp)
		if camp.Name == "" {
			logger.Error("CampaignHandler: Campaign name is required")
			http.Error(w, "Campaign name is required", http.StatusBadRequest)
			return
		}

		camp, err := assets.SetCampaign(camp, userid)
		if err != nil {
			logger.Error("CampaignHandler: Error setting campaign: " + err.Error())
			http.Error(w, "Error setting campaign", server_utils.ErrorStatus(err))
			return
		}

		logger.PostRequest("CampaignHandler: Campaign set successfully: " + camp.Name)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(camp); err != nil {
			logger.Error("CampaignHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		logger.DeleteRequest("CampaignHandler: DELETE request")
		defer r.Body.Close()

		name := strings.TrimPrefix(r.URL.Path, "/campaign/")
		if name == "" {
			logger.Error("CampaignHandler: Campaign name is required for deletion")
			http.Error(w, "Campaign name is required", http.StatusBadRequest)
			return
		}

		logger.DeleteRequest("Deleting Campaign: " + name)

		err := assets.DeleteCampaign(name, userid)
		if err != nil {
			logger.Error("CampaignHandler: Error deleting campaign: " + err.Error())
			http.Error(w, "Error deleting campaign", server_utils.ErrorStatus(err))
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
