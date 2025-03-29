package stat_blocks

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	actions "github.com/TimTwigg/EncounterManagerBackend/types/actions"
	generics "github.com/TimTwigg/EncounterManagerBackend/types/generics"
	data_types "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

type StatBlock struct {
	Name                string
	ChallengeRating     float32
	ProficiencyBonus    int
	Source              string
	Description         EntityDescription
	Stats               NumericalAttributes
	DamageModifiers     DamageModifiers
	ConditionImmunities []string
	Details             DetailBlock
	Actions             []actions.Action
	BonusActions        []generics.SimpleItem
	Reactions           []generics.SimpleItem
	LegendaryActions    actions.Legendary
	MythicActions       actions.Mythic
	Lair                Lair
}

func (sb StatBlock) Dict() map[string]any {
	return map[string]any{
		"Name":                sb.Name,
		"ChallengeRating":     sb.ChallengeRating,
		"ProficiencyBonus":    sb.ProficiencyBonus,
		"Description":         sb.Description,
		"Stats":               sb.Stats,
		"DamageModifiers":     sb.DamageModifiers,
		"ConditionImmunities": sb.ConditionImmunities,
		"Details":             sb.Details,
		"Actions":             sb.Actions,
		"BonusActions":        sb.BonusActions,
		"Reactions":           sb.Reactions,
		"LegendaryActions":    sb.LegendaryActions,
		"MythicActions":       sb.MythicActions,
		"Lair":                sb.Lair,
	}
}

// Parse a Stat Block from a dictionary.
func ParseStatBlockData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name", "ChallengeRating", "ProficiencyBonus", "Description", "Stats", "DamageModifiers", "ConditionImmunities", "Details"})
	if missingKey != nil {
		return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from StatBlock dictionary! (%v)", *missingKey, dict)}
	}

	Description, err := parse.PARSERS.Get("EntityDescription")(dict["Description"].(map[string]any))
	if err != nil {
		return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Description: %s", err)}
	}

	Stats, err := parse.PARSERS.Get("NumericalAttributes")(dict["Stats"].(map[string]any))
	if err != nil {
		return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Stats: %s", err)}
	}

	DamageMods, err := parse.PARSERS.Get("DamageModifiers")(dict["DamageModifiers"].(map[string]any))
	if err != nil {
		return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing DamageModifiers: %s", err)}
	}

	ConditionImmunities_raw := lists.UnpackArray(dict["ConditionImmunities"])
	ConditionImmunities := make([]string, 0)
	for _, conditionImmunity := range ConditionImmunities_raw {
		ConditionImmunities = append(ConditionImmunities, conditionImmunity.(string))
	}

	Details, err := parse.PARSERS.Get("DetailBlock")(dict["Details"].(map[string]any))
	if err != nil {
		return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Details: %s", err)}
	}

	ActionsRaw := lists.UnpackArray(dict["Actions"])
	Actions := make([]actions.Action, 0)
	for _, action := range ActionsRaw {
		actionParsed, err := parse.PARSERS.Get("Action")(action.(map[string]any))
		if err != nil {
			return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Action: %s", err)}
		}
		Actions = append(Actions, actionParsed.(actions.Action))
	}

	BonusActions := make([]generics.SimpleItem, 0)
	if _, ok := dict["BonusActions"]; ok {
		BonusActionsRaw := lists.UnpackArray(dict["BonusActions"])
		for _, bonusAction := range BonusActionsRaw {
			bonusActionParsed, err := parse.PARSERS.Get("SimpleItem")(bonusAction.(map[string]any))
			if err != nil {
				return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing BonusAction: %s", err)}
			}
			BonusActions = append(BonusActions, bonusActionParsed.(generics.SimpleItem))
		}
	}

	Reactions := make([]generics.SimpleItem, 0)
	if _, ok := dict["Reactions"]; ok {
		ReactionsRaw := lists.UnpackArray(dict["Reactions"])
		for _, reaction := range ReactionsRaw {
			reactionParsed, err := parse.PARSERS.Get("SimpleItem")(reaction.(map[string]any))
			if err != nil {
				return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Reaction: %s", err)}
			}
			Reactions = append(Reactions, reactionParsed.(generics.SimpleItem))
		}
	}

	var LegendaryActions parse.Parseable
	if _, ok := dict["LegendaryActions"]; ok {
		LegendaryActions, err = parse.PARSERS.Get("Legendary")(dict["LegendaryActions"].(map[string]any))
		if err != nil {
			return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing LegendaryActions: %s", err)}
		}
	} else {
		LegendaryActions = actions.Legendary{}
	}

	var MythicActions parse.Parseable
	if _, ok := dict["MythicActions"]; ok {
		MythicActions, err = parse.PARSERS.Get("Mythic")(dict["MythicActions"].(map[string]any))
		if err != nil {
			return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing MythicActions: %s", err)}
		}
	} else {
		MythicActions = actions.Mythic{}
	}

	var LairData parse.Parseable
	if _, ok := dict["Lair"]; ok {
		LairData, err = parse.PARSERS.Get("Lair")(dict["Lair"].(map[string]any))
		if err != nil {
			return StatBlock{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Lair: %s", err)}
		}
	} else {
		LairData = Lair{}
	}

	return StatBlock{
		Name:                dict["Name"].(string),
		ChallengeRating:     float32(dict["ChallengeRating"].(float64)),
		ProficiencyBonus:    int(dict["ProficiencyBonus"].(float64)),
		Description:         Description.(EntityDescription),
		Stats:               Stats.(NumericalAttributes),
		DamageModifiers:     DamageMods.(DamageModifiers),
		ConditionImmunities: ConditionImmunities,
		Details:             Details.(DetailBlock),
		Actions:             Actions,
		BonusActions:        BonusActions,
		Reactions:           Reactions,
		LegendaryActions:    LegendaryActions.(actions.Legendary),
		MythicActions:       MythicActions.(actions.Mythic),
		Lair:                LairData.(Lair),
	}, nil
}

var DEFAULT_STAT_BLOCKS = data_types.LockableMap[string, StatBlock]{}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("StatBlock", ParseStatBlockData)

	// Build dictionary of default stat blocks from files in the assets/stat_blocks folder.
	statBlocks, err := parse.ParseAllFilesInFolder("assets/stat_blocks", ParseStatBlockData)
	if err != nil {
		panic(fmt.Errorf("error initializing 'stat_block' objects: %s", err))
	}
	for _, statBlock := range statBlocks {
		DEFAULT_STAT_BLOCKS.Set(statBlock.(StatBlock).Name, statBlock.(StatBlock))
	}
	log.Init("Stat blocks initialized!")
}
