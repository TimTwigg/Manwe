package stat_blocks

import "github.com/TimTwigg/EncounterManagerBackend/types/damage"

type DamageModifiers struct {
	Vulnerabilities []damage.DamageType
	Resistances     []damage.DamageType
	Immunities      []damage.DamageType
}
