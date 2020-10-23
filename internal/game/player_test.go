package game_test

import (
	"testing"

	"github.com/necrophonic/gopher-maze/internal/game"
	"github.com/stretchr/testify/assert"
)

func TestPlayerFrontPoint(t *testing.T) {

	type tc struct {
		facing             uint8
		playerPoint        game.Point
		expectedFrontPoint game.Point
	}

	cases := []tc{
		{'n', game.NewPointInt(12, 17), game.NewPointInt(12, 16)},
		{'s', game.NewPointInt(12, 17), game.NewPointInt(12, 18)},
		{'e', game.NewPointInt(12, 17), game.NewPointInt(13, 17)},
		{'w', game.NewPointInt(12, 17), game.NewPointInt(11, 17)},
	}

	for _, test := range cases {
		p := game.NewPlayer(test.playerPoint, test.facing)
		assert.Equal(t, test.expectedFrontPoint, p.FrontPoint())
	}

}
