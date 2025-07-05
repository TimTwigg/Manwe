package cmapaign

import (
	"time"

	player "github.com/TimTwigg/Manwe/types/player"
)

type Campaign struct {
	Name         string
	Description  string
	CreationDate time.Time
	LastModified time.Time
	Players      []player.Player
}

type CampaignOverview struct {
	Name         string
	Description  string
	CreationDate time.Time
	LastModified time.Time
}
