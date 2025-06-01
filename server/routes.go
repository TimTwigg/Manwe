package routes

import (
	"net/http"

	conditionroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/conditions"
	encounterroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/encounters"
	statblockroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/statblocks"
	server_utils "github.com/TimTwigg/EncounterManagerBackend/server/utils"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func WrapHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received: [" + r.Method + "] " + r.URL.Path)
		server_utils.EnableCORS(&w)
		handler(w, r)
	}
}

func RegisterRoutes() {
	http.HandleFunc("/statblock", WrapHandler((statblockroutes.StatBlockHandler)))
	http.HandleFunc("/statblock/all", WrapHandler((statblockroutes.StatBlockOverviewHandler)))
	http.HandleFunc("/encounter", WrapHandler((encounterroutes.EncounterHandler)))
	http.HandleFunc("/encounter/", WrapHandler((encounterroutes.EncounterHandler)))
	http.HandleFunc("/encounter/all", WrapHandler((encounterroutes.EncounterOverviewHandler)))
	http.HandleFunc("/condition/all", WrapHandler((conditionroutes.AllConditionsHandler)))
}
