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
			return "AND RecordType = 'Statblock'"
		} else {
			return "WHERE RecordType = 'Statblock'"
		}
	case PLAYER:
		if useAnd {
			return "AND RecordType = 'Player'"
		} else {
			return "WHERE RecordType = 'Player'"
		}
	case CUSTOM:
		if useAnd {
			return "AND RecordType = 'Custom'"
		} else {
			return "WHERE RecordType = 'Custom'"
		}
	default:
		if useAnd {
			return "AND RecordType = 'ERROR'"
		} else {
			return "WHERE RecordType = 'ERROR'"
		}
	}
}
