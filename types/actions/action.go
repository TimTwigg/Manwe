package actions

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	utils "github.com/TimTwigg/EncounterManagerBackend/utils/functions"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
)

type AltDamageT struct {
	Amount string
	Type   string
	Note   string
}

type SavingThrowDamageT struct {
	Ability    string
	DC         int
	HalfDamage bool
	Note       string
}

type DamageT struct {
	Amount            string
	Type              string
	AlternativeDamage AltDamageT
	SavingThrow       SavingThrowDamageT
}

type Action struct {
	Name                  string
	AttackType            string
	ToHitModifier         int
	Reach                 int
	Targets               int
	Damage                []DamageT
	AdditionalDescription string
}

func (a AltDamageT) Dict() map[string]any {
	return map[string]interface{}{
		"Amount": a.Amount,
		"Type":   a.Type,
		"Note":   a.Note,
	}
}

func (s SavingThrowDamageT) Dict() map[string]any {
	return map[string]interface{}{
		"Ability":    s.Ability,
		"DC":         s.DC,
		"HalfDamage": s.HalfDamage,
		"Note":       s.Note,
	}
}

func (d DamageT) Dict() map[string]any {
	return map[string]interface{}{
		"Amount":            d.Amount,
		"Type":              d.Type,
		"AlternativeDamage": d.AlternativeDamage.Dict(),
		"SavingThrow":       d.SavingThrow.Dict(),
	}
}

func (a Action) Dict() map[string]any {
	return map[string]interface{}{
		"data_type":             "Action",
		"Name":                  a.Name,
		"AttackType":            a.AttackType,
		"ToHitModifier":         a.ToHitModifier,
		"Reach":                 a.Reach,
		"Targets":               a.Targets,
		"Damage":                a.Damage,
		"AdditionalDescription": a.AdditionalDescription,
	}
}

func ParseAltDamageData(dict map[string]any) (AltDamageT, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Amount", "Type", "Note"})
	if missingKey != nil {
		return AltDamageT{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from AltDamage dictionary! (%v)", *missingKey, dict)}
	}

	return AltDamageT{
		Amount: dict["Amount"].(string),
		Type:   dict["Type"].(string),
		Note:   dict["Note"].(string),
	}, nil
}

func ParseSavingThrowDamageData(dict map[string]any) (SavingThrowDamageT, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Ability", "DC", "HalfDamage"})
	if missingKey != nil {
		return SavingThrowDamageT{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from SavingThrowDamage dictionary! (%v)", *missingKey, dict)}
	}

	return SavingThrowDamageT{
		Ability:    dict["Ability"].(string),
		DC:         int(dict["DC"].(float64)),
		HalfDamage: dict["HalfDamage"].(bool),
		Note:       utils.GetOptional(dict, "Note", ""),
	}, nil
}

func ParseDamageData(dict map[string]any) (DamageT, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Amount", "Type"})
	if missingKey != nil {
		return DamageT{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Damage dictionary! (%v)", *missingKey, dict)}
	}

	var AlternativeDamage AltDamageT
	var err error
	if altdmg, ok := dict["AlternativeDamage"]; ok {
		AlternativeDamage, err = ParseAltDamageData(altdmg.(map[string]any))
	}
	if err != nil {
		return DamageT{}, err
	}

	var SavingThrowDamage SavingThrowDamageT
	if svngthrwdmg, ok := dict["SavingThrow"]; ok {
		SavingThrowDamage, err = ParseSavingThrowDamageData(svngthrwdmg.(map[string]any))
	}
	if err != nil {
		return DamageT{}, err
	}

	return DamageT{
		Amount:            dict["Amount"].(string),
		Type:              dict["Type"].(string),
		AlternativeDamage: AlternativeDamage,
		SavingThrow:       SavingThrowDamage,
	}, nil
}

// Parse an Action from a dictionary.
func ParseActionData(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Name"})
	if missingKey != nil {
		return Action{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Action dictionary! (%v)", *missingKey, dict)}
	}

	Damages := make([]DamageT, 0)
	if dmg, ok := dict["Damage"]; ok {
		damages_raw := lists.UnpackArray(dmg)
		for _, damage := range damages_raw {
			Damage, err := ParseDamageData(damage.(map[string]any))
			if err != nil {
				return Action{}, err
			}
			Damages = append(Damages, Damage)
		}
	}

	return Action{
		Name:                  dict["Name"].(string),
		AttackType:            utils.GetOptional(dict, "AttackType", ""),
		ToHitModifier:         utils.GetOptional(dict, "ToHitModifier", 0),
		Reach:                 utils.GetOptional(dict, "Reach", 0),
		Targets:               utils.GetOptional(dict, "Targets", 0),
		Damage:                Damages,
		AdditionalDescription: utils.GetOptional(dict, "AdditionalDescription", ""),
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("Action", ParseActionData)
}
