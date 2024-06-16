package stat_blocks

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	"github.com/TimTwigg/EncounterManagerBackend/utils/lists"
)

type DamageModifiers struct {
	Vulnerabilities []string
	Resistances     []string
	Immunities      []string
}

func (d DamageModifiers) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Vulnerabilities": d.Vulnerabilities,
		"Resistances":     d.Resistances,
		"Immunities":      d.Immunities,
	}
}

// Parse a Damage Modifiers from a dictionary.
func ParseDamageModifiersData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Vulnerabilities", "Resistances", "Immunities"})
	if missingKey != nil {
		return DamageModifiers{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Damage Modifiers dictionary! (%v)", *missingKey, dict)}
	}

	vulnerabilities_raw := lists.UnpackArray(dict["Vulnerabilities"])
	vulnerabilities := make([]string, 0)
	for _, vulnerability := range vulnerabilities_raw {
		vulnerabilities = append(vulnerabilities, vulnerability.(string))
	}

	resistances_raw := lists.UnpackArray(dict["Resistances"])
	resistances := make([]string, 0)
	for _, resistance := range resistances_raw {
		resistances = append(resistances, resistance.(string))
	}

	immunities_raw := lists.UnpackArray(dict["Immunities"])
	immunities := make([]string, 0)
	for _, immunity := range immunities_raw {
		immunities = append(immunities, immunity.(string))
	}

	return DamageModifiers{
		Vulnerabilities: vulnerabilities,
		Resistances:     resistances,
		Immunities:      immunities,
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("DamageModifiers", ParseDamageModifiersData)
}
