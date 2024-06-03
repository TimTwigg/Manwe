package damage

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"

	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

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

var DEFAULT_DAMAGE_TYPES = data_type_utils.LockableMap[string, DamageType]{}

func init() {
	// Register the parser with the parser map.
	parse.PARSERS.Set("DamageType", ParseDamageType)

	// Build dictionary of default damage types from files in the assets/damage_types folder.
	damageTypes, err := parse.ParseAllFilesInFolder("assets/damage_types", ParseDamageType)
	if err != nil {
		panic(fmt.Errorf("error initializing 'damage_type' objects: %s", err))
	}
	for _, damageType := range damageTypes {
		DEFAULT_DAMAGE_TYPES.Set(damageType.(DamageType).DamageType, damageType.(DamageType))
	}
	log.Init("Damage types initialized!")
}
