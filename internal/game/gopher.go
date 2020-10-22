package game

import (
	"fmt"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/necrophonic/gopher-maze/internal/game/element"
)

type gopher struct {
	p point
}

func (g *gopher) GetPoint() point {
	return g.p
}

// ╔════════════════════════╗
// ║ ▓▓                  ▓▓ ║
// ║ ▓▓▓▓              ▓▓▓▓ ║
// ║ ▓▓▓▓▓▓            ▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓  ░░  ░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓▓ ░░ ▓░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓  ░░  ░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓            ▓▓▓▓ ║
// ║ ▓▓▓▓              ▓▓▓▓ ║
// ║ ▓▓                  ▓▓ ║
// ╚════════════════════════╝
//
// ╔════════════════════════╗
// ║ ▓▓                  ▓▓ ║
// ║ ▓▓▓▓              ▓▓▓▓ ║
// ║ ▓▓▓▓▓   ░    ░    ▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓ ░░░░░░ ░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓ ░░░░ ░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓ ░░░░░░ ░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓    ░  ░     ▓▓▓▓ ║
// ║ ▓▓▓▓              ▓▓▓▓ ║
// ║ ▓▓                  ▓▓ ║
// ╚════════════════════════╝
//
// ╔════════════════════════╗
// ║ ▓▓                  ▓▓ ║
// ║ ▓▓▓▓  ░░      ░░  ▓▓▓▓ ║
// ║ ▓▓▓▓▓ ░░▓▓▓▓▓▓░░  ▓▓▓▓ ║
// ║ ▓▓▓▓▓ ░░▓ ▓▓ ▓░░ ░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓ ░░░░░░░░ ░░▓▓▓▓ ║
// ║ ▓▓▓▓▓ ░░▓▓░░▓▓░░ ░▓▓▓▓ ║
// ║ ▓▓▓▓▓ ░░░░░░░░░░  ▓▓▓▓ ║
// ║ ▓▓▓▓    ░░  ░░    ▓▓▓▓ ║
// ║ ▓▓                  ▓▓ ║
// ╚════════════════════════╝

// Shorthand - see element package for full definitions
const (
	T  = element.T
	O  = element.O
	W  = element.W
	G  = element.G
	LE = element.LE
	RE = element.RE
	LI = element.LI
	RI = element.RI
	LO = element.LO
	RO = element.RO
	GM = element.GM
)

var gopherFrames = map[int]element.PixelMatrix{
	4: {
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, LO, G, RO, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
	},
	3: {
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, LO, G, RO, T, T, T, T},
		{T, T, T, T, LO, G, RO, T, T, T, T},
		{T, T, T, T, LO, G, RO, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
	},
	2: {
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, LO, G, T, G, RO, T, T, T},
		{T, T, T, LO, G, G, G, RO, T, T, T},
		{T, T, T, T, LI, G, RI, T, T, T, T},
		{T, T, T, LO, G, G, G, RO, T, T, T},
		{T, T, T, T, LI, T, RI, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
	},
	1: {
		{T, T, T, T, T, T, T, T, T, T, T},
		{T, T, LO, LI, G, T, G, RI, RO, T, T},
		{T, T, LO, G, G, G, G, G, RO, T, T},
		{T, T, LO, G, LE, G, RE, G, RO, T, T},
		{T, T, T, LI, G, G, G, RI, T, T, T},
		{T, T, LO, G, GM, GM, GM, G, RO, T, T},
		{T, T, LO, G, GM, GM, GM, G, RO, T, T},
		{T, T, T, LO, G, T, G, RO, T, T, T},
		{T, T, T, T, T, T, T, T, T, T, T},
	},
}

// GetMatrix returns an overlay for the given distance
func (g *gopher) GetMatrix(distance int) (element.PixelMatrix, error) {
	if distance < 1 || distance > len(gopherFrames) {
		return nil, fmt.Errorf("cannot render gopher at distance (%d)", distance)
	}

	debug.Printf("Returning gopher overlay for distance '%d'", distance)

	return gopherFrames[distance], nil
}
