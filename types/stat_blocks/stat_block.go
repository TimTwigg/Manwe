package stat_blocks

import (
	actions "github.com/TimTwigg/EncounterManagerBackend/types/actions"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
)

type StatBlock struct {
	ID                  int
	Name                string
	ChallengeRating     float32
	ProficiencyBonus    int
	Source              string
	Description         EntityDescription
	Stats               NumericalAttributes
	DamageModifiers     DamageModifiers
	ConditionImmunities []string
	Details             DetailBlock
	Actions             []actions.Action
	BonusActions        []generics.SimpleItem
	Reactions           []generics.SimpleItem
	LegendaryActions    actions.Legendary
	MythicActions       actions.Mythic
	Lair                Lair
}

func (sb StatBlock) Dict() map[string]any {
	return map[string]any{
		"ID":                  sb.ID,
		"Name":                sb.Name,
		"ChallengeRating":     sb.ChallengeRating,
		"ProficiencyBonus":    sb.ProficiencyBonus,
		"Source":              sb.Source,
		"Description":         sb.Description,
		"Stats":               sb.Stats,
		"DamageModifiers":     sb.DamageModifiers,
		"ConditionImmunities": sb.ConditionImmunities,
		"Details":             sb.Details,
		"Actions":             sb.Actions,
		"BonusActions":        sb.BonusActions,
		"Reactions":           sb.Reactions,
		"LegendaryActions":    sb.LegendaryActions,
		"MythicActions":       sb.MythicActions,
		"Lair":                sb.Lair,
	}
}
