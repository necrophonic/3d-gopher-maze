package game

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/pkg/errors"
)

type (
	// Player represents the current player state
	Player struct {
		x int8
		y int8
		o byte
	}

	moveVector struct {
		x int8
		y int8
	}

	// Game represents the current game state
	Game struct {
		p    *Player
		m    *Maze
		v    *view
		move moveVector
	}
)

// Constants for types of maze space
const (
	SpaceEmpty       spaceType = ' '
	SpaceWall                  = 'X'
	SpacePlayerStart           = 'p'
)

type (
	spaceType uint8

	space struct {
		t spaceType
	}
)

// New creates a new game state
func New() *Game {
	return &Game{
		p: &Player{
			o: 'n',
		},
		m: &Maze{
			grid:   grid{},
			panels: nil,
		},
		v:    &view{},
		move: moveVector{0, -1},
	}
}

// Swatch returns a string with the defined game elements
func Swatch() string {
	return fmt.Sprintf(
		"%-10s: %c\n%-10s: %c\n",
		"Wall #1", walls[0],
		"Wall #2", walls[1],
	)
}

// Run performs the main game loop
func (g *Game) Run() error {
	// TODO split out scaler
	g.m.scale = "1"
	if s := os.Getenv("SCALE"); s != "" {
		g.m.scale = s
	}
	if err := g.setUpScaledPanels(); err != nil {
		return errors.WithMessage(err, "failed to set up scaling")
	}

	// TODO randomly (totally or from set of criteria) select a maze
	// TODO would be nice to able to dynamically create one!
	if err := g.importMaze(mazes[2]); err != nil {
		return errors.WithMessage(err, "failed to import maze")
	}

	for {
		if !debug.Debug {
			// TODO compile for windows
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

		if err := g.updateView(); err != nil {
			return errors.WithMessage(err, "failed to update view")
		}
		fmt.Print(g.render())

		reader := bufio.NewReader(os.Stdin)

		char, _, err := reader.ReadRune()
		if err != nil {
			return errors.WithMessage(err, "error reading rune from terminal")
		}

		// TODO remove need to hit return!

		switch char {
		case 'w':
			debug.Println("Move forward")
			g.moveForward()
			break
		case 's':
			debug.Println("Move backward")
			g.moveBackwards()
			break
		case 'd':
			debug.Println("Turn right")
			g.rotateRight()
			break
		case 'a':
			debug.Println("Turn left")
			g.rotateLeft()
			break
		case 'q':
			debug.Println("Exiting game")
			fmt.Println("Goodbye!")
			return nil
		}
		debug.Printf("Player is now at (%d,%d). Facing (%c)\n", g.p.x, g.p.y, g.p.o)
	}
}
