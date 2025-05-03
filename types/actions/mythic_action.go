package actions

type MythicAction struct {
	Name        string
	Description string
	Cost        int
}

type Mythic struct {
	Description string
	Actions     []MythicAction
}

func (a MythicAction) Dict() map[string]any {
	return map[string]interface{}{
		"Name":        a.Name,
		"Description": a.Description,
		"Cost":        a.Cost,
	}
}

func (l Mythic) Dict() map[string]any {
	actions := make([]map[string]any, len(l.Actions))
	for i, action := range l.Actions {
		actions[i] = action.Dict()
	}

	return map[string]any{
		"Description": l.Description,
		"Actions":     actions,
	}
}
