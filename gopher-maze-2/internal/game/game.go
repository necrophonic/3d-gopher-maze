package game

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/developer"
	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/terminal"
)

var ds developer.Stack

func init() {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		ds = developer.NewDebugStack(11)
		ds.Add("[debug enabled]")
	} else {
		ds = &developer.NullStack{}
	}
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
func New() *Game {
	return &Game{
		Tick:   100 * time.Millisecond,
		Player: NewPlayer(),
	}
}

// Loop is the main game loop. It is designed
// to be called as a goroutine
func (g *Game) Loop(ctx context.Context) {
	ds.Add("Game loop started")
	ticker := time.NewTicker(g.Tick)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("Main game loop closing")
			return
		case <-ticker.C:
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
		}
	}
}
