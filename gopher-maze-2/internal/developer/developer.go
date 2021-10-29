// Package developer contains various tools for helping develop the game
package developer

import (
	"fmt"
	"time"
)

type DebugStack struct {
	size uint8
	msgs []string
}

// NewDebugStack initialises an empty debug stack with its
// window set to the size of `windowSize`.
func NewDebugStack(windowSize uint8) *DebugStack {
	return &DebugStack{
		size: windowSize,
		msgs: make([]string, windowSize), // TODO size
	}
}

// AddMsg appends a new message to the debug stack.
func (ds *DebugStack) AddMsg(msg string) {
	ds.msgs = append(ds.msgs, fmt.Sprintf("%s: %s", time.Now().Format("15:04:05"), msg))
}

// Get fetches the given indexed line from the debug stack
// (if it exists). If the index isn't populated or is out
// of range then an empty string is returned instead.
func (ds *DebugStack) Get(i int) string {
	upper := len(ds.msgs)
	lower := upper - int(ds.size)
	if lower < 0 {
		lower = 0
	}
	window := ds.msgs[lower:upper]
	if i >= len(window) {
		return ""
	}
	return window[i]
}
