package actions

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
		"Name":                  a.Name,
		"AttackType":            a.AttackType,
		"ToHitModifier":         a.ToHitModifier,
		"Reach":                 a.Reach,
		"Targets":               a.Targets,
		"Damage":                a.Damage,
		"AdditionalDescription": a.AdditionalDescription,
	}
}
