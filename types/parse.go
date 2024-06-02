package types

import data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"

var PARSERS = data_type_utils.SmartMap[string, func(map[string]any) (Parseable, error)]{}

type Parseable interface {
	Dict() map[string]any
}
