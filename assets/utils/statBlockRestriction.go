package asset_utils

type EntityTypeRestriction int

const (
	ANY EntityTypeRestriction = iota
	STATBLOCK
	PLAYER
	CUSTOM
)

func StatBlockRestrictionClause(restriction EntityTypeRestriction, useAnd bool) string {
	switch restriction {
	case ANY:
		return ""
	case STATBLOCK:
		if useAnd {
			return " AND recordtype = 'Statblock'"
		} else {
			return " WHERE recordtype = 'Statblock'"
		}
	case PLAYER:
		if useAnd {
			return " AND recordtype = 'Player'"
		} else {
			return " WHERE recordtype = 'Player'"
		}
	case CUSTOM:
		if useAnd {
			return " AND recordtype = 'Custom'"
		} else {
			return " WHERE recordtype = 'Custom'"
		}
	default:
		if useAnd {
			return " AND recordtype = 'ERROR'"
		} else {
			return " WHERE recordtype = 'ERROR'"
		}
	}
}
