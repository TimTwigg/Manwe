package assets

import (
	"context"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	logger "github.com/TimTwigg/Manwe/utils/log"
	pgx "github.com/jackc/pgx/v5"
	errors "github.com/pkg/errors"
)

func ReadAllUsedSources(userid string) ([]string, error) {
	rows, _ := asset_utils.DBPool.Query(context.Background(), "SELECT DISTINCT Source FROM public.statblock WHERE (username = 'public' OR username = $1 OR published = true)", userid)
	sources, err := pgx.CollectRows(rows, pgx.RowTo[string])
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error reading used sources: " + err.Error())
		return nil, errors.Wrap(err, "Error reading used sources from database")
	}
	return sources, nil
}
