package game

import (
	"bytes"
	"fmt"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/necrophonic/gopher-maze/internal/game/element"
)

type (
	// pixel       rune
	pixelMatrix [][]rune
)

const numPanels = 7
const viewHeight = 9
const viewWidth = 11

type windowSlice [][]rune
type point struct {
	x int8
	y int8
}

// TODO refactor swtches
var walls = [2]rune{'▒', '░'}

// ErrBadSpace is returned if the space definition in a
// point location is unexpected
type ErrBadSpace struct {
	p point
}

func (e ErrBadSpace) Error() string {
	return fmt.Sprintf("bad space definition at (%d,%d)", e.p.x, e.p.y)
}

// ErrBadDistance is returned if a distance used is not "Near", "Middle" or "Far"
type ErrBadDistance struct {
	d string
}

func (e ErrBadDistance) Error() string {
	return fmt.Sprintf("bad distance definition '%s'", e.d)
}

// // Constants defining screen pixels to display in the view window
// const (
// 	PW uint8 = iota // Wall
// 	PO              // Open
// 	PE              // Empty
// 	PF              // Floor
// 	PC              // Ceiling
// 	GB              // Gopher body
// )
// const (
// 	// TN represents transparent (used for overlays)
// 	TN = 32
// )

type displayType uint8

// Wall panel dispositions
const (
	DSideWall       = "sidewall"
	DOpenWallNear   = "opennewar"
	DOpenWallMiddle = "openmiddle"
	DOpenWallFar    = "openfar"
	DEmpty          = "empty"
)

// type view struct {
// 	window [numPanels]element.PixelMatrix
// }

func (g *Game) updateView() error {

	// The window is comprised of 7 vertical slices
	// The outer two each side are two columns;
	// the middle three are single column
	//
	//   |  |  | | | |  |  |
	//   |  |  | | | |  |  |
	//   |  |  | | | |  |  |
	//
	// Scan around us to render (adjust for direction!)
	//
	//    +----+----+----+
	//    |    | F2 |    |
	//    +----+----+----+
	//    | L2 | F1 | R2 |
	//    +----+----+----+
	//    | L1 | F  | R1 |
	//    +----+----+----+
	//    | L  | P  | R  |
	//    +----+----+----+

	// TODO Assuming NORTH look for now

	// Determine the points to our left and right and then we can
	// position from there into the grid.
	var lp point
	var rp point
	var fp point

	var mx, my int8
	p := g.p

	switch p.o {
	case 'n':
		lp = point{p.x - 1, p.y}
		rp = point{p.x + 1, p.y}
		fp = point{p.x, p.y - 1}
		mx = 0
		my = -1
	case 's':
		lp = point{p.x + 1, p.y}
		rp = point{p.x - 1, p.y}
		fp = point{p.x, p.y + 1}
		mx = 0
		my = 1
	case 'e':
		lp = point{p.x, p.y - 1}
		rp = point{p.x, p.y + 1}
		fp = point{p.x + 1, p.y}
		mx = 1
		my = 0
	case 'w':
		lp = point{p.x, p.y + 1}
		rp = point{p.x, p.y - 1}
		fp = point{p.x - 1, p.y}
		mx = -1
		my = 0
	}

	var isWall bool
	var err error

	// L, R and F
	isWall, err = g.renderSpace(lp, rp, fp, 0, 6, "Near")
	if err != nil {
		return err
	}
	if isWall {
		return nil
	}

	// L1, R1 and F1
	lp = point{lp.x + mx, lp.y + my}
	rp = point{rp.x + mx, rp.y + my}
	fp = point{fp.x + mx, fp.y + my}

	debug.Printf("Checking L1 (%d,%d)[%c] R1 (%d,%d)[%c]", lp.x, lp.y, g.m.getSpace(lp).t, rp.x, rp.y, g.m.getSpace(rp).t)
	isWall, err = g.renderSpace(lp, rp, fp, 1, 5, "Middle")
	if err != nil {
		return err
	}
	if isWall {
		return nil
	}

	// L2, R2 and F2
	lp = point{lp.x + mx, lp.y + my}
	rp = point{rp.x + mx, rp.y + my}
	fp = point{fp.x + mx, fp.y + my}

	debug.Printf("Checking L2 (%d,%d)[%c] R2 (%d,%d)[%c]", lp.x, lp.y, g.m.getSpace(lp).t, rp.x, rp.y, g.m.getSpace(rp).t)
	isWall, err = g.renderSpace(lp, rp, fp, 2, 4, "Far")
	if err != nil {
		return err
	}
	if isWall {
		return nil
	}

	// Final step
	fp = point{fp.x + mx, fp.y + my}
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp3) is (wall)")
		g.v[3] = element.Panels[3]["OpenWallFar"]
	case SpaceEmpty:
		debug.Println("Space in (fp3) is (empty)")
		g.v[3] = element.Panels[3]["Empty"]
	default:
		return ErrBadSpace{fp}
	}

	return nil
}

