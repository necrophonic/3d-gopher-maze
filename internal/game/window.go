package game

import (
	"bytes"
	"fmt"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/necrophonic/gopher-maze/internal/game/element"
	"github.com/pkg/errors"
)

const numPanels = 7
const viewHeight = 9
const viewWidth = 11

type point struct {
	x int8
	y int8
}

// Is returns whether the given point is
// the same location as this point
func (p point) Is(p2 point) bool {
	return p.x == p2.x && p.y == p2.y
}

// TODO refactor swtches
var walls = [2]rune{'▒', '░'}

// View is a compiled slice of pixel matricies representing
// panels to be displayed in the viewport.
type view struct {
	screen  []element.PixelMatrix
	overlay element.PixelMatrix
}

// Clear empties an existing view
func (v view) Clear() {
	for i := 0; i < len(v.screen); i++ {
		v.screen[i].Clear()
	}
}

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

type displayType uint8

// Wall panel dispositions
const (
	DSideWall       = "sidewall"
	DOpenWallNear   = "opennewar"
	DOpenWallMiddle = "openmiddle"
	DOpenWallFar    = "openfar"
	DEmpty          = "empty"
)

func (g *Game) updateView() error {

	g.v.overlay = nil

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

	fp, isWall, err := g.renderSpace(lp, rp, fp, mx, my, Near)
	if err != nil {
		return err
	}
	if isWall {
		debug.Println("Render break at wall")
		return nil
	}

	return nil
}

// Distances
const (
	Near   = 0
	Middle = 1
	Far    = 2
)

var distances = []string{"Near", "Middle", "Far"}
var panelPairs = [][]int{{0, 6}, {1, 5}, {2, 4}}

func (g *Game) checkSpaceForItem(p point, distance int) (err error) {
	for _, item := range g.items {
		if p.Is(item.GetPoint()) {
			g.v.overlay, err = item.GetMatrix(distance)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Game) renderSpace(lp, rp, fp point, mx, my int8, distance int) (nfp point, isWall bool, err error) {

	if distance == 3 {
		// As far as we render, so break out
		return fp, false, nil
	}
	leftPanel := panelPairs[distance][0]
	rightPanel := panelPairs[distance][1]

	debug.Printf("Render L (%v), R (%v), F (%v) at %s\n", lp, rp, fp, distances[distance])

	if err = g.renderLeftRight(lp, leftPanel, distances[distance]); err != nil {
		return point{}, false, err
	}
	if err = g.renderLeftRight(rp, rightPanel, distances[distance]); err != nil {
		return point{}, false, err
	}

	if isWall, err = g.renderFront(fp, distances[distance]); err != nil {
		return point{}, false, err
	}
	if !isWall && distance != 4 {
		lp = point{lp.x + mx, lp.y + my}
		rp = point{rp.x + mx, rp.y + my}
		fp = point{fp.x + mx, fp.y + my}
		distance++

		// Check for items
		if err := g.checkSpaceForItem(fp, distance); err != nil {
			return point{}, false, errors.WithMessage(err, "failed to check space for items")
		}

		return g.renderSpace(lp, rp, fp, mx, my, distance)
	}
	debug.Printf("Return FP %v", fp)
	return fp, true, nil
}

func (g *Game) renderLeftRight(p point, panel int, distance string) error {
	switch g.m.getSpace(p).t {
	case SpaceWall:
		g.v.screen[panel] = element.Panels[panel]["SideWall"]
		debug.Printf("Render  side point (%v) as wall", p)
	case SpaceEmpty:
		g.v.screen[panel] = element.Panels[panel]["OpenWall"+distance]
		debug.Printf("Render  side (%v) as open", p)
	default:
		return ErrBadSpace{p}
	}
	return nil
}

func (g *Game) renderFront(fp point, distance string) (bool, error) {
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Printf("Render front point (%v) as wall", fp)
		switch distance {
		case "Near":
			g.v.screen[1] = element.Panels[1]["OpenWall"+distance]
			g.v.screen[5] = element.Panels[5]["OpenWall"+distance]
			fallthrough
		case "Middle":
			g.v.screen[2] = element.Panels[2]["OpenWall"+distance]
			g.v.screen[4] = element.Panels[4]["OpenWall"+distance]
			fallthrough
		case "Far":
			g.v.screen[3] = element.Panels[3]["OpenWall"+distance]
		default:
			return false, ErrBadDistance{distance}
		}
		return true, nil
	case SpaceEmpty:
		debug.Printf("Render front point (%v) as empty", fp)
		g.v.screen[3] = element.Panels[3]["Empty"]
		return false, nil
	}
	return false, ErrBadSpace{fp}
}

var (
	gopherBody = []rune{'█', '█'}
)

func (g *Game) render() (string, error) {
	// render() will render a compiled view to a displayable string

	// Construct the viewport first before we overlay and sprites.
	// This comprises a set of virtual scanlines.
	scanlines := make([][]rune, viewHeight)

	wallColourMod := 0
	if g.p.o == 'e' || g.p.o == 'w' {
		wallColourMod = 1
	}

	// SPRITE
	// overlay, err := g.gopher.sprite(1)
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
		x := -2
		xi := -1

		for c := 0; c < numPanels; c++ {
			panel := g.v.screen[c][y]
			for _, pixel := range panel {
				x += 2
				xi++

				// TODO Account for overlay needing to define every
				// half "pixel" rather than the doubling of the main
				// view
				if g.v.overlay != nil {
					debug.Printf("Overlay on panel %2d at %d, %d from overlay %d,%d\n", c, x, y, xi, y)
					if g.v.overlay[y][xi] != element.T {
						switch g.v.overlay[y][xi] {
						case element.G:
							scanline[x] = gopherBody[0]
							scanline[x+1] = gopherBody[1]
						}
						continue
					}
				}

				// TODO Better handle doubling - interpolate into slices?
				switch pixel {
				case element.W:
					// debug.Println(" Pixel WALL")
					scanline[x] = walls[((wallColourMod + 1) % 2)]
					scanline[x+1] = walls[((wallColourMod + 1) % 2)]
				case element.O:
					// debug.Println(" Pixel OPEN")
					scanline[x] = walls[(wallColourMod % 2)]
					scanline[x+1] = walls[(wallColourMod % 2)]
				default:
					// debug.Println(" Pixel TRANSPARENT")
					scanline[x] = rune(' ')
					scanline[x+1] = rune(' ')
				}
			}

		}
		scanlines[y] = scanline
	}

	output := "╔════════════════════════╗\n"
	for _, sl := range scanlines {
		output += fmt.Sprintf("║ %s ║\n", string(sl))
	}
	output += "╚════════════════════════╝\n" + fmt.Sprintf("Facing: %s\n", bytes.ToUpper([]byte{g.p.o})) + "\nWhich way?: "

	return output, nil
}
