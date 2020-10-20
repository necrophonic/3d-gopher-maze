package game

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

// Slice of multi-dim slices
//
// 'X' = wall
// ' ' = space
// 'p' = player starting position (always starts facing north)
//
type mazeDefinition [][]uint8

var mazes = []mazeDefinition{
	{
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
}
