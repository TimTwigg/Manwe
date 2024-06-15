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
		"Name":        a.Name,
		"Description": a.Description,
	}
}

// Parse a Simple Action from a dictionary.
func ParseSimpleActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "Description"})
	if missingKey != nil {
		return SimpleAction{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Simple Action dictionary! (%v)", *missingKey, dict)}
	}

	return SimpleAction{
		Name:        dict["Name"].(string),
		Description: dict["Description"].(string),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("SimpleAction", ParseSimpleActionData)
}
