package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	condition "github.com/TimTwigg/Manwe/types/conditions"
	logger "github.com/TimTwigg/Manwe/utils/log"
	pgx "github.com/jackc/pgx/v5"
)

func ReadAllConditions(userid string) ([]condition.Condition, error) {
	rows, _ := asset_utils.DBPool.Query(context.Background(), "SELECT condition as Name FROM public.condition WHERE (username = 'public' OR username = $1 OR published = true)", userid)
	conditions, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (condition.Condition, error) {
		var cond condition.Condition
		if err := row.Scan(&cond.Name); err != nil {
			return condition.Condition{}, err
		}
		effectRows, _ := asset_utils.DBPool.Query(context.Background(), "SELECT description FROM public.conditioneffect WHERE condition = $1", cond.Name)
		effects, err := pgx.CollectRows(effectRows, pgx.RowTo[string])
		if err != nil {
			logger.Error("Error reading effects for condition " + cond.Name + ": " + err.Error())
			return condition.Condition{}, err
		}
		cond.Effects = effects
		return cond, nil
	})
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading conditions from database" + err.Error())
		return nil, err
	}
	return conditions, nil
}
