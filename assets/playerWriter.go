package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	player "github.com/TimTwigg/Manwe/types/player"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

func playerExists(plyr player.Player, userid string) bool {
	var count int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT COUNT(*) FROM public.campaignentities WHERE id = $1 AND username = $2 AND rowid = $3", plyr.CampaignID, userid, plyr.RowID).Scan(&count)
	if err != nil {
		logger.Error("Error checking Player: " + err.Error())
		return false
	}
	return count > 0
}

func SetPlayer(plyr player.Player, userid string) (player.Player, error) {
	if plyr.CampaignID == 0 {
		logger.Error("Campaign is required")
		return player.Player{}, errors.New("Campaign is required")
	}
	if !campaignExists(plyr.CampaignID, userid) {
		return player.Player{}, errors.New("Campaign does not exist")
	}

	sb, err := SetStatblock(plyr.StatBlock, userid)
	if err != nil {
		return player.Player{}, errors.Wrap(err, "Error setting StatBlock")
	}
	plyr.StatBlock = sb

	if plyr.RowID == 0 || !playerExists(plyr, userid) {
		// Create New Player
		err := asset_utils.DBPool.QueryRow(
			context.Background(),
			"INSERT INTO public.campaignentities (id, username, statblockid, notes) VALUES ($1, $2, $3, $4) RETURNING rowid",
			plyr.CampaignID,
			userid,
			plyr.StatBlock.ID,
			plyr.Notes,
		).Scan(&plyr.RowID)
		if err != nil {
			logger.Error("Error inserting Player: " + err.Error())
			return player.Player{}, errors.Wrap(err, "Error inserting Player")
		}
	} else {
		// Update Existing Player
		_, err := asset_utils.DBPool.Exec(
			context.Background(),
			"UPDATE public.campaignentities SET statblockid = $1, notes = $2 WHERE id = $3 AND username = $4 AND rowid = $5",
			plyr.StatBlock.ID,
			plyr.Notes,
			plyr.CampaignID,
			userid,
			plyr.RowID,
		)
		if err != nil {
			logger.Error("Error updating Player: " + err.Error())
			return player.Player{}, errors.Wrap(err, "Error updating Player")
		}
	}

	return plyr, nil
}

func DeletePlayer(campaignID int, rowID int, userid string) error {
	_, err := asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.campaignentities WHERE id = $1 AND username = $2 AND rowid = $3", campaignID, userid, rowID)
	if err != nil {
		logger.Error("Error deleting Player: " + err.Error())
		return errors.Wrap(err, "Error deleting Player")
	}
	return nil
}
