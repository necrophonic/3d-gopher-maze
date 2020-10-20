package game

import (
	"errors"

	"github.com/necrophonic/gopher-maze/internal/debug"
)

type grid [][]space

// Maze is a fully built maze: Maze[y][x]
type Maze struct {
	grid   grid
	panels map[windowColumn]map[displayType]windowSlice
	scale  string
	height int
	width  int
}

func (m *Maze) getSpace(p point) space {
	return m.grid[p.y][p.x]
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
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
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
		{'X', ' ', 'X', 'X', ' ', ' ', ' ', ' ', 'X'},
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
		{'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', 'X', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', 'X', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'p', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X'},
		{'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X'},
		{'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', ' ', 'X'},
		{'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X', ' ', ' ', ' ', 'X'},
		{'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', 'X', 'X', ' ', 'X', 'X', 'X'},
		{'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', 'X', ' ', 'X', ' ', 'X'},
		{'X', ' ', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X', 'X', 'X', 'X', 'X', 'X', 'X', ' ', 'X', ' ', 'X', ' ', 'X'},
		{'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'X', ' ', ' ', ' ', ' ', ' ', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
	},
}
