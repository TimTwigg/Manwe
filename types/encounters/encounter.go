package encounters

import (
	entities "github.com/TimTwigg/EncounterManagerBackend/types/entities"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
)

type Encounter struct {
	ID             int
	Name           string
	Description    string
	Metadata       EncounterMetadata
	Entities       []entities.Entity
	HasLair        bool
	Lair           stat_blocks.Lair
	LairEntityName string
	ActiveID       string
}

func (e Encounter) Dict() map[string]interface{} {
	return map[string]interface{}{
		"data_type":   "Encounter",
		"ID":          e.ID,
		"Name":        e.Name,
		"Description": e.Description,
		"Metadata":    e.Metadata.Dict(),
		"Entities":    e.Entities,
		"HasLair":     e.HasLair,
	}
}

type EncounterOverview struct {
	ID          int
	Name        string
	Description string
	Metadata    EncounterMetadata
}

func (eo EncounterOverview) Dict() map[string]interface{} {
	return map[string]interface{}{
		"data_type":   "EncounterOverview",
		"ID":          eo.ID,
		"Name":        eo.Name,
		"Description": eo.Description,
		"Metadata":    eo.Metadata.Dict(),
	}
}
