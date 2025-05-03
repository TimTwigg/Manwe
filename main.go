package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	asset_utils "github.com/TimTwigg/EncounterManagerBackend/assets/utils"
	routes "github.com/TimTwigg/EncounterManagerBackend/server"
	logger "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func cleanup() {
	asset_utils.CloseDB(asset_utils.DB)
	logger.Info("Database closed.")
	logger.Info("Server stopped.")
}

func main() {
	routes.RegisterRoutes()

	logger.Info("Loading database...")
	database, err := asset_utils.GetDB()
	if err != nil {
		logger.Error("Error loading database: " + err.Error())
		return
	}
	asset_utils.DB = database
	logger.Info("Database loaded.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	logger.Info("Server started on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		asset_utils.CloseDB(asset_utils.DB)
		log.Fatal(err)
	}
}
