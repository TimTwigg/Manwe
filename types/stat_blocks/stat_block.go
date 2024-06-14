package stat_blocks

import "github.com/TimTwigg/EncounterManagerBackend/types/actions"

type StatBlock struct {
	Name             string
	ChallengeRating  int
	Description      EntityDescription
	Stats            NumericalAttributes
	Details          DetailBlock
	DamageModifiers  []DamageModifiers
	Actions          []actions.Action
	BonusActions     []actions.SimpleAction
	Reactions        []actions.SimpleAction
	LegendaryActions []actions.LegendaryAction
}
