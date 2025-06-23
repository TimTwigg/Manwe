package routes

import (
	"net/http"
	"strings"

	campaignroutes "github.com/TimTwigg/Manwe/server/routes/campaigns"
	conditionroutes "github.com/TimTwigg/Manwe/server/routes/conditions"
	encounterroutes "github.com/TimTwigg/Manwe/server/routes/encounters"
	metadataroutes "github.com/TimTwigg/Manwe/server/routes/metadata"
	statblockroutes "github.com/TimTwigg/Manwe/server/routes/statblocks"
	server_utils "github.com/TimTwigg/Manwe/server/utils"
	logger "github.com/TimTwigg/Manwe/utils/log"
	supertokens "github.com/supertokens/supertokens-golang/supertokens"
)

func CORSMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			logger.OptionsRequest(r.URL.Path)
			response.Header().Set("Access-Control-Allow-Methods", "*")
			response.Header().Set("Access-Control-Allow-Headers",
				strings.Join(append([]string{"Content-Type"},
					supertokens.GetAllCORSHeaders()...), ","))
			response.WriteHeader(http.StatusOK)
		} else {
			supertokens.Middleware(next).ServeHTTP(response, r)
		}
	})
}

func HandleRoute(w http.ResponseWriter, r *http.Request) {
	userid, _ := server_utils.GetSessionUserID(r)

	if r.URL.Path == "/metadata" {
		metadataroutes.MetadataHandler(w, r, userid)
		// } else if r.URL.Path == "/support" {
		// 	supportroutes.SupportHandler(w, r, userid)
	} else if r.URL.Path == "/condition/all" {
		conditionroutes.AllConditionsHandler(w, r, userid)
	} else if r.URL.Path == "/statblock" {
		statblockroutes.StatBlockHandler(w, r, userid)
	} else if r.URL.Path == "/statblock/all" {
		statblockroutes.StatBlockOverviewHandler(w, r, userid)
	} else if r.URL.Path == "/encounter" || r.URL.Path == "/encounter/" {
		encounterroutes.EncounterHandler(w, r, userid)
	} else if r.URL.Path == "/encounter/all" {
		encounterroutes.EncounterOverviewHandler(w, r, userid)
	} else if r.URL.Path == "/campaign" || r.URL.Path == "/campaign/" {
		campaignroutes.CampaignHandler(w, r, userid)
	} else if r.URL.Path == "/campaign/all" {
		campaignroutes.CampaignOverviewHandler(w, r, userid)
	} else {
		logger.Error("Unhandled route: " + r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}
