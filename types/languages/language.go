package language

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TimTwigg/EncounterManagerBackend/utils"
)

type Language struct {
	Language    string
	Description string
}

var DEFAULT_LANGUAGES = utils.LockableMap[string, Language]{}

func initializeLanguage(file_contents string) error {
	language := Language{}
	err := json.Unmarshal([]byte(file_contents), &language)
	if err != nil {
		return err
	}
	DEFAULT_LANGUAGES.Set(language.Language, language)
	return nil
}

func InitializeDefaultLanguages() error {
	err := utils.ApplyToAll("assets/languages", initializeLanguage)
	if err != nil {
		fmt.Println("Error initializing languages!")
		log.Fatal(err)
	}
	DEFAULT_LANGUAGES.Lock()
	fmt.Println("Languages initialized!")
	fmt.Println(DEFAULT_LANGUAGES.ToString())

	return nil
}
