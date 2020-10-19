package game

import "github.com/necrophonic/gopher-maze/internal/debug"

func (g *Game) moveForward() {
	if g.isMoveToWall(g.move.x, g.move.y) {
		debug.Println("Move forward to wall - stopping")
		return
	}
	g.p.x += g.move.x
	g.p.y += g.move.y
}

func (g *Game) moveBackwards() {
	if g.isMoveToWall(g.move.x*-1, g.move.y*-1) {
		debug.Println("Move backward to wall - stopping")
		return
	}
	g.p.x += (g.move.x * -1)
	g.p.y += (g.move.y * -1)
}

func (g *Game) rotateRight() {
	debug.Println("Rotate right")
	switch g.p.o {
	case 'n':
		g.p.o = 'e'
		g.move.x = 1
		g.move.y = 0
	case 'e':
		g.p.o = 's'
		g.move.x = 0
		g.move.y = 1
	case 's':
		g.p.o = 'w'
		g.move.x = -1
		g.move.y = 0
	case 'w':
		g.p.o = 'n'
		g.move.x = 0
		g.move.y = -1
	}
}

func (g *Game) rotateLeft() {
	debug.Println("Rotate left")
	switch g.p.o {
	case 'n':
		g.p.o = 'w'
		g.move.x = -1
		g.move.y = 0
	case 'w':
		g.p.o = 's'
		g.move.x = 0
		g.move.y = 1
	case 's':
		g.p.o = 'e'
		g.move.x = 1
		g.move.y = 0
	case 'e':
		g.p.o = 'n'
		g.move.x = 0
		g.move.y = -1
	}
}

func (g *Game) isMoveToWall(mx, my int8) bool {
	px := g.p.x + mx
	py := g.p.y + my
	if g.m.getSpace(point{px, py}).t == SpaceWall {
		return true
	}
	return false
}
