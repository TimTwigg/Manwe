package damage

type DamageType struct {
	DamageType  string
	Description string
}

func (d DamageType) Dict() map[string]any {
	return map[string]any{
		"DamageType":  d.DamageType,
		"Description": d.Description,
	}
}
