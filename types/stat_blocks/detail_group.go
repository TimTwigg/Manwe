package stat_blocks

type DetailBlock struct {
	ArmorType    string
	Skills       []Stat
	SavingThrows []Stat
	Senses       []Sense
	Languages    []string
	Traits       []Trait
}
