package stat_blocks

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
)

type Lair struct {
	Name            string
	Description     string
	Initiative      int
	Actions         generics.ItemList
	RegionalEffects generics.ItemList
}

func (l Lair) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Name":            l.Name,
		"Description":     l.Description,
		"Initiative":      l.Initiative,
		"Actions":         l.Actions,
		"RegionalEffects": l.RegionalEffects,
	}
}

// Parse a Lair from a dictionary.
func ParseLairData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Description"})
	if missingKey != nil {
		return Lair{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Lair dictionary! (%v)", *missingKey, dict)}
	}

	var Actions parse.Parseable
	var err error
	if _, ok := dict["Actions"]; ok {
		Actions, err = parse.PARSERS.Get("ItemList")(dict["Actions"].(map[string]interface{}))
		if err != nil {
			return Lair{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Actions: %s", err)}
		}
	} else {
		Actions = generics.ItemList{}
	}

	var RegionalEffects parse.Parseable
	if _, ok := dict["RegionalEffects"]; ok {
		RegionalEffects, err = parse.PARSERS.Get("ItemList")(dict["RegionalEffects"].(map[string]interface{}))
		if err != nil {
			return Lair{}, errors.ParseError{Message: fmt.Sprintf("Error parsing RegionalEffects: %s", err)}
		}
	} else {
		RegionalEffects = generics.ItemList{}
	}

	return Lair{
		Name:            dict["Name"].(string),
		Description:     dict["Description"].(string),
		Initiative:      utils.GetOptional(dict, "Initiative", 0),
		Actions:         Actions.(generics.ItemList),
		RegionalEffects: RegionalEffects.(generics.ItemList),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("Lair", ParseLairData)
}
