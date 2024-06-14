package language

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"

	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	log "github.com/TimTwigg/EncounterManagerBackend/utils/log"
)

type Language struct {
	Language    string
	Description string
}

// Turn a language into a dictionary
func (language Language) Dict() map[string]any {
	return map[string]any{
		"data_type":   "Language",
		"language":    language.Language,
		"description": language.Description,
	}
}

// Parse a language from a dictionary.
func ParseLanguage(dict map[string]any) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"language", "description"})
	if missingKey != nil {
		return Language{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Language dictionary! (%v)", *missingKey, dict)}
	}

	return Language{
		Language:    dict["language"].(string),
		Description: dict["description"].(string),
	}, nil
}

var DEFAULT_LANGUAGES = data_type_utils.LockableMap[string, Language]{}

func init() {
	// Register the parser with the parser map.
	parse.PARSERS.Set("Language", ParseLanguage)

	// Build dictionary of default languages from files in the assets/languages folder.
	languages, err := parse.ParseAllFilesInFolder("assets/languages", ParseLanguage)
	if err != nil {
		log.Error("Failure while initializing 'language' objects")
		panic(err)
	}
	for _, language := range languages {
		DEFAULT_LANGUAGES.Set(language.(Language).Language, language.(Language))
	}
	log.Init("Languages initialized!")
}
