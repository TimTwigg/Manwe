package cmapaign

import (
	player "github.com/TimTwigg/Manwe/types/player"
)

type Campaign struct {
	Name        string
	Description string
	Players     []player.Player
}

type CampaignOverview struct {
	Name        string
	Description string
}
