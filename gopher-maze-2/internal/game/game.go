package game

import (
	"fmt"
	"os"
	"time"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/developer"
	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/terminal"
)

var ds *developer.DebugStack

func init() {
	// TODO make debug switchable via env vars
	ds = developer.NewDebugStack(true, 11)
	ds.Add("[debug enabled]")
}

type (
	// Point is a point on the game grid denoted by
	// its x, y co-ordinate. The 0,0 point is top left.
	Point struct {
		X, Y int
	}

	// Game represents the overall current running game state
	Game struct {
		// Tick is the interval between engine refreshes
		Tick time.Duration

		// Player represents the current player state
		Player *Player
	}
)

// New initialises a new game with defaults
func New(debug bool) *Game {
	return &Game{
		Tick:   100_000 * time.Microsecond,
		Player: NewPlayer(),
	}
}

// Loop is the main game loop. It is designed
// to be called as a goroutine
func (g *Game) Loop() {
	ds.Add("Game loop started")
	ticker := time.NewTicker(g.Tick)
	defer ticker.Stop()
	for range ticker.C {
		// Refresh the terminal before
		// calling render to build the view.
		terminal.Clear()
		rendered, err := g.Render()
		if err != nil {
			// TODO better error handling
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(rendered)
		// time.Sleep(g.Tick)
	}
}
