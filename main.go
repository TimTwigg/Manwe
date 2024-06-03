package main

import (
	damage_types "github.com/TimTwigg/EncounterManagerBackend/types/damage"
	language "github.com/TimTwigg/EncounterManagerBackend/types/languages"

	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

func init() {
	log.Init("Fully initialized!")
}

func main() {
	log.Info(language.DEFAULT_LANGUAGES.Get("Common"))
	log.Info(damage_types.DEFAULT_DAMAGE_TYPES.Get("Fire"))
}
