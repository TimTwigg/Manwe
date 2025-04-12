package main

import (
	"log"
	"net/http"

	routes "github.com/TimTwigg/EncounterManagerBackend/server"
	dbutils "github.com/TimTwigg/EncounterManagerBackend/utils/database"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	validate "github.com/TimTwigg/EncounterManagerBackend/utils/validation"
)

func init() {
	logger.Info("Fully initialized!")
}

func Validate() {
	assetPath := "./assets"
	hideOutput := true
	validate.ValidateStatBlocks(assetPath, hideOutput)
	validate.ValidateLanguage(assetPath, hideOutput)
	validate.ValidateDamageTypes(assetPath, hideOutput)
	validate.ValidateConditions(assetPath, hideOutput)
	logger.Info("Validation complete!")
}

func main() {
	routes.RegisterRoutes()

	logger.Info("Loading database...")
	database, err := dbutils.GetDB()
	if err != nil {
		logger.Error("Error loading database: " + err.Error())
		return
	}
	dbutils.DB = database
	logger.Info("Database loaded.")

	logger.Info("Server started on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
