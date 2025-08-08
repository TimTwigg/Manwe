package stat_blocks

type EntityDescription struct {
	Size      string
	Type      string
	Alignment string
	Category  string
}

func (e EntityDescription) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Size":      e.Size,
		"Type":      e.Type,
		"Alignment": e.Alignment,
		"Category":  e.Category,
	}
}
