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
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			logger.OptionsRequest(r.URL.Path)
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
	if userid == "" {
		userid = "public"
	}
	// } else if r.Method != http.MethodOptions {
	// 	_ = asset_utils.UpsertUser(asset_utils.DB, userid)
	// }

	route := strings.TrimPrefix(r.URL.Path, "/")
	if strings.Contains(route, "/") && !strings.Contains(route, "all") {
		route = strings.Split(route, "/")[0]
	}

	switch route {
	case "metadata":
		metadataroutes.MetadataHandler(w, r, userid)
	case "condition/all":
		conditionroutes.AllConditionsHandler(w, r, userid)
	case "statblock":
		statblockroutes.StatBlockHandler(w, r, userid)
	case "statblock/all":
		statblockroutes.StatBlockOverviewHandler(w, r, userid)
	case "encounter":
		encounterroutes.EncounterHandler(w, r, userid)
	case "encounter/all":
		encounterroutes.EncounterOverviewHandler(w, r, userid)
	case "campaign":
		campaignroutes.CampaignHandler(w, r, userid)
	case "campaign/all":
		campaignroutes.CampaignOverviewHandler(w, r, userid)
	default:
		logger.Error("Unhandled route: " + r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}
