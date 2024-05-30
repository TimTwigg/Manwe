package stat_blocks

import "github.com/TimTwigg/EncounterManagerBackend/damage"

type DamageModifiers struct {
	Vulnerabilities []damage.DamageType
	Resistances     []damage.DamageType
	Immunities      []damage.DamageType
}
