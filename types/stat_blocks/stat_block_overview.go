package stat_blocks

type StatBlockOverview struct {
	ID              string
	Name            string
	Type            string
	Size            string
	ChallengeRating int
	Source          string
}

func (sbo StatBlockOverview) Dict() map[string]any {
	return map[string]any{
		"ID":              sbo.ID,
		"Name":            sbo.Name,
		"Type":            sbo.Type,
		"Size":            sbo.Size,
		"ChallengeRating": sbo.ChallengeRating,
		"Source":          sbo.Source,
	}
}
