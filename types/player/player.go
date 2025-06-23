package player

import (
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
)

type Player struct {
	StatBlock stat_blocks.StatBlock
	Notes     string
}
