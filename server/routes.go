package routes

import (
	"net/http"

	encounterroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/encounters"
	statblockroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/statblocks"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	requests_utils "github.com/TimTwigg/EncounterManagerBackend/utils/requests"
)

func WrapHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received: [" + r.Method + "] " + r.URL.Path)
		requests_utils.EnableCORS(&w)
		handler(w, r)
	}
}

func RegisterRoutes() {
	http.HandleFunc("/statblock", WrapHandler((statblockroutes.StatBlockHandler)))
	// http.HandleFunc("/statblock/overview", WrapHandler((statblockroutes.StatBlockOverviewHandler)))
	http.HandleFunc("/encounter", WrapHandler((encounterroutes.EncounterHandler)))
	http.HandleFunc("/encounter/all", WrapHandler((encounterroutes.EncounterOverviewHandler)))
}
