package game

const (
	RankInfantry  = "infantry"
	RankCavalry   = "cavalry"
	RankArtillery = "artillery"

	powerArtillery = 10
	powerCavalry   = 5
	powerInfantry  = 1
)

type Player struct {
	Units    map[int]Unit
	Username string
}

type (
	UnitRank string
	Location string
)

type Unit struct {
	Rank     UnitRank
	Location Location
	ID       int
}

type ArmyMove struct {
	Player     Player
	ToLocation Location
	Units      []Unit
}

type RecognitionOfWar struct {
	Attacker Player
	Defender Player
}

func getAllRanks() map[UnitRank]struct{} {
	return map[UnitRank]struct{}{
		RankInfantry:  {},
		RankCavalry:   {},
		RankArtillery: {},
	}
}

func getAllLocations() map[Location]struct{} {
	return map[Location]struct{}{
		"americas":   {},
		"europe":     {},
		"africa":     {},
		"asia":       {},
		"australia":  {},
		"antarctica": {},
	}
}
