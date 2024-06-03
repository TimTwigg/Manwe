package damage

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"

	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

var DEFAULT_DAMAGE_TYPES = data_type_utils.LockableMap[string, DamageType]{}

type DamageType struct {
	DamageType  string
	Description string
}

func (d DamageType) Dict() map[string]any {
	return map[string]any{
		"data_type":   "DamageType",
		"damage_type": d.DamageType,
		"description": d.Description,
	}
}

func ParseDamageType(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"damage_type", "description"})
	if missingKey != nil {
		return DamageType{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from DamageType dictionary!", *missingKey)}
	}

	return DamageType{
		DamageType:  dict["damage_type"].(string),
		Description: dict["description"].(string),
	}, nil
}

func InitializeDefaultDamageTypes() error {
	damageTypes, err := parse.ParseAllFilesInFolder("assets/damage_types", ParseDamageType)
	if err != nil {
		return err
	}
	for _, damageType := range damageTypes {
		DEFAULT_DAMAGE_TYPES.Set(damageType.(DamageType).DamageType, damageType.(DamageType))
	}
	return nil
}

// Register the parser with the parser map.
func init() {
	parse.PARSERS.Set("DamageType", ParseDamageType)
}
