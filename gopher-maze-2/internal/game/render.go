package game

import (
	"strings"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/terminal"
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

func (g *Game) getDebugMsg(i int) string {
	if !g.debug {
		return ""
	}
	return g.debugStack.Get(i)
}

// Render contructs the terminal view. If debug is enabled, the debug stack
// is rendered alongside the viewport.
func (g *Game) Render() (string, error) {
	// Build the output
	rendered := ""

	// if g.debug {
	// 	debugWindow := g.debugStack.Get()
	// }

	rendered += windowTopLeft + strings.Repeat(windowHoriz, windowWidth+2) + windowTopRight + "\n"
	for i := 0; i < windowHeight; i++ {
		// TODO - content!
		rendered += windowSide + " " + strings.Repeat(" ", windowWidth) + " " + windowSide + "   " + g.getDebugMsg(i) + "\n"
	}
	rendered += windowBottomLeft + strings.Repeat(windowHoriz, windowWidth+2) + windowBottomRight + "\n"

	return rendered, nil
}
