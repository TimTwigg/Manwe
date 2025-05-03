package stat_blocks

import (
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
)

type LanguageInfo struct {
	Note      string
	Languages []string
}

func (d LanguageInfo) Dict() map[string]any {
	return map[string]any{
		"Note":      d.Note,
		"Languages": d.Languages,
	}
}

type DetailBlock struct {
	ArmorType    string
	Skills       []generics.NumericalItem
	SavingThrows []generics.NumericalItem
	Senses       []generics.NumericalItem
	Languages    LanguageInfo
	Traits       []generics.SimpleItem
	SpellSaveDC  int
}

func (d DetailBlock) Dict() map[string]any {
	return map[string]any{
		"ArmorType":    d.ArmorType,
		"Skills":       d.Skills,
		"SavingThrows": d.SavingThrows,
		"Senses":       d.Senses,
		"Languages":    d.Languages,
		"Traits":       d.Traits,
	}
}
