package routes

import (
	"net/http"
	"strings"

	conditionroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/conditions"
	encounterroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/encounters"
	statblockroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/statblocks"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	supertokens "github.com/supertokens/supertokens-golang/supertokens"
)

func Middleware(next http.Handler) http.Handler {
	return supertokens.Middleware(http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		logger.Info("Request received: [" + r.Method + "] " + r.URL.Path)
		response.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		response.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			response.Header().Set("Access-Control-Allow-Headers",
				strings.Join(append([]string{"Content-Type"},
					supertokens.GetAllCORSHeaders()...), ","))
			response.Header().Set("Access-Control-Allow-Methods", "*")
			response.Write([]byte(""))
		} else {
			next.ServeHTTP(response, r)
		}
	}))
}

func HandleRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/statblock" {
		statblockroutes.StatBlockHandler(w, r)
	} else if r.URL.Path == "/statblock/all" {
		statblockroutes.StatBlockOverviewHandler(w, r)
	} else if r.URL.Path == "/encounter" || r.URL.Path == "/encounter/" {
		encounterroutes.EncounterHandler(w, r)
	} else if r.URL.Path == "/encounter/all" {
		encounterroutes.EncounterOverviewHandler(w, r)
	} else if r.URL.Path == "/condition/all" {
		conditionroutes.AllConditionsHandler(w, r)
	} else {
		logger.Error("Unhandled route: " + r.URL.Path)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
}
