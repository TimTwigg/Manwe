package routes

import (
	"net/http"
	"strings"

	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	conditionroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/conditions"
	encounterroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/encounters"
	statblockroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/statblocks"
	server_utils "github.com/TimTwigg/EncounterManagerBackend/server/utils"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	supertokens "github.com/supertokens/supertokens-golang/supertokens"
)

func CORSMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		logger.Info("Request received: [" + r.Method + "] " + r.URL.Path)
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
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
	logger.Info("User ID: " + userid)
	err := asset_utils.UpsertUser(asset_utils.DB, userid)
	if err != nil {
		logger.Error("Error upserting user: " + err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.URL.Path == "/statblock" {
		statblockroutes.StatBlockHandler(w, r, userid)
	} else if r.URL.Path == "/statblock/all" {
		statblockroutes.StatBlockOverviewHandler(w, r, userid)
	} else if r.URL.Path == "/encounter" || r.URL.Path == "/encounter/" {
		encounterroutes.EncounterHandler(w, r, userid)
	} else if r.URL.Path == "/encounter/all" {
		encounterroutes.EncounterOverviewHandler(w, r, userid)
	} else if r.URL.Path == "/condition/all" {
		conditionroutes.AllConditionsHandler(w, r, userid)
	} else {
		logger.Error("Unhandled route: " + r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}
