package main

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"

	// "strcconv"
	// "sync"
	"log"

	// condition "github.com/TimTwigg/EncounterManagerBackend/types/conditions"
	// damage_types "github.com/TimTwigg/EncounterManagerBackend/types/damage"
	// language "github.com/TimTwigg/EncounterManagerBackend/types/languages"
	// stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
	routes "github.com/TimTwigg/EncounterManagerBackend/server"
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
	logger.Info("Server started on port 8080")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
