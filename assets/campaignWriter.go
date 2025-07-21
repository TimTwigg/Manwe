package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	campaign "github.com/TimTwigg/Manwe/types/campaign"
	player "github.com/TimTwigg/Manwe/types/player"
	utils "github.com/TimTwigg/Manwe/utils/functions"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

// Check if a campaign exists in the database for a given user
func campaignExists(campaignName string, userid string) bool {
	rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT COUNT(*) FROM public.campaigns WHERE campaign = $1 AND username = $2", campaignName, userid)
	if err != nil {
		logger.Error("Error checking Campaign: " + err.Error())
		return false
	}
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			logger.Error("Error scanning count from Campaign: " + err.Error())
			return false
		}
	}
	return count > 0
}

// Update the CampaignEntities table with the provided entities for a specific campaign and user
func setCampaignEntities(entities []player.Player, campaignName string, userid string) error {
	_, err := asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.campaignentities WHERE campaign = $1 AND username = $2", campaignName, userid)
	if err != nil {
		logger.Error("Error deleting CampaignEntities: " + err.Error())
		return err
	}

	for row, entity := range entities {
		_, err := asset_utils.DBPool.Exec(context.Background(), "INSERT INTO public.campaignentities (campaign, username, rowid, statblockid, notes) VALUES ($1, $2, $3, $4, $5)", campaignName, userid, row+1, entity.StatBlock.ID, entity.Notes)
		if err != nil {
			logger.Error("Error inserting CampaignEntity: " + err.Error())
			return err
		}
	}
	return nil
}

func SetCampaign(campaignData campaign.Campaign, userid string) (campaign.Campaign, error) {
	if campaignExists(campaignData.Name, userid) {
		_, err := asset_utils.DBPool.Exec(context.Background(), "UPDATE public.campaigns SET description = $1 AND lastModified = $2 WHERE campaign = $3 AND username = $4", campaignData.Description, utils.FormatDate(campaignData.LastModified), campaignData.Name, userid)
		if err != nil {
			logger.Error("Error updating Campaign: " + err.Error())
			return campaign.Campaign{}, err
		}
	} else {
		_, err := asset_utils.DBPool.Exec(context.Background(), "INSERT INTO public.campaigns (campaign, username, description, creationDate, lastModified) VALUES ($1, $2, $3, $4, $5)", campaignData.Name, userid, campaignData.Description, utils.FormatDate(campaignData.CreationDate), utils.FormatDate(campaignData.LastModified))
		if err != nil {
			logger.Error("Error inserting Campaign: " + err.Error())
			return campaign.Campaign{}, err
		}
	}

	err := setCampaignEntities(campaignData.Players, campaignData.Name, userid)
	if err != nil {
		logger.Error("Error setting Campaign Entities: " + err.Error())
		return campaign.Campaign{}, err
	}

	return campaignData, nil
}

func DeleteCampaign(campaignName string, userid string) error {
	if !campaignExists(campaignName, userid) {
		logger.Error("Campaign does not exist: " + campaignName)
		return errors.New("Campaign does not exist: " + campaignName)
	}

	_, err := asset_utils.DBPool.Exec(context.Background(), "DELETE FROM public.campaigns WHERE campaign = $1 AND username = $2", campaignName, userid)
	if err != nil {
		logger.Error("Error deleting Campaign: " + err.Error())
		return err
	}

	return nil
}
