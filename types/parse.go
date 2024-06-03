package parse

import (
	data_type_utils "github.com/TimTwigg/EncounterManagerBackend/utils/data_types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	io_utils "github.com/TimTwigg/EncounterManagerBackend/utils/io"
	list_utils "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
)

var PARSERS = data_type_utils.SmartMap[string, func(map[string]any) (Parseable, error)]{}

type Parseable interface {
	Dict() map[string]any
}

func BuildDataType(data map[string]any) (Parseable, error) {
	dataType, ok := data["data_type"].(string)
	if !ok {
		return nil, errors.ParseError{Message: "JSON Object has no 'data_type' key to point to the reconstructor!"}
	}
	return PARSERS.Get(dataType)(data)
}

func ReadDataTypeFromFile(address string) (Parseable, error) {
	data, err := io_utils.ReadJSON(address)
	if err != nil {
		return nil, err
	}
	return BuildDataType(data)
}

func ParseAllFilesInFolder(address string, parser func(a map[string]any) (Parseable, error)) ([]Parseable, error) {
	defaultFiles, err := io_utils.ListDir(address)
	if err != nil {
		return nil, err
	}

	dictedFiles, err := list_utils.MapWithError(defaultFiles, io_utils.ReadJSON)
	if err != nil {
		return nil, err
	}

	parsedFiles, err := list_utils.MapWithError(dictedFiles, parser)
	if err != nil {
		return nil, err
	}

	return parsedFiles, nil
}
