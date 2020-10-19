package game

import (
	"fmt"
	"strings"

	"github.com/necrophonic/gopher-maze/internal/debug"
)

// type panel uint8
type pixel uint8
type windowSlice [][]pixel
type point struct {
	x int8
	y int8
}

// ErrBadSpace is returned if the space definition in a
// point location is unexpected
type ErrBadSpace struct {
	p point
}

func (e ErrBadSpace) Error() string {
	return fmt.Sprintf("bad space definition at (%d,%d)", e.p.x, e.p.y)
}

const scale = 2

// Constants defining screen pixels to display in the view window
const (
	PW pixel = iota
	PO
	PE // Floor, ceiling and distant wall out of view range
)

type displayType uint8

const (
	DSideWall displayType = iota // A closed wall
	DOpenWallNear
	DOpenWallMiddle
	DOpenWallFar
	DEmpty
)

type view [7]windowSlice

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

	g.v[0] = panels[0][DSideWall]
	g.v[1] = panels[1][DSideWall]
	g.v[2] = panels[2][DSideWall]
	g.v[3] = panels[3][DEmpty]
	g.v[4] = panels[4][DSideWall]
	g.v[5] = panels[5][DSideWall]
	g.v[6] = panels[6][DSideWall]

	// window[0] = panels[0][DOpenWallNear]
	// window[1] = panels[1][DSideWall]
	// window[2] = panels[2][DOpenWallFar]
	// window[3] = panels[3][DOpenWallFar]
	// window[4] = panels[4][DSideWall]
	// window[5] = panels[5][DSideWall]
	// window[6] = panels[6][DSideWall]

	// window[0] = panels[0][DSideWall]
	// window[1] = panels[1][DOpenWallMiddle]
	// window[2] = panels[2][DSideWall]
	// window[3] = panels[3][DEmpty]
	// window[4] = panels[4][DSideWall]
	// window[5] = panels[5][DOpenWallMiddle]
	// window[6] = panels[6][DSideWall]

	// window[0] = panels[0][DOpenWallNear]
	// window[1] = panels[1][DSideWall]
	// window[2] = panels[2][DSideWall]
	// window[3] = panels[3][DEmpty]
	// window[4] = panels[4][DSideWall]
	// window[5] = panels[5][DOpenWallMiddle]
	// window[6] = panels[6][DSideWall]

	p := g.p

	// TODO Assuming NORTH look for now

	// Determine the points to our left and right and then we can
	// position from there into the grid.
	var lp point
	var rp point
	var fp point

	// Move modifiers
	var mmx int8
	var mmy int8

	switch p.o {
	case 'n':
		lp = point{p.x - 1, p.y}
		rp = point{p.x + 1, p.y}
		fp = point{p.x, p.y - 1}
		mmx = 0
		mmy = -1
	case 's':
		lp = point{p.x + 1, p.y}
		rp = point{p.x - 1, p.y}
		fp = point{p.x, p.y + 1}
		mmx = 0
		mmy = 1
	case 'e':
		lp = point{p.x, p.y - 1}
		rp = point{p.x, p.y + 1}
		fp = point{p.x + 1, p.y}
		mmx = 1
		mmy = 0
	case 'w':
		lp = point{p.x, p.y + 1}
		rp = point{p.x, p.y - 1}
		fp = point{p.x - 1, p.y}
		mmx = -1
		mmy = 0
	}

	// Check left (L) and right (R) first

	// L
	switch g.m.getSpace(lp).t {
	case SpaceWall:
		g.v[0] = panels[0][DSideWall]
	case SpaceEmpty:
		g.v[0] = panels[0][DOpenWallNear]
	}

	// R
	switch g.m.getSpace(rp).t {
	case SpaceWall:
		g.v[6] = panels[6][DSideWall]
	case SpaceEmpty:
		g.v[6] = panels[6][DOpenWallNear]
	}

	// Then check front. If it's a wall then panels 1-6 are Wall Near.
	// Otherwise we can check L1 and R1
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp) is (wall)")
		g.v[1] = panels[1][DOpenWallNear]
		g.v[2] = panels[2][DOpenWallNear]
		g.v[3] = panels[3][DOpenWallNear]
		g.v[4] = panels[4][DOpenWallNear]
		g.v[5] = panels[5][DOpenWallNear]
		return nil
	case SpaceEmpty:
		debug.Println("Space in (fp) is (empty)")

		// Move up a row in the direction we're facing
		lp = point{lp.x + mmx, lp.y + mmy}
		rp = point{rp.x + mmx, rp.y + mmy}

		debug.Printf("Checking L1 (%d,%d)[%c] R1 (%d,%d)[%c]", lp.x, lp.y, g.m.getSpace(lp).t, rp.x, rp.y, g.m.getSpace(rp).t)

		// L1
		switch g.m.getSpace(lp).t {
		case SpaceWall:
			g.v[1] = panels[1][DSideWall]
		case SpaceEmpty:
			g.v[1] = panels[1][DOpenWallMiddle]
		default:
			return ErrBadSpace{lp}
		}

		// R1
		switch g.m.getSpace(rp).t {
		case SpaceWall:
			g.v[5] = panels[5][DSideWall]
		case SpaceEmpty:
			g.v[5] = panels[5][DOpenWallMiddle]
		default:
			return ErrBadSpace{rp}
		}

	}

	// FP 1 -------

	// Move forward again
	fp = point{fp.x + mmx, fp.y + mmy}

	// Then check front. If it's a wall then panels 1-6 are Wall Near.
	// Otherwise we can check L1 and R1
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp1) is (wall)")
		g.v[2] = panels[2][DOpenWallMiddle]
		g.v[3] = panels[3][DOpenWallMiddle]
		g.v[4] = panels[4][DOpenWallMiddle]
		return nil
	case SpaceEmpty:
		debug.Println("Space in (fp1) is (empty)")

		// Move up a row in the direction we're facing
		lp = point{lp.x + mmx, lp.y + mmy}
		rp = point{rp.x + mmx, rp.y + mmy}

		debug.Printf("Checking L2 (%d,%d)[%c] R2 (%d,%d)[%c]", lp.x, lp.y, g.m.getSpace(lp).t, rp.x, rp.y, g.m.getSpace(rp).t)

		// L2
		switch g.m.getSpace(lp).t {
		case SpaceWall:
			g.v[2] = panels[2][DSideWall]
		case SpaceEmpty:
			g.v[2] = panels[2][DOpenWallFar]
		default:
			return ErrBadSpace{lp}
		}

		// R2
		switch g.m.getSpace(rp).t {
		case SpaceWall:
			g.v[4] = panels[4][DSideWall]
		case SpaceEmpty:
			g.v[4] = panels[4][DOpenWallFar]
		default:
			return ErrBadSpace{rp}
		}

	}

	// FP 2 -------

	// Move forward again
	fp = point{fp.x + mmx, fp.y + mmy}
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp2) is (wall)")
		g.v[3] = panels[3][DOpenWallFar]
		return nil
	case SpaceEmpty:
		debug.Println("Space in (fp2) is (empty)")

		// Move up a row in the direction we're facing
		lp = point{lp.x + mmx, lp.y + mmy}
		rp = point{rp.x + mmx, rp.y + mmy}
	}

	// Final step
	fp = point{fp.x + mmx, fp.y + mmy}
	switch g.m.getSpace(fp).t {
	case SpaceWall:
		debug.Println("Space in (fp3) is (wall)")
		g.v[3] = panels[3][DOpenWallFar]
	case SpaceEmpty:
		debug.Println("Space in (fp3) is (empty)")
		g.v[3] = panels[3][DEmpty]
	}

	return nil
}

