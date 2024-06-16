package main

import (
	condition "github.com/TimTwigg/EncounterManagerBackend/types/conditions"
	damage_types "github.com/TimTwigg/EncounterManagerBackend/types/damage"
	language "github.com/TimTwigg/EncounterManagerBackend/types/languages"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
	validate "github.com/TimTwigg/EncounterManagerBackend/utils/validation"
)

func init() {
	log.Init("Fully initialized!")
}

func Validate() {
	assetPath := "./assets"
	hideOutput := true
	validate.ValidateStatBlocks(assetPath, hideOutput)
	validate.ValidateLanguage(assetPath, hideOutput)
	validate.ValidateDamageTypes(assetPath, hideOutput)
	validate.ValidateConditions(assetPath, hideOutput)
	log.Info("Validation complete!")
}

func main() {
	Validate()
	log.Info(language.DEFAULT_LANGUAGES.Get("Common"))
	log.Info(damage_types.DEFAULT_DAMAGE_TYPES.Get("Fire"))
	log.Info(condition.DEFAULT_CONDITIONS.Get("Blinded"))
	log.Info(stat_blocks.DEFAULT_STAT_BLOCKS.Get("Aurelia"))
}
