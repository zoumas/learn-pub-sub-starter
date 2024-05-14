package game

import (
	"sync"
)

type GameState struct {
	mu     *sync.RWMutex
	Player Player
	Paused bool
}

func NewGameState(username string) *GameState {
	return &GameState{
		Player: Player{
			Username: username,
			Units:    make(map[int]Unit),
		},
		Paused: false,
		mu:     &sync.RWMutex{},
	}
}

func (gs *GameState) resume() {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Paused = false
}

func (gs *GameState) pause() {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Paused = true
}

func (gs *GameState) isPaused() bool {
	gs.mu.RLock()
	defer gs.mu.RUnlock()
	return gs.Paused
}

func (gs *GameState) removeUnitsInLocation(l Location) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	for k, v := range gs.Player.Units {
		if v.Location == l {
			delete(gs.Player.Units, k)
		}
	}
}

func (gs *GameState) addUnit(u Unit) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Player.Units[u.ID] = u
}

func (gs *GameState) UpdateUnit(u Unit) {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Player.Units[u.ID] = u
}

func (gs *GameState) GetUsername() string {
	return gs.Player.Username
}

func (gs *GameState) getUnitsSnapshot() []Unit {
	gs.mu.RLock()
	defer gs.mu.RUnlock()
	units := []Unit{}
	for _, v := range gs.Player.Units {
		units = append(units, v)
	}
	return units
}

func (gs *GameState) GetUnit(id int) (Unit, bool) {
	gs.mu.RLock()
	defer gs.mu.RUnlock()
	u, ok := gs.Player.Units[id]
	return u, ok
}

func (gs *GameState) GetPlayerSnapshot() Player {
	gs.mu.RLock()
	defer gs.mu.RUnlock()
	units := map[int]Unit{}
	for k, v := range gs.Player.Units {
		units[k] = v
	}
	return Player{
		Username: gs.Player.Username,
		Units:    units,
	}
}
