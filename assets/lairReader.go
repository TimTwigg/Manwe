package assets

import (
	"context"
	"strconv"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	generics "github.com/TimTwigg/Manwe/types/generics"
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
	logger "github.com/TimTwigg/Manwe/utils/log"
	pgx "github.com/jackc/pgx/v5"
	errors "github.com/pkg/errors"
)

func ReadLairByEntityID(id int) (stat_blocks.Lair, error) {
	var lair stat_blocks.Lair
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT name, description, initiative FROM public.lair WHERE statblockid = $1", id).Scan(&lair.Name, &lair.Description, &lair.Initiative)
	if pgx.ErrNoRows == err {
		logger.Error("No Lair found with ID: " + strconv.Itoa(id))
		return stat_blocks.Lair{}, errors.New("No Lair found with ID: " + strconv.Itoa(id))
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return stat_blocks.Lair{}, errors.Wrap(err, "Error querying Lair by ID")
	}
	lair.OwningEntityDBID = id
	lair.Actions = generics.ItemList{Description: "", Items: make([]generics.SimpleItem, 0)}
	lair.RegionalEffects = generics.ItemList{Description: "", Items: make([]generics.SimpleItem, 0)}

	lair_actions_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT name, description, isregional FROM public.lairactionv WHERE statblockid = $1", id)
	var Name string
	var Description string
	var IsRegional bool
	_, err = pgx.ForEachRow(lair_actions_rows, []any{&Name, &Description, &IsRegional}, func() error {
		if Name == "X" {
			if IsRegional {
				lair.RegionalEffects.Description = Description
			} else {
				lair.Actions.Description = Description
			}
		} else {
			if IsRegional {
				lair.RegionalEffects.Items = append(lair.RegionalEffects.Items, generics.SimpleItem{Name: Name, Description: Description})
			} else {
				lair.Actions.Items = append(lair.Actions.Items, generics.SimpleItem{Name: Name, Description: Description})
			}
		}
		return nil
	})
	err = lair_actions_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading Lair Actions from database: " + err.Error())
		return stat_blocks.Lair{}, errors.Wrap(err, "Error reading Lair Actions from database")
	}

	return lair, nil
}