func (g *Game) renderSpace(lp, rp, fp point, lpanel, rpanel int, distance string) (isWall bool, err error) {
	if err = g.renderLeftRight(lp, lpanel, distance); err != nil {
		return false, err
	}
	if err = g.renderLeftRight(rp, rpanel, distance); err != nil {
		return false, err
	}
	if isWall, err = g.renderFront(fp, "Far"); err != nil {
		return false, err
	}
	return isWall, nil
}

func (g *Game) renderLeftRight(p point, panel int, distance string) error {
	switch g.m.getSpace(p).t {
	case SpaceWall:
		g.v[panel] = element.Panels[panel]["SideWall"]
		debug.Printf("Render point (%v) as wall", p)
	case SpaceEmpty:
		g.v[panel] = element.Panels[panel]["OpenWall"+distance]
		debug.Printf("Render point (%v) as open", p)
	default:
		return ErrBadSpace{p}
	}
	return nil
}

func (g *Game) renderFront(fp point, distance string) (bool, error) {
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		switch distance {
		case "Near":
			g.v[1] = element.Panels[1]["OpenWall"+distance]
			g.v[5] = element.Panels[5]["OpenWall"+distance]
			fallthrough
		case "Middle":
			g.v[2] = element.Panels[2]["OpenWall"+distance]
			g.v[4] = element.Panels[4]["OpenWall"+distance]
			fallthrough
		case "Far":
			g.v[3] = element.Panels[3]["OpenWall"+distance]
		default:
			return false, ErrBadDistance{distance}
		}
		return true, nil
	case SpaceEmpty:
		return false, nil
	}
	return false, ErrBadSpace{fp}
}

func (g *Game) render() (string, error) {

	// Construct the viewport first before we overlay and sprites.
	// This comprises a set of virtual scanlines.
	scanlines := make(pixelMatrix, viewHeight)

	wallColourMod := 0
	if g.p.o == 'e' || g.p.o == 'w' {
		wallColourMod = 1
	}

	// SPRITE
	// overlay, err := g.g.sprite(1)
	// if err != nil {
	// 	return "", errors.New("error rendering gopther sprite to viewport")
	// }
	// SPRITE

	for y := 0; y < viewHeight; y++ {
		debug.Println("Render scan line:", y)

		// Instantiate to the width * 2 as each "pixel" block is
		// actally two runes wide
		scanline := make([]rune, viewWidth*2)

		// Mark the absolute x position as we read through the panels so
		// we can replace with any overlay in the correct positions as and
		// when appropriate.
		x := 0

		for c := 0; c < numPanels; c++ {
			panel := g.v[c][y]
			debug.Println("Panel", c)

			for _, pixel := range panel {

				// 	// debug.Printf("Overlay on panel %2d at %d, %d\n", c, x, y)
				// 	// if overlay[y][x] != TN {
				// 	// 	scanline += "XX"
				// 	// 	continue
				// 	// }

				// TODO Better handle doubling - interpolate into slices?
				switch pixel {
				case element.W:
					debug.Println(" Pixel WALL")
					scanline[x] = walls[(wallColourMod % 2)]
					scanline[x+1] = walls[(wallColourMod % 2)]
				case element.O:
					debug.Println(" Pixel OPEN")
					scanline[x] = walls[((wallColourMod + 1) % 2)]
					scanline[x+1] = walls[((wallColourMod + 1) % 2)]
				default:
					debug.Println(" Pixel TRANSPARENT")
					scanline[x] = rune(' ')
					scanline[x+1] = rune(' ')
				}
				x += 2
			}

		}
		debug.Println("Scanline:", scanline)
		scanlines[y] = []rune(scanline)
	}

	// for _, sl := range scanlines {
	// 	debug.Println(sl)
	// }

	output := "╔════════════════════════╗\n"
	for _, sl := range scanlines {
		output += fmt.Sprintf("║ %s ║\n", string(sl))
	}
	output += "╚════════════════════════╝\n" + fmt.Sprintf("Facing: %s\n", bytes.ToUpper([]byte{g.p.o})) + "\nWhich way?: "

	return output, nil
}
