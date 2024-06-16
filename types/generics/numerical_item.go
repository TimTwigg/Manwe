package generics

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

type NumericalItem struct {
	Name     string
	Modifier int
}

func (i NumericalItem) Dict() map[string]interface{} {
	return map[string]interface{}{
		"data_type": "NumericalItem",
		"Name":      i.Name,
		"Modifier":  i.Modifier,
	}
}

// Parse a Numerical Item from a dictionary.
func ParseNumericalItemData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "Modifier"})
	if missingKey != nil {
		return NumericalItem{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Numerical Item dictionary! (%v)", *missingKey, dict)}
	}

	return NumericalItem{
		Name:     dict["Name"].(string),
		Modifier: int(dict["Modifier"].(float64)),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("NumericalItem", ParseNumericalItemData)
}
