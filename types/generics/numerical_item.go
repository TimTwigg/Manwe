package generics

type NumericalItem struct {
	Name     string
	Modifier int
}

func (i NumericalItem) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Name":     i.Name,
		"Modifier": i.Modifier,
	}
}
