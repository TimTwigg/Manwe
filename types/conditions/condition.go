package condition

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"

	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

type Condition struct {
	Name    string
	Effects []string
}

// Turn a condition into a dictionary
func (condition Condition) Dict() map[string]interface{} {
	return map[string]interface{}{
		"data_type": "Condition",
		"Name":      condition.Name,
		"Effects":   condition.Effects,
	}
}

// Parse a condition from a dictionary.
func ParseCondition(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "Effects"})
	if missingKey != nil {
		return Condition{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Condition dictionary! (%v)", *missingKey, dict)}
	}

	effects_raw := lists.UnpackArray(dict["Effects"])
	effects := make([]string, 0)
	for _, effect := range effects_raw {
		effects = append(effects, effect.(string))
	}

	return Condition{
		Name:    dict["Name"].(string),
		Effects: effects,
	}, nil
}

var DEFAULT_CONDITIONS = data_type_utils.LockableMap[string, Condition]{}

func init() {
	// Register the parser with the parser map.
	parse.PARSERS.Set("Condition", ParseCondition)

	// Build dictionary of default conditions from files in the assets/conditions folder.
	conditions, err := parse.ParseAllFilesInFolder("assets/conditions", ParseCondition)
	if err != nil {
		log.Error("Failure while initializing 'condition' objects")
		panic(err)
	}
	for _, condition := range conditions {
		DEFAULT_CONDITIONS.Set(condition.(Condition).Name, condition.(Condition))
	}
	log.Init("Conditions initialized!")
}
