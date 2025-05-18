package generics

type ProficiencyItem struct {
	Name     string
	Level    int
	Override int
}

func (i ProficiencyItem) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Name":     i.Name,
		"Level":    i.Level,
		"Override": i.Override,
	}
}
