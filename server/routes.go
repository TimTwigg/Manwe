package routes

import (
	"net/http"

	encounterroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/encounters"
	statblockroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/statblocks"
	requests_utils "github.com/TimTwigg/EncounterManagerBackend/utils/requests"
)

func WrapHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requests_utils.EnableCORS(&w)
		handler(w, r)
	}
}

func RegisterRoutes() {
	http.HandleFunc("/statblock", WrapHandler((statblockroutes.StatBlockHandler)))
	http.HandleFunc("/encounter", WrapHandler((encounterroutes.EncounterHandler)))
}
