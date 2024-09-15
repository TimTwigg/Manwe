package stat_blocks

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
)

type Speeds struct {
	Walk   int
	Fly    int
	Swim   int
	Climb  int
	Burrow int
}

type HitPointsT struct {
	Average int
	Dice    string
}

type NumericalAttributes struct {
	ArmorClass   int
	HitPoints    HitPointsT
	Speed        Speeds
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}

func (s Speeds) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Walk":   s.Walk,
		"Fly":    s.Fly,
		"Swim":   s.Swim,
		"Climb":  s.Climb,
		"Burrow": s.Burrow,
	}
}

func (h HitPointsT) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Average": h.Average,
		"Dice":    h.Dice,
	}
}

func (n NumericalAttributes) Dict() map[string]interface{} {
	return map[string]interface{}{
		"ArmorClass":   n.ArmorClass,
		"HitPoints":    n.HitPoints,
		"Speed":        n.Speed,
		"Strength":     n.Strength,
		"Dexterity":    n.Dexterity,
		"Constitution": n.Constitution,
		"Intelligence": n.Intelligence,
		"Wisdom":       n.Wisdom,
		"Charisma":     n.Charisma,
	}
}

func ParseSpeedsData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Walk"})
	if missingKey != nil {
		return Speeds{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Speeds dictionary! (%v)", *missingKey, dict)}
	}

	return Speeds{
		Walk:   int(dict["Walk"].(float64)),
		Fly:    utils.GetOptional(dict, "Fly", 0),
		Swim:   utils.GetOptional(dict, "Swim", 0),
		Climb:  utils.GetOptional(dict, "Climb", 0),
		Burrow: utils.GetOptional(dict, "Burrow", 0),
	}, nil
}

func ParseHitPointsData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Average", "Dice"})
	if missingKey != nil {
		return HitPointsT{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from HitPoints dictionary! (%v)", *missingKey, dict)}
	}

	return HitPointsT{
		Average: int(dict["Average"].(float64)),
		Dice:    dict["Dice"].(string),
	}, nil
}

// Parse a Numerical Attributes from a dictionary.
func ParseNumericalAttributesData(dict map[string]interface{}) (parse.Parseable, error) {
	speedDict, err := ParseSpeedsData(dict["Speed"].(map[string]interface{}))
	if err != nil {
		return NumericalAttributes{}, errors.ParseError{Message: fmt.Sprintf("Speed key missing from Numerical Attributes dictionary! (%v)", dict)}
	}

	hitPointsDict, err := ParseHitPointsData(dict["HitPoints"].(map[string]interface{}))
	if err != nil {
		return NumericalAttributes{}, errors.ParseError{Message: fmt.Sprintf("HitPoints key missing from Numerical Attributes dictionary! (%v)", dict)}
	}

	missingKey := errors.ValidateKeyExistance(dict, []string{"ArmorClass", "HitPoints", "Speed", "Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"})
	if missingKey != nil {
		return NumericalAttributes{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Numerical Attributes dictionary! (%v)", *missingKey, dict)}
	}

	return NumericalAttributes{
		ArmorClass:   int(dict["ArmorClass"].(float64)),
		HitPoints:    hitPointsDict.(HitPointsT),
		Speed:        speedDict.(Speeds),
		Strength:     int(dict["Strength"].(float64)),
		Dexterity:    int(dict["Dexterity"].(float64)),
		Constitution: int(dict["Constitution"].(float64)),
		Intelligence: int(dict["Intelligence"].(float64)),
		Wisdom:       int(dict["Wisdom"].(float64)),
		Charisma:     int(dict["Charisma"].(float64)),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("NumericalAttributes", ParseNumericalAttributesData)
}
