package game

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
}
