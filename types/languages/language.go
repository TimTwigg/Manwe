package damage

import (
	"database/sql"

	"github.com/TimTwigg/EncounterManagerBackend/utils"
)

type Language struct {
	Language    string
	Description string
}

var LANGUAGES = map[string]Language{}

func InitializeLanguages() {
	db := utils.Must(sql.Open("sqlite3", "./database/database.sqlite3"))
	rows := utils.Must(db.Query("SELECT * FROM languages"))
	defer rows.Close()
	for rows.Next() {
		var language Language
		err := rows.Scan(&language.Language, &language.Description)
		if err != nil {
			panic(err)
		}
		LANGUAGES[language.Language] = language
	}
}
