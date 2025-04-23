package asset_utils

import (
	"os"
)

// Check if a stat block exists in the database
func StatBlockExists(name string) bool {
	_, err := os.Stat("assets\\stat_blocks\\" + name + ".json")
	return err == nil
}
