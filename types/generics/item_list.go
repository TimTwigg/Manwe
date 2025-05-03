package generics

type ItemList struct {
	Description string
	Items       []SimpleItem
}

func (i ItemList) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Description": i.Description,
		"Items":       i.Items,
	}
}
