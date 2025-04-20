package encounters

import (
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
)

type Encounter struct {
	Name           string
	Description    string
	Metadata       EncounterMetadata
	Entities       []entities.Entity
	HasLair        bool
	Lair           stat_blocks.Lair
	LairEntityName string
	ActiveID       string
}

type EncounterOverview struct {
	Name        string
	Description string
	Metadata    EncounterMetadata
}
