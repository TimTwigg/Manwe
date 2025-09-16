package player

import (
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
)

type Player struct {
	CampaignID int
	StatBlock  stat_blocks.StatBlock
	Notes      string
	RowID      int
}

func (p Player) Dict() map[string]interface{} {
	return map[string]interface{}{
		"CampaignID": p.CampaignID,
		"StatBlock":  p.StatBlock.Dict(),
		"Notes":      p.Notes,
		"RowID":      p.RowID,
	}
}
