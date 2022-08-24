package game

import "github.com/necrophonic/gopher-maze/internal/debug"

// Player represents the current player state
type Player struct {
	p Point
	o byte
}

// NewPlayer returns a new player populated with a positional
// point and orientation
func NewPlayer(p Point, o byte) *Player {
	return &Player{
		p: p,
		o: o,
	}
}

// FrontPoint returns the point directly in front of the player
// taking into account the direction they are facing.
func (p *Player) FrontPoint() Point {
	switch p.o {
	case 'n':
		return Point{p.p.X, p.p.Y - 1}
	case 's':
		return Point{p.p.X, p.p.Y + 1}
	case 'e':
		return Point{p.p.X + 1, p.p.Y}
	case 'w':
		return Point{p.p.X - 1, p.p.Y}
	}
	return Point{}
}

func (g *Game) moveForward() {
	if g.isMoveToWall(g.move.x, g.move.y) {
		g.Msg = "Can't go that way!"
		debug.Println("Move forward to wall - stopping")
		return
	}
	g.player.p.X += g.move.x
	g.player.p.Y += g.move.y

	if g.gopher.p.Is(g.player.FrontPoint()) {
		g.state = sWin
	}
	g.stats.moves++
}

func (g *Game) moveBackwards() {
	if g.isMoveToWall(g.move.x*-1, g.move.y*-1) {
		g.Msg = "Can't go that way!"
		debug.Println("Move backward to wall - stopping")
		return
	}
	g.player.p.X += (g.move.x * -1)
	g.player.p.Y += (g.move.y * -1)
	g.stats.moves++
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
	px := g.player.p.X + mx
	py := g.player.p.Y + my
	return g.m.getSpace(Point{px, py}).t == SpaceWall
}
