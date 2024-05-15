package game

import (
	"fmt"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
)

func (gs *GameState) HandlePause(ps routing.PlayingState) {
	defer fmt.Println("------------------------")
	fmt.Println()
	if ps.IsPaused {
		fmt.Println("==== Pause Detected ====")
		gs.pause()
	} else {
		fmt.Println("==== Resume Detected ====")
		gs.resume()
	}
}
