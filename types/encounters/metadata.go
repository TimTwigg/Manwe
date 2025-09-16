package encounters

import "time"

type EncounterMetadata struct {
	CreationDate time.Time
	AccessedDate time.Time
	CampaignID   int
	Started      bool
	Round        int
	Turn         int
}

func (m EncounterMetadata) Dict() map[string]any {
	return map[string]any{
		"CreationDate": m.CreationDate,
		"AccessedDate": m.AccessedDate,
		"CampaignID":   m.CampaignID,
		"Started":      m.Started,
		"Round":        m.Round,
		"Turn":         m.Turn,
	}
}
