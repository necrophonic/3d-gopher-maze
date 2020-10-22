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
	T = element.T
	O = element.O
	W = element.W
	G = element.G
)

// GetMatrix returns an overlay for the given distance
func (g *gopher) GetMatrix(distance int) (element.PixelMatrix, error) {
	if distance < 1 || distance > 3 {
		return nil, fmt.Errorf("cannot render gopher at distance (%d)", distance)
	}

	debug.Printf("Returning gopher overlay for distance '%d'", distance)

	return map[int]element.PixelMatrix{
		3: {
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, G, T, T, T, T, T},
			{T, T, T, T, T, G, T, T, T, T, T},
			{T, T, T, T, T, G, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
		},
		2: {
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, G, T, G, T, T, T, T},
			{T, T, T, T, G, G, G, T, T, T, T},
			{T, T, T, T, G, G, G, T, T, T, T},
			{T, T, T, T, G, G, G, T, T, T, T},
			{T, T, T, T, G, T, G, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
		},
		1: {
			{T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, G, T, T, T, G, T, T, T},
			{T, T, T, G, G, G, G, G, T, T, T},
			{T, T, T, G, G, G, G, G, T, T, T},
			{T, T, T, T, G, G, G, T, T, T, T},
			{T, T, T, G, G, G, G, G, T, T, T},
			{T, T, T, G, G, G, G, G, T, T, T},
			{T, T, T, T, G, T, G, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T},
		},
	}[distance], nil
}
