package actions

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
)

type MythicAction struct {
	Name        string
	Description string
	Cost        int
}

type Mythic struct {
	Description string
	Actions     []MythicAction
}

func (a MythicAction) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":   "MythicAction",
		"Name":        a.Name,
		"Description": a.Description,
		"Cost":        a.Cost,
	}
}

func (l Mythic) Dict() map[string]any {
	actions := make([]map[string]any, len(l.Actions))
	for i, action := range l.Actions {
		actions[i] = action.Dict()
	}

	return map[string]any{
		"Description": l.Description,
		"Actions":     actions,
	}
}

func ParseMythicActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "Description", "Cost"})
	if missingKey != nil {
		return MythicAction{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from MythicAction dictionary! (%v)", *missingKey, dict)}
	}

	return MythicAction{
		Name:        dict["Name"].(string),
		Description: dict["Description"].(string),
		Cost:        int(dict["Cost"].(float64)),
	}, nil
}

// Parse a Mythic Action from a dictionary.
func ParseMythicData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Description", "Actions"})
	if missingKey != nil {
		return Mythic{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Mythic dictionary! (%v)", *missingKey, dict)}
	}

	actions_raw := lists.UnpackArray(dict["Actions"])
	Actions := make([]MythicAction, 0)
	for _, action := range actions_raw {
		missingKey := errors.ValidateKeyExistance(action.(map[string]any), []string{"Name", "Description", "Cost"})
		if missingKey != nil {
			return Mythic{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from MythicAction dictionary! (%v)", *missingKey, action)}
		}
		act, err := ParseMythicActionData(action.(map[string]any))
		if err != nil {
			return Mythic{}, err
		}
		Actions = append(Actions, act.(MythicAction))
	}

	return Mythic{
		Description: dict["Description"].(string),
		Actions:     Actions,
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("Mythic", ParseMythicData)
}
