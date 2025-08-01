package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	campaign "github.com/TimTwigg/Manwe/types/campaign"
	player "github.com/TimTwigg/Manwe/types/player"
	logger "github.com/TimTwigg/Manwe/utils/log"
	pgx "github.com/jackc/pgx/v5"
	errors "github.com/pkg/errors"
)

func ReadCampaign(campaignName string, userid string) (campaign.Campaign, error) {
	var camp campaign.Campaign
	camp.Players = make([]player.Player, 0)
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT campaign, description, creationdate, lastmodified FROM public.campaigns WHERE campaign = $1 AND username = $2", campaignName, userid).Scan(&camp.Name, &camp.Description, &camp.CreationDate, &camp.LastModified)
	if pgx.ErrNoRows == err {
		logger.Error("No Campaign found with name: " + campaignName)
		return campaign.Campaign{}, errors.New("No Campaign found with name: " + campaignName)
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return campaign.Campaign{}, errors.Wrap(err, "Error querying database for campaign")
	}

	player_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT statblockid, notes FROM public.campaignentities WHERE campaign = $1 AND username = $2", campaignName, userid)
	players, err := pgx.CollectRows(player_rows, func(row pgx.CollectableRow) (player.Player, error) {
		var p player.Player
		var statblockID int
		if err := row.Scan(&statblockID, &p.Notes); err != nil {
			logger.Error("Error scanning CampaignEntities row: " + err.Error())
			return player.Player{}, errors.Wrap(err, "Error scanning CampaignEntities row")
		}
		statblock, err := ReadStatBlockByID(statblockID, userid, asset_utils.PLAYER)
		if err != nil {
			logger.Error("Error reading StatBlock by ID: " + err.Error())
			return player.Player{}, errors.Wrap(err, "Error reading StatBlock by ID")
		}
		p.StatBlock = statblock
		return p, nil
	})
	err = player_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading players from database: " + err.Error())
		return campaign.Campaign{}, errors.Wrap(err, "Error reading players from database")
	}
	camp.Players = players
	return camp, nil
}

func ReadAllCampaignOverviews(userid string) ([]campaign.CampaignOverview, error) {
	rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT campaign, description, creationDate, lastModified FROM public.campaigns WHERE username = $1", userid)
	campaigns, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (campaign.CampaignOverview, error) {
		var camp campaign.CampaignOverview
		if err := row.Scan(&camp.Name, &camp.Description, &camp.CreationDate, &camp.LastModified); err != nil {
			logger.Error("Error scanning CampaignOverview row: " + err.Error())
			return campaign.CampaignOverview{}, errors.Wrap(err, "Error scanning CampaignOverview row")
		}
		return camp, nil
	})
	err = rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading campaigns from database: " + err.Error())
		return nil, errors.Wrap(err, "Error reading campaigns from database")
	}
	return campaigns, nil
}
