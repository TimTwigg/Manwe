package stat_blocks

type Speeds struct {
	Walk  int
	Fly   int
	Swim  int
	Climb int
}

type NumericalAttributes struct {
	ArmorClass   int
	HitPoints    int
	Speed        Speeds
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}
