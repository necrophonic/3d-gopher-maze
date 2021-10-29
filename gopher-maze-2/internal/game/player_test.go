package game_test

import (
	"testing"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	var p *game.Player
	t.Run("Create new player", func(t *testing.T) {
		p = game.NewPlayer()
		assert.NotNil(t, p)
	})

	t.Run("Rotate and move (no walls)", func(t *testing.T) {
		p = game.NewPlayer()
		assert.Equal(t, "Player N @{0,0}", p.String())
		p.RotateLeft()
		assert.Equal(t, "Player W @{0,0}", p.String())
		p.RotateLeft()
		assert.Equal(t, "Player S @{0,0}", p.String())
		p.RotateLeft()
		assert.Equal(t, "Player E @{0,0}", p.String())
		p.RotateRight()
		assert.Equal(t, "Player S @{0,0}", p.String())
		p.MoveForward()
		assert.Equal(t, "Player S @{0,1}", p.String())
		p.MoveForward()
		assert.Equal(t, "Player S @{0,2}", p.String())
		p.RotateLeft()
		p.MoveForward()
		assert.Equal(t, "Player E @{1,2}", p.String())
		p.MoveBackward()
		assert.Equal(t, "Player E @{0,2}", p.String())

	})
}
