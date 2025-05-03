package condition

type Condition struct {
	Name    string
	Effects []string
}

// Turn a condition into a dictionary
func (condition Condition) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Name":    condition.Name,
		"Effects": condition.Effects,
	}
}
