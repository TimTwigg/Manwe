package stat_blocks

import (
	actions "github.com/TimTwigg/EncounterManagerBackend/types/actions"
	condition "github.com/TimTwigg/EncounterManagerBackend/types/conditions"
)

type StatBlock struct {
	Name                string
	ChallengeRating     int
	ProficiencyBonus    int
	Description         EntityDescription
	Stats               NumericalAttributes
	DamageModifiers     DamageModifiers
	ConditionImmunities []condition.Condition
	Details             DetailBlock
	Actions             []actions.Action
	BonusActions        []actions.SimpleAction
	Reactions           []actions.SimpleAction
	LegendaryActions    actions.Legendary
}
