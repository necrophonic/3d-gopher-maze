package game

import (
	"fmt"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/pkg/errors"
)

type (
	// Player represents the current player state
	Player struct {
		x uint
		y uint
		o byte
	}

	// Game represents the current game state
	Game struct {
		p *Player
		m *Maze
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

	// Maze is a fully built maze
	Maze [][]space
)

var walls = [2]rune{'░', '▓'}

// New creates a new game state
func New() *Game {
	return &Game{
		p: &Player{
			o: 'n',
		},
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

	g.render()
	return nil
}

func (g *Game) importMaze(m mazeDefinition) error {

	height := len(m)
	width := len(m[0])

	debug.Printf("Importing maze: w[%d] h[%d]\n", width, height)

	playerFound := false

	newMaze := make(Maze, height)
	for y := range m {
		newMaze[y] = make([]space, width)
		for x, sp := range m[y] {
			// If this is a player start point then we want to
			// mark that point, and set the space as "empty"
			if sp == SpacePlayerStart {
				g.p.x = uint(x)
				g.p.y = uint(y)
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

	g.m = &newMaze
	return nil
}

func (g *Game) render() {

}
