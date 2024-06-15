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

type Legendary struct {
	Points      int
	Description string
	Actions     []LegendaryAction
}

func (a LegendaryAction) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":   "LegendaryAction",
		"Name":        a.Name,
		"Description": a.Description,
		"Cost":        a.Cost,
	}
}

func (l Legendary) Dict() map[string]any {
	actions := make([]map[string]any, len(l.Actions))
	for i, action := range l.Actions {
		actions[i] = action.Dict()
	}

	return map[string]any{
		"Points":      l.Points,
		"Description": l.Description,
		"Actions":     actions,
	}
}

// Parse a Legendary Action from a dictionary.
func ParseLegendaryActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Points", "Description", "Actions"})
	if missingKey != nil {
		return Legendary{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Legendary dictionary! (%v)", *missingKey, dict)}
	}

	actions := dict["Actions"].([]map[string]any)
	for _, action := range actions {
		missingKey := errors.ValidateKeyExistance(action, []string{"Name", "Description", "Cost"})
		if missingKey != nil {
			return Legendary{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from LegendaryAction dictionary! (%v)", *missingKey, action)}
		}
	}

	return Legendary{
		Points:      dict["Points"].(int),
		Description: dict["Description"].(string),
		Actions:     dict["Actions"].([]LegendaryAction),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("LegendaryAction", ParseLegendaryActionData)
}
