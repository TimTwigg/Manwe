package entities

import (
	"github.com/TimTwigg/EncounterManagerBackend/types/generics"
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
)

type EntityType int

const (
	StatBlock EntityType = iota
	Player
)

type SpellSlotLevel struct {
	total int
	used  int
}

type Entity struct {
	DBID             int
	ID               string
	Name             string
	Suffix           string
	Initiative       int
	MaxHitPoints     int
	TempHitPoints    int
	CurrentHitPoints int
	ArmorClass       int
	ArmorClassBonus  int
	Speed            stat_blocks.Speeds
	Conditions       map[string]int
	SpellSaveDC      int
	SpellSlots       map[int]SpellSlotLevel
	Concentration    bool
	Notes            string
	IsHostile        bool
	EncounterLocked  bool
	Displayable      any
	EntityType       EntityType
	SavingThrows     []generics.ProficiencyItem
	ChallengeRating  float32
}
