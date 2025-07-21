package assets

import (
	"context"
	"strconv"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	generics "github.com/TimTwigg/Manwe/types/generics"
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func ReadLairByEntityID(id int) (stat_blocks.Lair, error) {
	lair_row, err := asset_utils.DBPool.Query(context.Background(), "SELECT name, description, initiative FROM public.lair WHERE statblockid = $1", id)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.Lair{}, error_utils.ParseError{Message: err.Error()}
	}
	defer lair_row.Close()
	// Read row from Lair table
	if lair_row.Next() {
		var Name, Description string
		var Initiative int
		if err := lair_row.Scan(
			&Name,
			&Description,
			&Initiative,
		); err != nil {
			logger.Error("Error Scanning Lair Row: " + err.Error())
			return stat_blocks.Lair{}, error_utils.ParseError{Message: err.Error()}
		}
		block := stat_blocks.Lair{Name: Name, OwningEntityDBID: id, Description: Description, Initiative: Initiative, Actions: generics.ItemList{Description: "", Items: make([]generics.SimpleItem, 0)}, RegionalEffects: generics.ItemList{Description: "", Items: make([]generics.SimpleItem, 0)}}

		// Read Lair Actions
		lair_actions_row, err := asset_utils.DBPool.Query(context.Background(), "SELECT name, description, isregional FROM public.lairactionv WHERE statblockid = $1", id)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return stat_blocks.Lair{}, error_utils.ParseError{Message: err.Error()}
		}
		defer lair_actions_row.Close()
		// Read row from LairActions table
		for lair_actions_row.Next() {
			var Name string
			var Description string
			var IsRegional bool
			if err := lair_actions_row.Scan(
				&Name,
				&Description,
				&IsRegional,
			); err != nil {
				logger.Error("Error Scanning Lair Action Row: " + err.Error())
				return stat_blocks.Lair{}, error_utils.ParseError{Message: err.Error()}
			}
			// Add to StatBlock
			if Name == "X" {
				if IsRegional {
					block.RegionalEffects.Description = Description
				} else {
					block.Actions.Description = Description
				}
			} else {
				if IsRegional {
					block.RegionalEffects.Items = append(block.RegionalEffects.Items, generics.SimpleItem{Name: Name, Description: Description})
				} else {
					block.Actions.Items = append(block.Actions.Items, generics.SimpleItem{Name: Name, Description: Description})
				}
			}
		}
		return block, nil
	} else {
		logger.Error("No Lair found for Entity ID: " + strconv.Itoa(id))
		return stat_blocks.Lair{}, error_utils.ParseError{Message: "No Lair found for Entity ID: " + strconv.Itoa(id)}
	}
}
