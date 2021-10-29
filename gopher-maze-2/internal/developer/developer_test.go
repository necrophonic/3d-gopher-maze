package developer_test

import (
	"testing"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/developer"
	"github.com/stretchr/testify/assert"
)

func TestDebugStack(t *testing.T) {

	var ds *developer.DebugStack

	stackSize := uint8(3)

	t.Run("Create a new debug stack", func(t *testing.T) {
		ds = developer.NewDebugStack(true, stackSize)
		assert.NotNil(t, ds)
	})

	t.Run("Test debug stack window", func(t *testing.T) {
		t.Run("Single message", func(t *testing.T) {
			ds.Add("Message 1")
			assert.Contains(t, ds.Get(2), "Message 1")
		})

		t.Run("Full window", func(t *testing.T) {
			ds.Add("Message 2")
			ds.Add("Message 3")
			assert.Contains(t, ds.Get(0), "Message 1")
			assert.Contains(t, ds.Get(1), "Message 2")
			assert.Contains(t, ds.Get(2), "Message 3")
		})

		t.Run("Roll window", func(t *testing.T) {
			ds.Add("Message 4")
			ds.Add("Message 5")
			assert.Contains(t, ds.Get(0), "Message 3")
			assert.Contains(t, ds.Get(1), "Message 4")
			assert.Contains(t, ds.Get(2), "Message 5")
		})

		t.Run("Stack overrun", func(t *testing.T) {
			assert.Equal(t, ds.Get(5), "")
		})
	})

	t.Run("Attempt fetch on empty stack", func(t *testing.T) {
		ds = developer.NewDebugStack(true, stackSize)
		assert.Equal(t, ds.Get(0), "")
	})

	t.Run("Attempt fetch on disbled stack", func(t *testing.T) {
		ds = developer.NewDebugStack(false, stackSize)
		assert.Equal(t, ds.Get(0), "")
	})

}
