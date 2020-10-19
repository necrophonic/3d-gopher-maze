package game

import (
	"fmt"

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

	// Game represents the current game state
	Game struct {
		p *Player
		m *Maze
		v *view
	}
)

// Constants for types of maze space
const (
	SpaceEmpty       spaceType = ' '
	SpaceWall                  = 'X'
	SpacePlayerStart           = 'p'
)

const displayWidth = 11
const displayHeight = 9

type (
	spaceType uint8

	space struct {
		t spaceType
	}
)

var walls = [2]rune{'░', '▓'}

// New creates a new game state
func New() *Game {
	return &Game{
		p: &Player{
			o: 'n',
		},
		m: &Maze{
			grid: grid{},
		},
		v: &view{},
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
	// TODO randomly (totally or from set of criteria) select a maze
	// TODO would be nice to able to dynamically create one!
	if err := g.importMaze(mazes[0]); err != nil {
		return errors.WithMessage(err, "failed to import maze")
	}

	if err := g.updateView(); err != nil {
		return errors.WithMessage(err, "failed to update view")
	}
	fmt.Println(g.render())
	return nil
}

func (g *Game) importMaze(m mazeDefinition) error {

	height := len(m)
	width := len(m[0])

	debug.Printf("Importing maze: w[%d] h[%d]\n", width, height)

	playerFound := false

	newMaze := make([][]space, height)
	for y := range m {
		newMaze[y] = make([]space, width)
		for x, sp := range m[y] {
			// If this is a player start point then we want to
			// mark that point, and set the space as "empty"
			if sp == SpacePlayerStart {
				g.p.x = int8(x)
				g.p.y = int8(y)
				sp = uint8(SpaceEmpty)
				playerFound = true
				debug.Printf("Found player start point at (%d,%d)", x, y)
			}
			newMaze[y][x] = space{
				t: spaceType(sp),
			}
		}
	}

	if !playerFound {
		return errors.New("bad maze definition: no player start point")
	}

	g.m.grid = newMaze
	return nil
}
