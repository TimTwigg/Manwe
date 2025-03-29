package stat_blocks

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
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

func ParseLanguageInfo(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Languages"})
	if missingKey != nil {
		return LanguageInfo{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from LanguageInfo dictionary! (%v)", *missingKey, dict)}
	}

	// Unpack the Languages array from the dictionary
	languages_raw := lists.UnpackArray(dict["Languages"])
	Languages := make([]string, 0)
	for _, language := range languages_raw {
		Languages = append(Languages, language.(string))
	}

	return LanguageInfo{
		Note:      dict["Note"].(string),
		Languages: Languages,
	}, nil
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

// Parse a Detail Block from a dictionary.
func ParseDetailBlockData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"ArmorType", "Skills", "SavingThrows", "Senses", "Languages", "Traits"})
	if missingKey != nil {
		return DetailBlock{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Detail Block dictionary! (%v)", *missingKey, dict)}
	}

	skills_raw := lists.UnpackArray(dict["Skills"])
	Skills := make([]generics.NumericalItem, 0)
	for _, skill := range skills_raw {
		s, err := parse.PARSERS.Get("NumericalItem")(skill.(map[string]any))
		if err != nil {
			return DetailBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Skill: %s", err)}
		}
		Skills = append(Skills, s.(generics.NumericalItem))
	}

	saving_throws_raw := lists.UnpackArray(dict["SavingThrows"])
	SavingThrows := make([]generics.NumericalItem, 0)
	for _, saving_throw := range saving_throws_raw {
		s, err := parse.PARSERS.Get("NumericalItem")(saving_throw.(map[string]any))
		if err != nil {
			return DetailBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Saving Throw: %s", err)}
		}
		SavingThrows = append(SavingThrows, s.(generics.NumericalItem))
	}

	senses_raw := lists.UnpackArray(dict["Senses"])
	Senses := make([]generics.NumericalItem, 0)
	for _, sense := range senses_raw {
		s, err := parse.PARSERS.Get("NumericalItem")(sense.(map[string]any))
		if err != nil {
			return DetailBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Sense: %s", err)}
		}
		Senses = append(Senses, s.(generics.NumericalItem))
	}

	Languages, err := ParseLanguageInfo(dict["Languages"].(map[string]any))
	if err != nil {
		return DetailBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Languages: %s", err)}
	}

	traits_raw := lists.UnpackArray(dict["Traits"])
	Traits := make([]generics.SimpleItem, 0)
	for _, trait := range traits_raw {
		t, err := parse.PARSERS.Get("SimpleItem")(trait.(map[string]any))
		if err != nil {
			return DetailBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Trait: %s", err)}
		}
		Traits = append(Traits, t.(generics.SimpleItem))
	}

	return DetailBlock{
		ArmorType:    dict["ArmorType"].(string),
		Skills:       Skills,
		SavingThrows: SavingThrows,
		Senses:       Senses,
		Languages:    Languages.(LanguageInfo),
		Traits:       Traits,
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("DetailBlock", ParseDetailBlockData)
}
