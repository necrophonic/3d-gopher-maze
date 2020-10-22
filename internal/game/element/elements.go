package element

// Shorthand constants defining screen pixels
const (
	W uint8 = iota // wall
	O              // open wall
	T              // transparent
	X              // default background for unrendered scenes
	G              // gopher body
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

// Clear fills a PixelMatrix with valid, but empty pixels
func (pm *PixelMatrix) Clear() {
	for y := 0; y < len(*pm); y++ {
		for x := 0; x < len((*pm)[y]); x++ {
			(*pm)[y][x] = X
		}
	}
}
