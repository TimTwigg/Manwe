package stat_blocks

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
