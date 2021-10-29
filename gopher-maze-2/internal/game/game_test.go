package game_test

import (
	"testing"
	"time"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	t.Run("Create new game", func(t *testing.T) {
		g := game.New()
		assert.NotNil(t, g)
		assert.Equal(t, 100*time.Millisecond, g.Tick, "default tick 100ms")
		assert.NotNil(t, g.Player, "created default player")
	})
}
