package game

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/necrophonic/gopher-maze/internal/debug"
)

type pixel uint8
type windowSlice [][]pixel
type point struct {
	x int8
	y int8
}

var walls = [2]rune{'▒', '░'}

// ErrBadSpace is returned if the space definition in a
// point location is unexpected
type ErrBadSpace struct {
	p point
}

func (e ErrBadSpace) Error() string {
	return fmt.Sprintf("bad space definition at (%d,%d)", e.p.x, e.p.y)
}

// Constants defining screen pixels to display in the view window
const (
	PW pixel = iota // Wall
	PO              // Open
	PE              // Empty
	PF              // Floor
	PC              // Ceiling
)

type displayType uint8

// Wall panel dispositions
const (
	DSideWall displayType = iota // A closed wall
	DOpenWallNear
	DOpenWallMiddle
	DOpenWallFar
	DEmpty
)

type view struct {
	w      int
	h      int
	window [7]windowSlice
}

// type view [7]windowSlice

func (g *Game) updateView() error {

	// The window is comprised of 7 vertical slices
	// The outer two each side are two columns;
	// the middle three are single column
	//
	//   |  |  | | | |  |  |
	//   |  |  | | | |  |  |
	//   |  |  | | | |  |  |
	//
	// window := [7]windowSlice{}

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

	// ------

	// g.v.window[0] = g.m.panels[0][DSideWall]
	// g.v.window[1] = g.m.panels[1][DSideWall]
	// g.v.window[2] = g.m.panels[2][DSideWall]
	// g.v.window[3] = g.m.panels[3][DEmpty]
	// g.v.window[4] = g.m.panels[4][DSideWall]
	// g.v.window[5] = g.m.panels[5][DSideWall]
	// g.v.window[6] = g.m.panels[6][DSideWall]

	// window[0] = g.m.panels[0][DOpenWallNear]
	// window[1] = g.m.panels[1][DSideWall]
	// window[2] = g.m.panels[2][DOpenWallFar]
	// window[3] = g.m.panels[3][DOpenWallFar]
	// window[4] = g.m.panels[4][DSideWall]
	// window[5] = g.m.panels[5][DSideWall]
	// window[6] = g.m.panels[6][DSideWall]

	// window[0] = g.m.panels[0][DSideWall]
	// window[1] = g.m.panels[1][DOpenWallMiddle]
	// window[2] = g.m.panels[2][DSideWall]
	// window[3] = g.m.panels[3][DEmpty]
	// window[4] = g.m.panels[4][DSideWall]
	// window[5] = g.m.panels[5][DOpenWallMiddle]
	// window[6] = g.m.panels[6][DSideWall]

	// window[0] = g.m.panels[0][DOpenWallNear]
	// window[1] = g.m.panels[1][DSideWall]
	// window[2] = g.m.panels[2][DSideWall]
	// window[3] = g.m.panels[3][DEmpty]
	// window[4] = g.m.panels[4][DSideWall]
	// window[5] = g.m.panels[5][DOpenWallMiddle]
	// window[6] = g.m.panels[6][DSideWall]

	p := g.p

	// TODO Assuming NORTH look for now

	// Determine the points to our left and right and then we can
	// position from there into the grid.
	var lp point
	var rp point
	var fp point

	var mx, my int8

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

	// Check left (L) and right (R) first

	// L
	switch g.m.getSpace(lp).t {
	case SpaceWall:
		g.v.window[0] = g.m.panels[0][DSideWall]
	case SpaceEmpty:
		g.v.window[0] = g.m.panels[0][DOpenWallNear]
	}

	// R
	switch g.m.getSpace(rp).t {
	case SpaceWall:
		g.v.window[6] = g.m.panels[6][DSideWall]
	case SpaceEmpty:
		g.v.window[6] = g.m.panels[6][DOpenWallNear]
	}

	// Then check front. If it's a wall then panels 1-6 are Wall Near.
	// Otherwise we can check L1 and R1
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp) is (wall)")
		g.v.window[1] = g.m.panels[1][DOpenWallNear]
		g.v.window[2] = g.m.panels[2][DOpenWallNear]
		g.v.window[3] = g.m.panels[3][DOpenWallNear]
		g.v.window[4] = g.m.panels[4][DOpenWallNear]
		g.v.window[5] = g.m.panels[5][DOpenWallNear]
		return nil
	case SpaceEmpty:
		debug.Println("Space in (fp) is (empty)")

		// Move up a row in the direction we're facing
		lp = point{lp.x + mx, lp.y + my}
		rp = point{rp.x + mx, rp.y + my}

		debug.Printf("Checking L1 (%d,%d)[%c] R1 (%d,%d)[%c]", lp.x, lp.y, g.m.getSpace(lp).t, rp.x, rp.y, g.m.getSpace(rp).t)

		// L1
		switch g.m.getSpace(lp).t {
		case SpaceWall:
			g.v.window[1] = g.m.panels[1][DSideWall]
		case SpaceEmpty:
			g.v.window[1] = g.m.panels[1][DOpenWallMiddle]
		default:
			return ErrBadSpace{lp}
		}

		// R1
		switch g.m.getSpace(rp).t {
		case SpaceWall:
			g.v.window[5] = g.m.panels[5][DSideWall]
		case SpaceEmpty:
			g.v.window[5] = g.m.panels[5][DOpenWallMiddle]
		default:
			return ErrBadSpace{rp}
		}

	}

	// FP 1 -------

	// Move forward again
	fp = point{fp.x + mx, fp.y + my}

	// Then check front. If it's a wall then panels 1-6 are Wall Near.
	// Otherwise we can check L1 and R1
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp1) is (wall)")
		g.v.window[2] = g.m.panels[2][DOpenWallMiddle]
		g.v.window[3] = g.m.panels[3][DOpenWallMiddle]
		g.v.window[4] = g.m.panels[4][DOpenWallMiddle]
		return nil
	case SpaceEmpty:
		debug.Println("Space in (fp1) is (empty)")

		// Move up a row in the direction we're facing
		lp = point{lp.x + mx, lp.y + my}
		rp = point{rp.x + mx, rp.y + my}

		debug.Printf("Checking L2 (%d,%d)[%c] R2 (%d,%d)[%c]", lp.x, lp.y, g.m.getSpace(lp).t, rp.x, rp.y, g.m.getSpace(rp).t)

		// L2
		switch g.m.getSpace(lp).t {
		case SpaceWall:
			g.v.window[2] = g.m.panels[2][DSideWall]
		case SpaceEmpty:
			g.v.window[2] = g.m.panels[2][DOpenWallFar]
		default:
			return ErrBadSpace{lp}
		}

		// R2
		switch g.m.getSpace(rp).t {
		case SpaceWall:
			g.v.window[4] = g.m.panels[4][DSideWall]
		case SpaceEmpty:
			g.v.window[4] = g.m.panels[4][DOpenWallFar]
		default:
			return ErrBadSpace{rp}
		}

	}

	// FP 2 -------

	// Move forward again
	fp = point{fp.x + mx, fp.y + my}
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp2) is (wall)")
		g.v.window[3] = g.m.panels[3][DOpenWallFar]
		return nil
	case SpaceEmpty:
		debug.Println("Space in (fp2) is (empty)")

		// Move up a row in the direction we're facing
		lp = point{lp.x + mx, lp.y + my}
		rp = point{rp.x + mx, rp.y + my}
	}

	// Final step
	fp = point{fp.x + mx, fp.y + my}
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp3) is (wall)")
		g.v.window[3] = g.m.panels[3][DOpenWallFar]
	case SpaceEmpty:
		debug.Println("Space in (fp3) is (empty)")
		g.v.window[3] = g.m.panels[3][DEmpty]
	}

	return nil
}

func (g *Game) render() string {

	frames := map[string][]string{
		"1": {
			"╔════════════════════════╗\n",
			"╚════════════════════════╝\n",
		},
		"2": {
			"╔════════════════════════════════════════════╗\n",
			"╚════════════════════════════════════════════╝\n",
		},
	}

	output := frames[g.m.scale][0]

	wallColourMod := 0
	if g.p.o == 'e' || g.p.o == 'w' {
		wallColourMod = 1
	}

	numPanels := 7

	for y := 0; y < g.v.h; y++ {
		output += "║ "
		for c := 0; c < numPanels; c++ {
			panel := g.v.window[c]

			for _, pxl := range panel[y] {
				switch pxl {
				case PW:
					output += strings.Repeat(string(walls[(wallColourMod%2)]), 2)
				case PO:
					output += strings.Repeat(string(walls[(wallColourMod+1)%2]), 2)
				default:
					output += "  "
				}
			}

		}
		output += " ║\n"
	}
	return output + frames[g.m.scale][1] + fmt.Sprintf("Facing: %s\n", bytes.ToUpper([]byte{g.p.o})) + "\nWhich way?: "
}
