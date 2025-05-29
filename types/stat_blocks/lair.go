package stat_blocks

import (
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
)

type Lair struct {
	Name             string
	OwningEntityDBID int
	Description      string
	Initiative       int
	Actions          generics.ItemList
	RegionalEffects  generics.ItemList
}

func (l Lair) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Name":             l.Name,
		"OwningEntityDBID": l.OwningEntityDBID,
		"Description":      l.Description,
		"Initiative":       l.Initiative,
		"Actions":          l.Actions,
		"RegionalEffects":  l.RegionalEffects,
	}
}
