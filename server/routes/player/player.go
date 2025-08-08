package playerroutes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	assets "github.com/TimTwigg/Manwe/assets"
	server_utils "github.com/TimTwigg/Manwe/server/utils"
	player "github.com/TimTwigg/Manwe/types/player"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func PlayerHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodPost:
		logger.PostRequest("PlayerHandler: POST request")
		defer r.Body.Close()

		plyr := player.Player{}
		if err := json.NewDecoder(r.Body).Decode(&plyr); err != nil {
			logger.Error("PlayerHandler: Failed to decode JSON: " + err.Error())
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		plyr, err := assets.SetPlayer(plyr, userid)
		if err != nil {
			logger.Error("PlayerHandler: Error setting player: " + err.Error())
			http.Error(w, "Error setting player", server_utils.ErrorStatus(err))
			return
		}

		logger.PostRequest("PlayerHandler: Player set successfully: " + plyr.StatBlock.Name)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(plyr); err != nil {
			logger.Error("PlayerHandler: Failed to encode JSON: " + err.Error())
			http.Error(w, "Error encoding player", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		logger.DeleteRequest("PlayerHandler: DELETE request")
		defer r.Body.Close()

		comboID := strings.TrimPrefix(r.URL.Path, "/player/")
		if comboID == "" {
			logger.Error("PlayerHandler: Failed to parse comboID")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		pieces := strings.Split(comboID, ",")
		if len(pieces) != 2 {
			logger.Error("PlayerHandler: Invalid comboID format")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		campaign := pieces[0]
		rowID, err := strconv.Atoi(pieces[1])
		if err != nil {
			logger.Error("PlayerHandler: Failed to parse playerID: " + err.Error())
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		logger.DeleteRequest("Deleting Player with rowID: " + strconv.Itoa(rowID) + " from campaign: " + campaign)

		err = assets.DeletePlayer(campaign, rowID, userid)
		if err != nil {
			logger.Error("PlayerHandler: Error deleting player: " + err.Error())
			http.Error(w, "Error deleting player", server_utils.ErrorStatus(err))
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
