package generics

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

type SimpleItem struct {
	Name        string
	Description string
}

func (i SimpleItem) Dict() map[string]interface{} {
	return map[string]interface{}{
		"data_type":   "SimpleItem",
		"Name":        i.Name,
		"Description": i.Description,
	}
}

// Parse a Simple Item from a dictionary.
func ParseSimpleItemData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "Description"})
	if missingKey != nil {
		return SimpleItem{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Simple Item dictionary! (%v)", *missingKey, dict)}
	}

	return SimpleItem{
		Name:        dict["Name"].(string),
		Description: dict["Description"].(string),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("SimpleItem", ParseSimpleItemData)
}
