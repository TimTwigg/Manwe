package encounters

import "time"

type EncounterMetadata struct {
	CreationDate time.Time
	AccessedDate time.Time
	Campaign     string
	Started      bool
	Round        int
	Turn         int
}

func (m EncounterMetadata) Dict() map[string]any {
	return map[string]any{
		"data_type":    "EncounterMetadata",
		"CreationDate": m.CreationDate,
		"AccessedDate": m.AccessedDate,
		"Campaign":     m.Campaign,
		"Started":      m.Started,
		"Round":        m.Round,
		"Turn":         m.Turn,
	}
}
