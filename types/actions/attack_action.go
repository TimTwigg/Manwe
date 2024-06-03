package actions

type ActionData struct {
	Name          string
	AttackType    string
	ToHitModifier int
	Reach         int
	Target        string
	DamageAmount  string
	DamageType    string
	Description   string
}

func (a ActionData) Dict() map[string]any {
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

func ParseActionData(dict map[string]any) ActionData {
	return ActionData{
		Name:          dict["name"].(string),
		AttackType:    dict["attack_type"].(string),
		ToHitModifier: dict["to_hit_mod"].(int),
		Reach:         dict["reach"].(int),
		Target:        dict["target"].(string),
		DamageAmount:  dict["damage_amount"].(string),
		DamageType:    dict["damage_type"].(string),
		Description:   dict["description"].(string),
	}
}
