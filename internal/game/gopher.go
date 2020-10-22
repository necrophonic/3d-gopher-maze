package game

import (
	"fmt"

	"github.com/necrophonic/gopher-maze/internal/game/element"
)

// type gopherElements struct {
// 	body uint8
// }

type gopher struct {
	p point
	// elements gopherElements
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

// func (g *Game) NewGopher(p point) *gopher {
// 	return *gopher{
// 		elements: {
// 			body: ''
// 		}
// 	}
// }

// Shorthand
const (
	T = element.T
	O = element.O
	W = element.W
)

// TODO Probably needs to support scaling if that actually gets done!
func (g *gopher) sprite(distance int) (element.PixelMatrix, error) {
	if distance < 1 || distance > 3 {
		return nil, fmt.Errorf("cannot render gopher at distance (%d)", distance)
	}

	return map[int]element.PixelMatrix{
		1: {
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, W, W, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, W, W, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, W, W, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
			{T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T, T},
		},
		// 2: {
		// 	"                      ",
		// 	"                      ",
		// 	"        ░    ░        ",
		// 	"        ░░░░░░        ",
		// 	"         ░░░░         ",
		// 	"        ░░░░░░        ",
		// 	"         ░  ░",
		// },
		// 3: {
		// 	"                      ",
		// 	"      ░░      ░░      ",
		// 	"      ░░▓▓▓▓▓▓░░      ",
		// 	"      ░░▓ ▓▓ ▓░░      ",
		// 	"       ░░░░░░░░       ",
		// 	"      ░░▓▓░░▓▓░░      ",
		// 	"      ░░░░░░░░░░      ",
		// 	"        ░░  ░░",
		// },
	}[distance], nil
}
