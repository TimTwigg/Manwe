package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	campaign "github.com/TimTwigg/Manwe/types/campaign"
	player "github.com/TimTwigg/Manwe/types/player"
	logger "github.com/TimTwigg/Manwe/utils/log"
	errors "github.com/pkg/errors"
)

func ReadCampaign(campaignName string, userid string) (campaign.Campaign, error) {
	rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT campaign, description FROM public.campaigns WHERE campaign = $1 AND username = $2", campaignName, userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return campaign.Campaign{}, err
	}
	defer rows.Close()

	var camp campaign.Campaign

	camp.Players = make([]player.Player, 0)
	if rows.Next() {
		if err = rows.Scan(&camp.Name, &camp.Description); err != nil {
			logger.Error("Error scanning row: " + err.Error())
			return campaign.Campaign{}, err
		}
	} else {
		logger.Error("No campaign found with name: " + campaignName)
		return campaign.Campaign{}, errors.New("No Campaign found with name: " + campaignName)
	}

	entity_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT statblockid, notes FROM public.campaignentities WHERE campaign = $1 AND username = $2", campaignName, userid)
	if err != nil {
		logger.Error("Error querying CampaignEntities: " + err.Error())
		return campaign.Campaign{}, err
	}
	defer entity_rows.Close()
	for entity_rows.Next() {
		var statblockID int
		var notes string
		if err = entity_rows.Scan(&statblockID, &notes); err != nil {
			logger.Error("Error scanning CampaignEntities row: " + err.Error())
			return campaign.Campaign{}, err
		}
		statblock, err := ReadStatBlockByID(statblockID, userid, asset_utils.PLAYER)
		if err != nil {
			logger.Error("Error reading StatBlock by ID: " + err.Error())
			return campaign.Campaign{}, err
		}
		p := player.Player{
			StatBlock: statblock,
			Notes:     notes,
		}
		camp.Players = append(camp.Players, p)
	}

	return camp, nil
}

func ReadAllCampaignOverviews(userid string) ([]campaign.CampaignOverview, error) {
	rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT campaign, description, creationDate, lastModified FROM public.campaigns WHERE username = $1", userid)
	if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var campaigns []campaign.CampaignOverview = make([]campaign.CampaignOverview, 0)

	for rows.Next() {
		var camp campaign.CampaignOverview
		if err = rows.Scan(&camp.Name, &camp.Description, &camp.CreationDate, &camp.LastModified); err != nil {
			logger.Error("Error scanning row: " + err.Error())
			return nil, err
		}
		campaigns = append(campaigns, camp)
	}

	return campaigns, nil
}
