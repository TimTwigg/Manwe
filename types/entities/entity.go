package entities

import (
	stat_blocks "github.com/TimTwigg/EncounterManagerBackend/types/stat_blocks"
)

type EntityType int

const (
	StatBlock EntityType = iota
	Player
)

type Entity struct {
	id               string
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
	SpellSlots       map[int]chan struct {
		total int
		used  int
	}
	Concentration bool
	Reactions     chan struct {
		total int
		used  int
	}
	Notes           string
	IsHostile       bool
	EncounterLocked bool
	Displayable     any
	EntityType      EntityType
	SavingThrows    chan struct {
		Strength     int
		Dexterity    int
		Constitution int
		Intelligence int
		Wisdom       int
		Charisma     int
	}
	DifficultyRating float32
}
