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
		ds = developer.NewDebugStack(stackSize)
		assert.NotNil(t, ds)
	})

	t.Run("Test debug stack window", func(t *testing.T) {
		t.Run("Single message", func(t *testing.T) {
			ds.AddMsg("Message 1")
			assert.Equal(t, "Message 1", ds.Get(2))
		})

		t.Run("Full window", func(t *testing.T) {
			ds.AddMsg("Message 2")
			ds.AddMsg("Message 3")
			assert.Equal(t, "Message 1", ds.Get(0))
			assert.Equal(t, "Message 2", ds.Get(1))
			assert.Equal(t, "Message 3", ds.Get(2))
		})

		t.Run("Roll window", func(t *testing.T) {
			ds.AddMsg("Message 4")
			ds.AddMsg("Message 5")
			assert.Equal(t, "Message 3", ds.Get(0))
			assert.Equal(t, "Message 4", ds.Get(1))
			assert.Equal(t, "Message 5", ds.Get(2))
		})

		t.Run("Stack overrun", func(t *testing.T) {
			assert.Equal(t, "", ds.Get(5))
		})
	})

	t.Run("Attempt fetch on empty stack", func(t *testing.T) {
		ds = developer.NewDebugStack(stackSize)
		assert.Equal(t, "", ds.Get(0))
	})

}
