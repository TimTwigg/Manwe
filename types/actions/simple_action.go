package actions

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

type SimpleAction struct {
	Name        string
	Description string
}

func (a SimpleAction) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":   "SimpleAction",
		"name":        a.Name,
		"description": a.Description,
	}
}

// Parse a Simple Action from a dictionary.
func ParseSimpleActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"name", "description"})
	if missingKey != nil {
		return SimpleAction{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Simple Action dictionary! (%v)", *missingKey, dict)}
	}

	return SimpleAction{
		Name:        dict["name"].(string),
		Description: dict["description"].(string),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("SimpleAction", ParseSimpleActionData)
}
