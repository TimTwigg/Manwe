package stat_blocks

type Speeds struct {
	Walk   int
	Fly    int
	Swim   int
	Climb  int
	Burrow int
}

type HitPointsT struct {
	Average int
	Dice    string
}

type NumericalAttributes struct {
	ArmorClass int
	HitPoints  HitPointsT
	Speed      Speeds
	Abilities  map[string]int
}

func (s Speeds) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Walk":   s.Walk,
		"Fly":    s.Fly,
		"Swim":   s.Swim,
		"Climb":  s.Climb,
		"Burrow": s.Burrow,
	}
}

func (h HitPointsT) Dict() map[string]interface{} {
	return map[string]interface{}{
		"Average": h.Average,
		"Dice":    h.Dice,
	}
}

func (n NumericalAttributes) Dict() map[string]interface{} {
	return map[string]interface{}{
		"ArmorClass": n.ArmorClass,
		"HitPoints":  n.HitPoints,
		"Speed":      n.Speed,
		"Abilities":  n.Abilities,
	}
}
