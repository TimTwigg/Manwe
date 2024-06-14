package actions

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

type Action struct {
	Name          string
	AttackType    string
	ToHitModifier int
	Reach         int
	Target        string
	DamageAmount  string
	DamageType    string
	Description   string
}

func (a Action) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":     "Action",
		"name":          a.Name,
		"attack_type":   a.AttackType,
		"to_hit_mod":    a.ToHitModifier,
		"reach":         a.Reach,
		"target":        a.Target,
		"damage_amount": a.DamageAmount,
		"damage_type":   a.DamageType,
		"description":   a.Description,
	}
}

// Parse an Action from a dictionary.
func ParseActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"name", "attack_type", "to_hit_mod", "reach", "target", "damage_amount", "damage_type", "description"})
	if missingKey != nil {
		return Action{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Action dictionary! (%v)", *missingKey, dict)}
	}

	return Action{
		Name:          dict["name"].(string),
		AttackType:    dict["attack_type"].(string),
		ToHitModifier: dict["to_hit_mod"].(int),
		Reach:         dict["reach"].(int),
		Target:        dict["target"].(string),
		DamageAmount:  dict["damage_amount"].(string),
		DamageType:    dict["damage_type"].(string),
		Description:   dict["description"].(string),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("Action", ParseActionData)
}
