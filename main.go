package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

func cleanup() {
	dbutils.CloseDB(dbutils.DB)
	logger.Info("Database closed.")
	logger.Info("Server stopped.")
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	logger.Info("Server started on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		dbutils.CloseDB(dbutils.DB)
		log.Fatal(err)
	}
}
