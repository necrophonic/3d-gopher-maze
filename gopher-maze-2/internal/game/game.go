package game

import (
	"fmt"
	"os"
	"time"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/developer"
	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/terminal"
)

// Point is a point on the game grid denoted by its x, y co-ordinate. The 0,0
// point is top left.
type Point struct {
	X, Y int
}

type (
	// Game represents the overall current running game state
	Game struct {
		Tick       time.Duration
		debug      bool
		debugStack *developer.DebugStack
	}
)

// New initialises a new game with defaults
func New(debug bool) *Game {
	gme := &Game{
		Tick: 100_000 * time.Microsecond,
	}
	if debug {
		// Enable the debugging stack
		// if debug was set to true.
		gme.debug = true
		gme.debugStack = developer.NewDebugStack(9)
		gme.Debug("[debug enabled]")
	}
	return gme
}

// Debug adds a message to the current debugging
// stack (if one has been initialised)
func (g *Game) Debug(msg string) {
	if !g.debug {
		return
	}
	g.debugStack.AddMsg(msg)
}

// Loop is the main game loop. It is designed
// to be called as a goroutine
func (g *Game) Loop() {
	g.Debug("Game loop started")
	for {
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
		time.Sleep(g.Tick)
	}
}
