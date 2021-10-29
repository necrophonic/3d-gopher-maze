package game

import (
	"strings"
)

const (
	windowHoriz       = "═"
	windowTopLeft     = "╔"
	windowTopRight    = "╗"
	windowBottomLeft  = "╚"
	windowBottomRight = "╝"
	windowSide        = "║"
)

const (
	windowWidth  = 22
	windowHeight = 9
)

// Render contructs the terminal view. If debug is enabled, the debug stack
// is rendered alongside the viewport.
func (g *Game) Render() (string, error) {
	// Build the output
	rendered := ""

	rendered += windowTopLeft + strings.Repeat(windowHoriz, windowWidth+2) + windowTopRight + "   " + ds.Get(0) + "\n"
	for i := 0; i < windowHeight; i++ {
		// TODO - content!
		rendered += windowSide + " " + strings.Repeat(" ", windowWidth) + " " + windowSide + "   " + ds.Get(i+1) + "\n"
	}
	rendered += windowBottomLeft + strings.Repeat(windowHoriz, windowWidth+2) + windowBottomRight + "   " + ds.Get(windowHeight+1) + "\n"

	return rendered, nil
}
