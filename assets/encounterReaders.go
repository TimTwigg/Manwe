package assets

import (
	"context"
	"strconv"
	"strings"

	asset_utils "github.com/TimTwigg/Manwe/assets/utils"
	encounters "github.com/TimTwigg/Manwe/types/encounters"
	entities "github.com/TimTwigg/Manwe/types/entities"
	generics "github.com/TimTwigg/Manwe/types/generics"
	error_utils "github.com/TimTwigg/Manwe/utils/errors"
	logger "github.com/TimTwigg/Manwe/utils/log"
	pgx "github.com/jackc/pgx/v5"
	errors "github.com/pkg/errors"
)

func ReadEncounterByID(id int, userid string) (encounters.Encounter, error) {
	var encounter encounters.Encounter
	encounter.Entities = make([]entities.Entity, 0)
	encounter.Lair.Actions.Items = make([]generics.SimpleItem, 0)
	encounter.Lair.RegionalEffects.Items = make([]generics.SimpleItem, 0)
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT encounterid, name, description, creationdate, accesseddate, campaign, started, round, turn, haslair, lairownerid, activeid FROM public.encounter WHERE encounterid = $1 AND (username = 'public' OR username = $2)", id, userid).Scan(
		&encounter.ID,
		&encounter.Name,
		&encounter.Description,
		&encounter.Metadata.CreationDate,
		&encounter.Metadata.AccessedDate,
		&encounter.Metadata.Campaign,
		&encounter.Metadata.Started,
		&encounter.Metadata.Round,
		&encounter.Metadata.Turn,
		&encounter.HasLair,
		&encounter.LairOwnerID,
		&encounter.ActiveID,
	)
	if pgx.ErrNoRows == err {
		logger.Error("No Encounter found with id: " + strconv.Itoa(id))
		return encounters.Encounter{}, errors.New("No Encounter found with id: " + strconv.Itoa(id))
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}

	entity_rows, _ := asset_utils.DBPool.Query(context.Background(), "SELECT rowid, statblockid, suffix, initiative, maxhitpoints, temphitpoints, currenthitpoints, armorclassbonus, concentration, notes, ishostile, encounterlocked, id FROM public.encounterentities WHERE encounterid = $1", id)
	entities, err := pgx.CollectRows(entity_rows, func(row pgx.CollectableRow) (entities.Entity, error) {
		var entity entities.Entity
		if err := row.Scan(
			&entity.DBID,
			&entity.ID,
			&entity.Suffix,
			&entity.Initiative,
			&entity.MaxHitPoints,
			&entity.TempHitPoints,
			&entity.CurrentHitPoints,
			&entity.ArmorClassBonus,
			&entity.Concentration,
			&entity.Notes,
			&entity.IsHostile,
			&entity.EncounterLocked,
			&entity.ID,
		); err != nil {
			logger.Error("Error Scanning Encounter Entity Row: " + err.Error())
			return entities.Entity{}, err
		}
		statblock, err := ReadStatBlockByID(entity.DBID, userid, asset_utils.ANY)
		if err != nil {
			logger.Error("Error reading statblock: " + err.Error())
			return entities.Entity{}, err
		}
		entity.Displayable = statblock

		conditions_rows, err := asset_utils.DBPool.Query(context.Background(), "SELECT condition, duration FROM public.encentconditions WHERE encounterid = $1 and rowid = $2", id, entity.DBID)
		if err != nil {
			logger.Error("Error querying database: " + err.Error())
			return entities.Entity{}, err
		}
		entity.Conditions = make(map[string]int, 0)
		var condition string
		var duration int
		pgx.ForEachRow(conditions_rows, []any{&condition, &duration}, func() error {
			entity.Conditions[condition] = duration
			return nil
		})

		return entity, nil
	})
	err = entity_rows.Err()
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}
	encounter.Entities = entities

	if encounter.HasLair && encounter.LairOwnerID > 0 {
		if encounter.Lair, err = ReadLairByEntityID(encounter.LairOwnerID); err != nil {
			if !strings.HasPrefix(err.Error(), "No Lair found") {
				logger.Error("Error reading lair: " + err.Error())
				return encounters.Encounter{}, error_utils.ParseError{Message: err.Error()}
			}
		}
	}

	return encounter, nil
}

