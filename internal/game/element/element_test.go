package element_test

import (
	"testing"

	"github.com/necrophonic/gopher-maze/internal/game/element"
	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	W := element.W
	X := element.X

	pm := element.PixelMatrix{
		{W, W, W, W, W},
		{W, W, W, W, W},
		{W, W, W, W}, // Intentionally not the same size!
	}

	pm.Clear()

	expected := element.PixelMatrix{
		{X, X, X, X, X},
		{X, X, X, X, X},
		{X, X, X, X},
	}

	assert.ElementsMatch(t, expected, pm)
}
