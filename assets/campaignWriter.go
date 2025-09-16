package assets

import (
	"context"
	"strconv"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	campaign "github.com/TimTwigg/Manwe/types/campaign"
	player "github.com/TimTwigg/Manwe/types/player"
	utils "github.com/TimTwigg/Manwe/utils/functions"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

// Check if a campaign exists in the database for a given user
func campaignExists(campaignID int, userid string) bool {
	var count int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT COUNT(*) FROM public.campaign WHERE id = $1 AND username = $2", campaignID, userid).Scan(&count)
	if err != nil {
		logger.Error("Error checking Campaign: " + err.Error())
		return false
	}
	return count > 0
}

// Update the CampaignEntities table with the provided entities for a specific campaign and user
func setCampaignEntities(entities []player.Player, campaignID int, userid string) error {
	_, err := asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.campaignentities WHERE id = $1 AND username = $2", campaignID, userid)
	if err != nil {
		logger.Error("Error deleting CampaignEntities: " + err.Error())
		return errors.Wrap(err, "Error deleting CampaignEntities")
	}

	for row, entity := range entities {
		_, err := asset_utils.DBPool.Exec(context.Background(), "INSERT INTO public.campaignentities (id, username, rowid, statblockid, notes) VALUES ($1, $2, $3, $4, $5)", campaignID, userid, row+1, entity.StatBlock.ID, entity.Notes)
		if err != nil {
			logger.Error("Error inserting CampaignEntity: " + err.Error())
			return errors.Wrap(err, "Error inserting CampaignEntity")
		}
	}
	return nil
}

func SetCampaign(campaignData campaign.Campaign, userid string) (campaign.Campaign, error) {
	if campaignExists(campaignData.ID, userid) {
		_, err := asset_utils.DBPool.Exec(context.Background(), "UPDATE public.campaign SET name = $1 AND description = $2 AND lastModified = $3 WHERE id = $4 AND username = $5", campaignData.Name, campaignData.Description, utils.FormatDate(campaignData.LastModified), campaignData.ID, userid)
		if err != nil {
			logger.Error("Error updating Campaign: " + err.Error())
			return campaign.Campaign{}, errors.Wrap(err, "Error updating Campaign")
		}
	} else {
		err := asset_utils.DBPool.QueryRow(context.Background(), "INSERT INTO public.campaign (username, name, description, creationDate, lastModified) VALUES ($1, $2, $3, $4) RETURNING id", userid, campaignData.Name, campaignData.Description, utils.FormatDate(campaignData.CreationDate), utils.FormatDate(campaignData.LastModified)).Scan(&campaignData.ID)
		if err != nil {
			logger.Error("Error inserting Campaign: " + err.Error())
			return campaign.Campaign{}, errors.Wrap(err, "Error inserting Campaign")
		}
	}

	err := setCampaignEntities(campaignData.Players, campaignData.ID, userid)
	if err != nil {
		logger.Error("Error setting Campaign Entities: " + err.Error())
		return campaign.Campaign{}, errors.Wrap(err, "Error setting Campaign Entities")
	}

	return campaignData, nil
}

func DeleteCampaign(campaignID int, userid string) error {
	if !campaignExists(campaignID, userid) {
		logger.Error("Campaign does not exist: " + strconv.Itoa(campaignID))
		return errors.New("Campaign does not exist: " + strconv.Itoa(campaignID))
	}

	_, err := asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.campaign WHERE id = $1 AND username = $2", campaignID, userid)
	if err != nil {
		logger.Error("Error deleting Campaign: " + err.Error())
		return errors.Wrap(err, "Error deleting Campaign")
	}

	return nil
}