func ReadEncounterByName(name string, userid string) (encounters.Encounter, error) {
	var id int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT encounterid FROM public.encounter WHERE name = $1 AND (username = 'public' OR username = $2)", name, userid).Scan(&id)
	if pgx.ErrNoRows == err {
		logger.Error("No Encounter found with name: " + name)
		return encounters.Encounter{}, errors.New("No Encounter found with name: " + name)
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.Encounter{}, err
	}

	return ReadEncounterByID(id, userid)
}

func ReadEncounterOverviewByID(id int, userid string) (encounters.EncounterOverview, error) {
	var encounter encounters.EncounterOverview
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT name, description, creationdate, accesseddate, campaign, started, round, turn FROM public.encounter WHERE encounterid = $1 AND (username = 'public' OR username = $2)", id, userid).Scan(
		&encounter.Name,
		&encounter.Description,
		&encounter.Metadata.CreationDate,
		&encounter.Metadata.AccessedDate,
		&encounter.Metadata.Campaign,
		&encounter.Metadata.Started,
		&encounter.Metadata.Round,
		&encounter.Metadata.Turn,
	)
	if pgx.ErrNoRows == err {
		logger.Error("No Encounter found with id: " + strconv.Itoa(id))
		return encounters.EncounterOverview{}, errors.New("No Encounter found with id: " + strconv.Itoa(id))
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.EncounterOverview{}, err
	}
	return encounter, nil
}

func ReadEncounterOverviewByName(name string, userid string) (encounters.EncounterOverview, error) {
	var id int
	err := asset_utils.DBPool.QueryRow(context.Background(), "SELECT encounterid FROM public.encounter WHERE name = $1 AND (username = 'public' OR username = $2)", name, userid).Scan(&id)
	if pgx.ErrNoRows == err {
		logger.Error("No Encounter found with name: " + name)
		return encounters.EncounterOverview{}, errors.New("No Encounter found with name: " + name)
	} else if err != nil {
		logger.Error("Error querying database: " + err.Error())
		return encounters.EncounterOverview{}, err
	}
	return ReadEncounterOverviewByID(id, userid)
}

func ReadAllEncounterOverviews(userid string) ([]encounters.EncounterOverview, error) {
	_rows, _ := asset_utils.DBPool.Query(context.Background(), "SELECT encounterid, name, description, creationdate, accesseddate, campaign, started, round, turn FROM public.encounter WHERE (username = 'public' OR username = $1)", userid)
	rows, err := pgx.CollectRows(_rows, func(row pgx.CollectableRow) (encounters.EncounterOverview, error) {
		var encounter encounters.EncounterOverview
		if err := row.Scan(
			&encounter.ID,
			&encounter.Name,
			&encounter.Description,
			&encounter.Metadata.CreationDate,
			&encounter.Metadata.AccessedDate,
			&encounter.Metadata.Campaign,
			&encounter.Metadata.Started,
			&encounter.Metadata.Round,
			&encounter.Metadata.Turn,
		); err != nil {
			logger.Error("Error Scanning Encounter Row: " + err.Error())
			return encounters.EncounterOverview{}, err
		}
		return encounter, nil
	})
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("Error querying database: " + err.Error())
		return nil, err
	}

	return rows, nil
}

func ReadEncounterByAccessType(accessType string, accessor string, userid string) (encounters.Encounter, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return encounters.Encounter{}, err
		}
		return ReadEncounterByID(id, userid)
	case "name":
		return ReadEncounterByName(accessor, userid)
	default:
		logger.Error("Invalid access type: " + accessType)
		return encounters.Encounter{}, errors.New("Invalid access type: " + accessType)
	}
}

func ReadEncounterOverviewByAccessType(accessType string, accessor string, userid string) (encounters.EncounterOverview, error) {
	switch accessType {
	case "id":
		id, err := strconv.Atoi(accessor)
		if err != nil {
			logger.Error("Error converting id to int: " + err.Error())
			return encounters.EncounterOverview{}, err
		}
		return ReadEncounterOverviewByID(id, userid)
	case "name":
		return ReadEncounterOverviewByName(accessor, userid)
	default:
		logger.Error("Invalid access type: " + accessType)
		return encounters.EncounterOverview{}, errors.New("Invalid access type: " + accessType)
	}
}
