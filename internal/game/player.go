package game

import "github.com/necrophonic/gopher-maze/internal/debug"

// Player represents the current player state
type Player struct {
	p point
	// x int8
	// y int8
	o byte
}

func (g *Game) moveForward() {
	if g.isMoveToWall(g.move.x, g.move.y) {
		debug.Println("Move forward to wall - stopping")
		return
	}
	g.player.p.x += g.move.x
	g.player.p.y += g.move.y
}

func (g *Game) moveBackwards() {
	if g.isMoveToWall(g.move.x*-1, g.move.y*-1) {
		debug.Println("Move backward to wall - stopping")
		return
	}
	g.player.p.x += (g.move.x * -1)
	g.player.p.y += (g.move.y * -1)
}

func (g *Game) rotateRight() {
	debug.Println("Rotate right")
	switch g.player.o {
	case 'n':
		g.player.o = 'e'
		g.move.x = 1
		g.move.y = 0
	case 'e':
		g.player.o = 's'
		g.move.x = 0
		g.move.y = 1
	case 's':
		g.player.o = 'w'
		g.move.x = -1
		g.move.y = 0
	case 'w':
		g.player.o = 'n'
		g.move.x = 0
		g.move.y = -1
	}
}

func (g *Game) rotateLeft() {
	debug.Println("Rotate left")
	switch g.player.o {
	case 'n':
		g.player.o = 'w'
		g.move.x = -1
		g.move.y = 0
	case 'w':
		g.player.o = 's'
		g.move.x = 0
		g.move.y = 1
	case 's':
		g.player.o = 'e'
		g.move.x = 1
		g.move.y = 0
	case 'e':
		g.player.o = 'n'
		g.move.x = 0
		g.move.y = -1
	}
}

func (g *Game) isMoveToWall(mx, my int8) bool {
	px := g.player.p.x + mx
	py := g.player.p.y + my
	if g.m.getSpace(point{px, py}).t == SpaceWall {
		return true
	}
	return false
}