// func (v *view) renderPanel(column int, p point) {

// }

func (g *Game) render() string {

	output := "╔════════════════════════╗\n"

	for y := 0; y < displayHeight; y++ {
		output += "║ "
		for c := 0; c < 7; c++ {
			panel := g.v[c]

			for _, pxl := range panel[y] {
				switch pxl {
				case PW:
					output += strings.Repeat(string(walls[0]), scale)
				case PO:
					output += strings.Repeat(string(walls[1]), scale)
				case PE:
					output += "  "
				}
			}

		}
		output += " ║\n"
	}
	return output + "╚════════════════════════╝\n"
}

type windowColumn uint8

// Common sequences
var (
	empty345          = windowSlice{{PE}, {PE}, {PE}, {PE}, {PE}, {PE}, {PE}, {PE}, {PE}}
	d06OpenWall       = windowSlice{{PE, PE}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PE, PE}}
	d15OpenWallMiddle = windowSlice{{PE, PE}, {PE, PE}, {PE, PE}, {PO, PO}, {PO, PO}, {PO, PO}, {PE, PE}, {PE, PE}, {PE, PE}}
	d15OpenWallNear   = windowSlice{{PE, PE}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PO, PO}, {PE, PE}}
	d24SideWall       = windowSlice{{PE}, {PE}, {PE}, {PE}, {PW}, {PE}, {PE}, {PE}, {PE}}
	d24OpenWallNear   = windowSlice{{PE}, {PO}, {PO}, {PO}, {PO}, {PO}, {PO}, {PO}, {PE}}
	d24OpenWallMiddle = windowSlice{{PE}, {PE}, {PE}, {PO}, {PO}, {PO}, {PE}, {PE}, {PE}}
	d24OpenWallFar    = windowSlice{{PE}, {PE}, {PE}, {PE}, {PO}, {PE}, {PE}, {PE}, {PE}}
)

var panels = map[windowColumn]map[displayType]windowSlice{
	0: {
		DSideWall:     {{PW, PE}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PE}},
		DOpenWallNear: d06OpenWall,
	},
	1: {
		DSideWall:       {{PE, PE}, {PE, PE}, {PW, PE}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PE}, {PE, PE}, {PE, PE}},
		DOpenWallMiddle: d15OpenWallMiddle,
		DOpenWallNear:   d15OpenWallNear,
	},
	2: {
		DSideWall:       d24SideWall,
		DOpenWallNear:   d24OpenWallNear,
		DOpenWallMiddle: d24OpenWallMiddle,
		DOpenWallFar:    d24OpenWallFar,
	},
	3: {
		DEmpty:          {{PE}, {PE}, {PE}, {PE}, {PE}, {PE}, {PE}, {PE}, {PE}},
		DOpenWallNear:   d24OpenWallNear,
		DOpenWallMiddle: d24OpenWallMiddle,
		DOpenWallFar:    d24OpenWallFar,
	},
	4: {
		DSideWall:       d24SideWall,
		DOpenWallNear:   d24OpenWallNear,
		DOpenWallMiddle: d24OpenWallMiddle,
		DOpenWallFar:    d24OpenWallFar,
	},
	5: {
		DSideWall:       {{PE, PE}, {PE, PE}, {PE, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PE, PW}, {PE, PE}, {PE, PE}},
		DOpenWallMiddle: d15OpenWallMiddle,
		DOpenWallNear:   d15OpenWallNear,
	},
	6: {
		DSideWall:     {{PE, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PE, PW}},
		DOpenWallNear: d06OpenWall,
	},
}
