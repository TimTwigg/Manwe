package stat_blocks

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

type EntityDescription struct {
	Size      string
	Type      string
	Alignment string
}

func (e EntityDescription) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Size":      e.Size,
		"Type":      e.Type,
		"Alignment": e.Alignment,
	}
}

// Parse an Entity Description from a dictionary.
func ParseEntityDescriptionData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Size", "Type", "Alignment"})
	if missingKey != nil {
		return EntityDescription{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Entity Description dictionary! (%v)", *missingKey, dict)}
	}

	return EntityDescription{
		Size:      dict["Size"].(string),
		Type:      dict["Type"].(string),
		Alignment: dict["Alignment"].(string),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("EntityDescription", ParseEntityDescriptionData)
}
