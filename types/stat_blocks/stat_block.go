package stat_blocks

type StatBlock struct {
	Name            string
	ChallengeRating int
	Description     EntityDescription
	Stats           NumericalAttributes
}
