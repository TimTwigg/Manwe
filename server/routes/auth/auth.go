package authroutes

import (
	"net/http"

	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		logger.OptionsRequest("AuthHandler: OPTIONS request")
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
