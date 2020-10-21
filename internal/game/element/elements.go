package element

// Shorthand constants defining screen pixels
const (
	W uint8 = iota // wall
	O              // open wall
	T              // transparent
)

// PixelMatrix defines a 2 dimensional array of pixels. It
// is indexed by row (y) then column (x): [y][x]uint8
type PixelMatrix [][]uint8

// Height returns the column height of the pixel matrix
func (pm PixelMatrix) Height() int {
	return len(pm)
}

// Width returns the row length of the pixel matrix
func (pm PixelMatrix) Width() int {
	return len(pm[0])
}
