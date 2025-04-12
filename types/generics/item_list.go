package generics

import (
	"fmt"

	parse "github.com/TimTwigg/EncounterManagerBackend/types"
	errors "github.com/TimTwigg/EncounterManagerBackend/utils/errors"
	lists "github.com/TimTwigg/EncounterManagerBackend/utils/lists"
)

type ItemList struct {
	Description string
	Items       []SimpleItem
}

func (i ItemList) Dict() map[string]interface{} {
	return map[string]interface{}{
		"data_type":   "ItemList",
		"Description": i.Description,
		"Items":       i.Items,
	}
}

// Parse a Item List from a dictionary.
func ParseItemListData(dict map[string]interface{}) (parse.Parseable, error) {
	missingKey := errors.ValidateKeyExistance(dict, []string{"Description", "Items"})
	if missingKey != nil {
		return ItemList{}, errors.ParseError{Message: fmt.Sprintf("Key '%s' missing from Item List dictionary! (%v)", *missingKey, dict)}
	}

	items_raw := lists.UnpackArray(dict["Items"])
	items := make([]SimpleItem, len(items_raw))
	for _, item := range items_raw {
		s, err := parse.PARSERS.Get("SimpleItem")(item.(map[string]any))
		if err != nil {
			return ItemList{}, errors.ParseError{Message: fmt.Sprintf("Error parsing Item: %s", err)}
		}
		items = append(items, s.(SimpleItem))
	}

	return ItemList{
		Description: dict["Description"].(string),
		Items:       items,
	}, nil
}

func init() {
	// register the parser with the parser map.
	parse.PARSERS.Set("ItemList", ParseItemListData)
}
