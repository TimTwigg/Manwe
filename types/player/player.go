package player

import (
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
)

type Player struct {
	Campaign  string
	StatBlock stat_blocks.StatBlock
	Notes     string
	RowID     int
}

func (p Player) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Campaign":  p.Campaign,
		"StatBlock": p.StatBlock.Dict(),
		"Notes":     p.Notes,
		"RowID":     p.RowID,
	}
}
