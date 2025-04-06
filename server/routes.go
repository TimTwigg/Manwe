package routes

import (
	"net/http"

	statblockroutes "github.com/TimTwigg/EncounterManagerBackend/server/routes/statblocks"
)

func RegisterRoutes() {
	http.HandleFunc("/data/statblock", statblockroutes.StatBlockHandler)
}
