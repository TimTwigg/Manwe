package damage

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TimTwigg/EncounterManagerBackend/utils"
	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
)

type DamageType struct {
	DamageType  string
	Description string
}

var DEFAULT_DAMAGE_TYPES = data_type_utils.LockableMap[string, DamageType]{}

func initializeDamageType(file_contents string) error {
	damageType := DamageType{}
	err := json.Unmarshal([]byte(file_contents), &damageType)
	if err != nil {
		return err
	}
	DEFAULT_DAMAGE_TYPES.Set(damageType.DamageType, damageType)
	return nil
}

func InitializeDefaultDamageTypes() error {
	err := utils.ApplyToAllFiles("assets/damage_types", initializeDamageType)
	if err != nil {
		fmt.Println("Error initializing damage types!")
		log.Fatal(err)
	}

	DEFAULT_DAMAGE_TYPES.Lock()
	fmt.Println("Damage types initialized!")

	return nil
}
