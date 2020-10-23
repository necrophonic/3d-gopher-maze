package game

import (
	"errors"
	"math/rand"
	"time"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/necrophonic/gopher-maze/internal/game/element"
)

type grid [][]space

// Maze is a fully built maze: Maze[y][x]
type Maze struct {
	grid   grid
	panels map[int]map[string]element.PixelMatrix
	scale  string
	height int
	width  int
}

func (m *Maze) getSpace(p Point) space {
	return m.grid[p.Y][p.X]
}

func (g *Game) importMaze(m mazeDefinition) error {

	height := len(m)
	width := len(m[0])

	debug.Printf("Importing maze: w[%d] h[%d]\n", width, height)

	playerFound := false
	gopherFound := false

	playerStarts := make([]Point, 0, 1)
	gopherStarts := make([]Point, 0, 1)

	newMaze := make([][]space, height)
	for y := range m {
		newMaze[y] = make([]space, width)
		for x, sp := range m[y] {
			// If this is a player start point then we want to
			// mark that point, and set the space as "empty"
			switch sp {
			case SpacePlayerStart:
				playerStarts = append(playerStarts, NewPointInt(x, y))
				sp = uint8(SpaceEmpty)
				playerFound = true
				debug.Printf("Found player start point at (%d,%d)", x, y)
			case SpaceGopherStart:
				g.gopher.p = NewPointInt(x, y)
				gopherStarts = append(gopherStarts, NewPointInt(x, y))
				sp = uint8(SpaceEmpty)
				gopherFound = true
				g.items = append(g.items, g.gopher)
				debug.Printf("Found gopher start point at (%d,%d)", x, y)
			}
			newMaze[y][x] = space{
				t: spaceType(sp),
			}
		}
	}

	if !playerFound {
		return errors.New("bad maze definition: no player start point")
	}
	if !gopherFound {
		return errors.New("bad maze definition: no gopher start point")
	}

	rand.Seed(time.Now().Unix())

	// Choose random player and gopher starts
	g.player.p = playerStarts[rand.Intn(len(playerStarts))]
	debug.Printf("Starting random player start at (%v)", g.player.p)

	g.gopher.p = gopherStarts[rand.Intn(len(gopherStarts))]
	debug.Printf("Starting random gopher start at (%v)", g.gopher.p)

	g.m.grid = newMaze
	return nil
}

// Slice of multi-dim slices
//
// 'X' = wall
// ' ' = space
// 'p' = player starting position (always starts facing north)
//
type mazeDefinition [][]uint8

var mazes = []mazeDefinition{
	{
		// Test maze
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', ' ', 'g', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', ' ', 'X', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', 'p', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', ' ', 'X', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	},
	{
		// Simple maze
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'g', 'X', 'X', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', 'X', 'X', ' ', 'X'},
		{'X', 'X', 'X', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', 'X', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', ' ', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', 'X', 'X', 'X', ' ', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', 'X', 'p', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	},
	{
		// Big maze
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X'},
		{'X', ' ', 'X', 'p', ' ', ' ', 'X', ' ', 'X', 'g', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', 'X', 'g', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', 'g', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', 'X', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'p', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', 'g', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X'},
		{'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X'},
		{'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', 'p', 'X'},
		{'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X'},
		{'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', 'p', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	},
}
