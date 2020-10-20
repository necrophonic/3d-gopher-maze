package game

import "fmt"

type windowColumn uint8

type components struct {
	width  int
	height int

	d06OpenWall       windowSlice
	d15OpenWallMiddle windowSlice
	d15OpenWallNear   windowSlice
	d2SideWall        windowSlice
	d4SideWall        windowSlice
	d24OpenWallNear   windowSlice
	d24OpenWallMiddle windowSlice
	d24OpenWallFar    windowSlice

	d3Empty windowSlice

	d0SideWall windowSlice
	d1SideWall windowSlice
	d5SideWall windowSlice
	d6SideWall windowSlice
}

var pe1 = []pixel{PE}

var po2 = []pixel{PO, PO}
var pe2 = []pixel{PE, PE}
var pf2 = []pixel{PF, PF}
var pw2 = []pixel{PW, PW}

var po4 = []pixel{PO, PO, PO, PO}
var pe4 = []pixel{PE, PE, PE, PE}
var pf4 = []pixel{PF, PF, PF, PF}
var pw4 = []pixel{PW, PW, PW, PW}

var po3pw = []pixel{PO, PO, PO, PW}

var scaled = map[string]components{
	"1": {
		width:  11,
		height: 9,

		d06OpenWall:       windowSlice{pe2, po2, po2, po2, po2, po2, po2, po2, {PF, PF}},
		d15OpenWallMiddle: windowSlice{pe2, pe2, pe2, po2, po2, po2, {PF, PF}, {PF, PF}, {PF, PF}},
		d15OpenWallNear:   windowSlice{pe2, po2, po2, po2, po2, po2, po2, po2, {PF, PF}},
		d2SideWall:        windowSlice{{PE}, {PE}, {PE}, {PE}, {PW}, {PF}, {PF}, {PF}, {PF}},
		d4SideWall:        windowSlice{{PE}, {PE}, {PE}, {PE}, {PW}, {PF}, {PF}, {PF}, {PF}},
		d24OpenWallNear:   windowSlice{{PE}, {PO}, {PO}, {PO}, {PO}, {PO}, {PO}, {PO}, {PF}},
		d24OpenWallMiddle: windowSlice{{PE}, {PE}, {PE}, {PO}, {PO}, {PO}, {PF}, {PF}, {PF}},
		d24OpenWallFar:    windowSlice{{PE}, {PE}, {PE}, {PE}, {PO}, {PF}, {PF}, {PF}, {PF}},

		d3Empty: windowSlice{pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1},

		d0SideWall: windowSlice{{PW, PE}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PE}},
		d1SideWall: windowSlice{pe2, pe2, {PW, PE}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PF}, {PF, PF}, {PF, PF}},
		d5SideWall: windowSlice{pe2, pe2, {PE, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PF, PW}, {PF, PF}, {PF, PF}},
		d6SideWall: windowSlice{{PE, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PW, PW}, {PF, PW}},
	},
	// TODO - Fix the scaling of scale 2!!!
	"2": {
		width:  23,
		height: 19,

		d06OpenWall:       windowSlice{pe4, pe4, {PO, PO, PO, PE}, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, po3pw, {PO, PO, PO, PE}, pf4, pf4},
		d15OpenWallMiddle: windowSlice{pe4, pe4, pe4, pe4, pe4, po4, po4, po4, po4, po4, po4, po4, po4, po4, pf4, pf4, pf4, pf4, pf4},
		d15OpenWallNear:   windowSlice{pe4, pe4, pe4, po4, po4, po4, po4, po4, po4, po4, po4, po4, po4, po4, po4, po4, pf4, pf4, pf4},

		d2SideWall: windowSlice{pe2, pe2, pe2, pe2, pe2, pe2, pe2, pe2, {PW, PE}, pw2, {PW, PE}, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2},
		d4SideWall: windowSlice{pe2, pe2, pe2, pe2, pe2, pe2, pe2, pe2, {PE, PW}, pw2, {PE, PW}, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2},

		d24OpenWallNear:   windowSlice{pe2, po2, po2, po2, po2, po2, po2, po2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2},
		d24OpenWallMiddle: windowSlice{pe2, pe2, pe2, po2, po2, po2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2},
		d24OpenWallFar:    windowSlice{pe2, pe2, pe2, pe2, po2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2, pf2},

		//
		d3Empty: windowSlice{pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1, pe1},

		d0SideWall: windowSlice{
			{PW, PE, PE, PE}, {PW, PW, PE, PE}, {PW, PW, PW, PE}, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, {PW, PW, PW, PE}, {PW, PW, PE, PE}, {PW, PE, PE, PE},
		},
		d1SideWall: windowSlice{
			pe4, pe4, pe4, pe4, {PW, PE, PE, PE}, {PW, PW, PE, PE}, {PW, PW, PW, PE}, pw4, pw4, pw4, pw4, pw4, {PW, PW, PW, PE}, {PW, PW, PE, PE}, {PW, PE, PE, PE}, pe4, pe4, pe4, pe4,
		},
		d5SideWall: windowSlice{
			pe4, pe4, pe4, pe4, {PE, PE, PE, PW}, {PE, PE, PW, PW}, {PE, PW, PW, PW}, pw4, pw4, pw4, pw4, pw4, {PE, PW, PW, PW}, {PE, PE, PW, PW}, {PE, PE, PE, PW}, pe4, pe4, pe4, pe4,
		},
		d6SideWall: windowSlice{
			{PE, PE, PE, PW}, {PE, PE, PW, PW}, {PE, PW, PW, PW}, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, pw4, {PE, PW, PW, PW}, {PE, PE, PW, PW}, {PE, PE, PE, PW},
		},
	},
}

func (g *Game) setUpScaledPanels() error {

	var cmpts components
	var ok bool

	scale := g.m.scale

	if cmpts, ok = scaled[scale]; !ok {
		return fmt.Errorf("Unrecognised scale value: %s", scale)
	}

	g.m.height = cmpts.height
	g.m.width = cmpts.width

	g.m.panels = map[windowColumn]map[displayType]windowSlice{
		0: {
			DSideWall:     cmpts.d0SideWall,
			DOpenWallNear: cmpts.d06OpenWall,
		},
		1: {
			DSideWall:       cmpts.d1SideWall,
			DOpenWallMiddle: cmpts.d15OpenWallMiddle,
			DOpenWallNear:   cmpts.d15OpenWallNear,
		},
		2: {
			DSideWall:       cmpts.d2SideWall,
			DOpenWallNear:   cmpts.d24OpenWallNear,
			DOpenWallMiddle: cmpts.d24OpenWallMiddle,
			DOpenWallFar:    cmpts.d24OpenWallFar,
		},
		3: {
			DEmpty:          cmpts.d3Empty,
			DOpenWallNear:   cmpts.d24OpenWallNear,
			DOpenWallMiddle: cmpts.d24OpenWallMiddle,
			DOpenWallFar:    cmpts.d24OpenWallFar,
		},
		4: {
			DSideWall:       cmpts.d4SideWall,
			DOpenWallNear:   cmpts.d24OpenWallNear,
			DOpenWallMiddle: cmpts.d24OpenWallMiddle,
			DOpenWallFar:    cmpts.d24OpenWallFar,
		},
		5: {
			DSideWall:       cmpts.d5SideWall,
			DOpenWallMiddle: cmpts.d15OpenWallMiddle,
			DOpenWallNear:   cmpts.d15OpenWallNear,
		},
		6: {
			DSideWall:     cmpts.d6SideWall,
			DOpenWallNear: cmpts.d06OpenWall,
		},
	}
	return nil
}
