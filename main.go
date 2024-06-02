package main

import (
	"github.com/TimTwigg/EncounterManagerBackend/types/damage"
	language "github.com/TimTwigg/EncounterManagerBackend/types/languages"
)

func initialize() {
	language.InitializeDefaultLanguages()
	damage.InitializeDefaultDamageTypes()
}

func main() {
	initialize()
}
