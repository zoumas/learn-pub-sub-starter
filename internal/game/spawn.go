package game

import (
	"errors"
	"fmt"
)

func (gs *GameState) CommandSpawn(words []string) error {
	if len(words) < 3 {
		return errors.New("usage: spawn <location> <rank>")
	}

	location := Location(words[1])
	if !validLocation(location) {
		return fmt.Errorf("error: %s is not a valid location", location)
	}

	rank := UnitRank(words[2])
	if !validRank(rank) {
		return fmt.Errorf("error: %s is not a valid unit", rank)
	}

	id := len(gs.getUnitsSnapshot()) + 1
	gs.addUnit(Unit{
		ID:       id,
		Rank:     rank,
		Location: location,
	})

	fmt.Printf("Spawned a(n) %s in %s with id %v\n", rank, location, id)
	return nil
}

func validLocation(l Location) bool {
	_, ok := getAllLocations()[l]
	return ok
}

func validRank(r UnitRank) bool {
	_, ok := getAllRanks()[r]
	return ok
}
