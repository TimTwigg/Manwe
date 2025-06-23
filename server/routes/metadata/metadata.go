package metadataroutes

import (
	"encoding/json"
	"net/http"

	logger "github.com/TimTwigg/Manwe/utils/log"
	usermetadata "github.com/supertokens/supertokens-golang/recipe/usermetadata"
)

func MetadataHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodGet:
		logger.GetRequest("MetadataHandler: GET request")

		metadata, err := usermetadata.GetUserMetadata(userid)
		if err != nil {
			logger.Error("MetadataHandler: Error getting user metadata: " + err.Error())
			http.Error(w, "Error getting user metadata", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(metadata); err != nil {
			logger.Error("MetadataHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		logger.PostRequest("MetadataHandler: POST request")
		var metadata map[string]string
		if err := json.NewDecoder(r.Body).Decode(&metadata); err != nil {
			logger.Error("MetadataHandler: Error decoding JSON: " + err.Error())
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Convert map[string]string to map[string]any
		metadataInterface := make(map[string]any, len(metadata))
		for key, value := range metadata {
			if value == "" {
				// Set empty values to nil to delete record from metadata
				metadataInterface[key] = nil
			} else {
				metadataInterface[key] = value
			}
		}
		// Update user metadata to Supertokens
		meta, err := usermetadata.UpdateUserMetadata(userid, metadataInterface)
		if err != nil {
			logger.Error("MetadataHandler: Error updating user metadata: " + err.Error())
			http.Error(w, "Error updating user metadata", http.StatusInternalServerError)
			return
		}

		logger.PostRequest("User metadata updated successfully")

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			logger.Error("MetadataHandler: Error encoding JSON: " + err.Error())
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
