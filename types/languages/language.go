package language

type Language struct {
	Language    string
	Description string
}

// Turn a language into a dictionary
func (language Language) Dict() map[string]any {
	return map[string]any{
		"Language":    language.Language,
		"Description": language.Description,
	}
}
