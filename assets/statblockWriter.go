package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	stat_blocks "github.com/TimTwigg/Manwe/types/stat_blocks"
	logger "github.com/TimTwigg/Manwe/utils/log"
)

func statblockExists(statblockID int, userid string) bool {
	var count int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT COUNT(*) FROM public.statblock WHERE statblockid = $1 and username = $2", statblockID, userid).Scan(&count)
	if err != nil {
		logger.Error("Error checking StatblockID: " + err.Error())
		return false
	}
	return count > 0
}

func SetStatblock(statblock stat_blocks.StatBlock, userid string) (stat_blocks.StatBlock, error) {
	if statblock.ID == 0 || !statblockExists(statblock.ID, userid) {
		// Insert new StatBlock
		err := asset_utils.DBPool.QueryRow(context.Background(), "INSERT INTO public.statblock (name, challengerating, proficiencybonus, source, size, type, alignment, armorclass, hitpoints1, hitpoints2, walkspeed, flyspeed, climbspeed, swimspeed, burrowspeed, armortype, recordtype, username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) RETURNING statblockid",
			statblock.Name,
			statblock.ChallengeRating,
			statblock.ProficiencyBonus,
			statblock.Source,
			statblock.Description.Size,
			statblock.Description.Type,
			statblock.Description.Alignment,
			statblock.Stats.ArmorClass,
			statblock.Stats.HitPoints.Average,
			statblock.Stats.HitPoints.Dice,
			statblock.Stats.Speed.Walk,
			statblock.Stats.Speed.Fly,
			statblock.Stats.Speed.Climb,
			statblock.Stats.Speed.Swim,
			statblock.Stats.Speed.Burrow,
			statblock.Details.ArmorType,
			statblock.Description.Category,
			userid,
		).Scan(&statblock.ID)
		if err != nil {
			logger.Error("Error inserting StatBlock: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
	} else {
		// Update existing StatBlock
		_, err := asset_utils.DBPool.Exec(
			context.Background(),
			"UPDATE public.statblock SET name = $1, challengerating = $2, proficiencybonus = $3, source = $4, size = $5, type = $6, alignment = $7, armorclass = $8, hitpoints1 = $9, hitpoints2 = $10, walkspeed = $11, flyspeed = $12, climbspeed = $13, swimspeed = $14, burrowspeed = $15, armortype = $16, recordtype = $17 WHERE statblockid = $18 AND username = $19",
			statblock.Name,
			statblock.ChallengeRating,
			statblock.ProficiencyBonus,
			statblock.Source,
			statblock.Description.Size,
			statblock.Description.Type,
			statblock.Description.Alignment,
			statblock.Stats.ArmorClass,
			statblock.Stats.HitPoints.Average,
			statblock.Stats.HitPoints.Dice,
			statblock.Stats.Speed.Walk,
			statblock.Stats.Speed.Fly,
			statblock.Stats.Speed.Climb,
			statblock.Stats.Speed.Swim,
			statblock.Stats.Speed.Burrow,
			statblock.Details.ArmorType,
			statblock.Description.Category,
			statblock.ID,
			userid,
		)
		if err != nil {
			logger.Error("Error updating StatBlock: " + err.Error())
			return stat_blocks.StatBlock{}, err
		}
	}

	return statblock, nil
}
