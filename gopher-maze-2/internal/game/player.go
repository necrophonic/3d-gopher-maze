package game

import (
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-tty"
)

const (
	N orientation = 0
	E orientation = 1
	S orientation = 2
	W orientation = 3
)

var directions = []orientation{N, E, S, W}

type (
	// Player represents the current player state
	Player struct {
		Position    Point
		Orientation orientation
	}
)

func NewPlayer() *Player {
	plyr := &Player{
		Orientation: N,
		Position:    Point{0, 0}, // TODO - starting position based on map
	}
	ds.Addf("[added player at %s]", plyr.String())
	return plyr
}

func (p *Player) String() string {
	return fmt.Sprintf("Player %s @{%d,%d}", p.Orientation.MustString(), p.Position.X, p.Position.Y)
}

func (p *Player) Loop() {
	// Need to open a tty to be able
	// to capture the user input.
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()
	for {
		r, err := t.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		ds.Addf("Keypress [%s]", string(r))

		switch r {
		case 'w':
			p.MoveForward()
		case 's':
			p.MoveBackward()
		case 'a':
			p.RotateLeft()
		case 'd':
			p.RotateRight()
		case 'q':
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

// RotateLeft rotates the player's orientation to the left (counter-clockwise)
func (p *Player) RotateLeft() {
	p.Orientation = directions[(p.Orientation-1+4)%4]
	ds.Addf("Rotate left: %s", p.String())
}

// RotateLeft rotates the player's orientation to the right (clockwise)
func (p *Player) RotateRight() {
	p.Orientation = directions[(p.Orientation+1+4)%4]
	ds.Addf("Rotate right: %s", p.String())
}

func (p *Player) move(modifier int) {
	switch p.Orientation {
	case N:
		p.Position.Y -= modifier
	case S:
		p.Position.Y += modifier
	case E:
		p.Position.X += modifier
	case W:
		p.Position.X -= modifier
	}
}

// MoveForward moves the player backwards 1 step
func (p *Player) MoveForward() {
	p.move(1)
	ds.Addf("Forward: %s", p.String())
}

// MoveBackward moves the player backwards 1 step
func (p *Player) MoveBackward() {
	p.move(-1)
	ds.Addf("Backward: %s", p.String())
}

type orientation byte

func (o orientation) MustString() string {
	switch o {
	case 0:
		return "N"
	case 1:
		return "E"
	case 2:
		return "S"
	case 3:
		return "W"
	}
	panic("Unable to parse orientation to string - incorrect range?")
}
