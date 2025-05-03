package generics

type SimpleItem struct {
	Name        string
	Description string
}

func (i SimpleItem) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Name":        i.Name,
		"Description": i.Description,
	}
}
