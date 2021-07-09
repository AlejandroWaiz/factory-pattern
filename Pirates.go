package main

type MonkeyDLuffy struct {
	Pirate
}

func NewMonkeyDLuffy() IPirate {
	return &MonkeyDLuffy{
		Pirate: Pirate{
			name:  "Monkey D. Luffy",
			power: "Gomu Gomu No Mi",
		},
	}
}

type TrafalgarDWaterLaw struct {
	Pirate
}

func NewTrafalgarDWaterLaw() IPirate {
	return &TrafalgarDWaterLaw{
		Pirate: Pirate{
			name:  "Trafalgar D. Water Law",
			power: "Ope Ope No Mi",
		},
	}
}

type GolDRoger struct {
	Pirate
}

func NewGolDRoger() IPirate {
	return &GolDRoger{
		Pirate: Pirate{
			name:  "Gol D. Roger",
			power: "Haoshoku No Haki",
		},
	}
}
