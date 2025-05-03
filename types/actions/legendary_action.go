package actions

type LegendaryAction struct {
	Name        string
	Description string
	Cost        int
}

type Legendary struct {
	Points      int
	Description string
	Actions     []LegendaryAction
}

func (a LegendaryAction) Dict() map[string]any {
	return map[string]interface{}{
		"Name":        a.Name,
		"Description": a.Description,
		"Cost":        a.Cost,
	}
}

func (l Legendary) Dict() map[string]any {
	actions := make([]map[string]any, len(l.Actions))
	for i, action := range l.Actions {
		actions[i] = action.Dict()
	}

	return map[string]any{
		"Points":      l.Points,
		"Description": l.Description,
		"Actions":     actions,
	}
}
