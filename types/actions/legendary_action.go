package actions

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

type LegendaryAction struct {
	Name        string
	Description string
	Cost        int
}

func (a LegendaryAction) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":   "LegendaryAction",
		"name":        a.Name,
		"description": a.Description,
		"cost":        a.Cost,
	}
}

// Parse a Legendary Action from a dictionary.
func ParseLegendaryActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"name", "description", "cost"})
	if missingKey != nil {
		return LegendaryAction{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Legendary Action dictionary! (%v)", *missingKey, dict)}
	}

	return LegendaryAction{
		Name:        dict["name"].(string),
		Description: dict["description"].(string),
		Cost:        dict["cost"].(int),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("LegendaryAction", ParseLegendaryActionData)
}
