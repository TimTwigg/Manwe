package language

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"

	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
)

var DEFAULT_LANGUAGES = data_type_utils.LockableMap[string, Language]{}

func init() {
	parse.PARSERS.Set("Language", ParseLanguage)
}

type Language struct {
	Language    string
	Description string
}

func (language Language) Dict() map[string]any {
	return map[string]any{
		"data_type":   "Language",
		"language":    language.Language,
		"description": language.Description,
	}
}

func ParseLanguage(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"language", "description"})
	if missingKey != nil {
		return Language{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Language dictionary!", *missingKey)}
	}

	return Language{
		Language:    dict["language"].(string),
		Description: dict["description"].(string),
	}, nil
}

func InitializeDefaultLanguages() error {
	languages, err := parse.ParseAllFilesInFolder("assets/languages", ParseLanguage)
	if err != nil {
		return err
	}
	for _, language := range languages {
		DEFAULT_LANGUAGES.Set(language.(Language).Language, language.(Language))
	}
	return nil
}

func init() {
	parse.PARSERS.Set("Language", ParseLanguage)
}
