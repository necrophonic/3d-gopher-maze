package element

// PanelDefinition defines the possible parts that can display
// for a panel depending on the current view
type PanelDefinition map[string]PixelMatrix

// Panels contains definitions for screen view display panels.
// These are first indexed by column number, then panel type,
// where panel type denotes whether it's a wall or an open wall etc
var Panels = map[int]PanelDefinition{
	0: {
		"SideWall":     d0SideWall,
		"OpenWallNear": d06OpenWall,
	},
	1: {
		"SideWall":       d1SideWall,
		"OpenWallMiddle": d15OpenWallMiddle,
		"OpenWallNear":   d15OpenWallNear,
	},
	2: {
		"SideWall":       d2SideWall,
		"OpenWallNear":   d24OpenWallNear,
		"OpenWallMiddle": d24OpenWallMiddle,
		"OpenWallFar":    d24OpenWallFar,
	},
	3: {
		"Empty":          d3Empty,
		"OpenWallNear":   d24OpenWallNear,
		"OpenWallMiddle": d24OpenWallMiddle,
		"OpenWallFar":    d24OpenWallFar,
	},
	4: {
		"SideWall":       d4SideWall,
		"OpenWallNear":   d24OpenWallNear,
		"OpenWallMiddle": d24OpenWallMiddle,
		"OpenWallFar":    d24OpenWallFar,
	},
	5: {
		"SideWall":       d5SideWall,
		"OpenWallMiddle": d15OpenWallMiddle,
		"OpenWallNear":   d15OpenWallNear,
	},
	6: {
		"SideWall":     d6SideWall,
		"OpenWallNear": d06OpenWall,
	},
}

var (
	d06OpenWall       = PixelMatrix{{T, T}, {O, O}, {O, O}, {O, O}, {O, O}, {O, O}, {O, O}, {O, O}, {T, T}}
	d15OpenWallMiddle = PixelMatrix{{T, T}, {T, T}, {T, T}, {O, O}, {O, O}, {O, O}, {T, T}, {T, T}, {T, T}}
	d15OpenWallNear   = PixelMatrix{{T, T}, {O, O}, {O, O}, {O, O}, {O, O}, {O, O}, {O, O}, {O, O}, {T, T}}
	d2SideWall        = PixelMatrix{{T}, {T}, {T}, {T}, {W}, {T}, {T}, {T}, {T}}
	d4SideWall        = PixelMatrix{{T}, {T}, {T}, {T}, {W}, {T}, {T}, {T}, {T}}
	d24OpenWallNear   = PixelMatrix{{T}, {O}, {O}, {O}, {O}, {O}, {O}, {O}, {T}}
	d24OpenWallMiddle = PixelMatrix{{T}, {T}, {T}, {O}, {O}, {O}, {T}, {T}, {T}}
	d24OpenWallFar    = PixelMatrix{{T}, {T}, {T}, {T}, {O}, {T}, {T}, {T}, {T}}
	d3Empty           = PixelMatrix{{T}, {T}, {T}, {T}, {T}, {T}, {T}, {T}, {T}}
	d0SideWall        = PixelMatrix{{W, T}, {W, W}, {W, W}, {W, W}, {W, W}, {W, W}, {W, W}, {W, W}, {W, T}}
	d1SideWall        = PixelMatrix{{T, T}, {T, T}, {W, T}, {W, W}, {W, W}, {W, W}, {W, T}, {T, T}, {T, T}}
	d5SideWall        = PixelMatrix{{T, T}, {T, T}, {T, W}, {W, W}, {W, W}, {W, W}, {T, W}, {T, T}, {T, T}}
	d6SideWall        = PixelMatrix{{T, W}, {W, W}, {W, W}, {W, W}, {W, W}, {W, W}, {W, W}, {W, W}, {T, W}}
)
