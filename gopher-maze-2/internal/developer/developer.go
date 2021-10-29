// Package developer contains various tools for helping develop the game
package developer

import (
	"fmt"
	"time"
)

type Stack interface {
	Add(string)
	Addf(string, ...string)
	Get(int) string
	Window() []string
}

// DebugStack implements a stack of debugging messages that can be rendered via
// a sliding window to the terminal. Should be instantiated via NewDebugStack()
type DebugStack struct {
	clean      bool
	size       uint8
	msgs       []string
	lastWindow []string
}

// NewDebugStack initialises an empty debug stack with its
// window set to the size of `windowSize`.
func NewDebugStack(windowSize uint8) *DebugStack {
	return &DebugStack{
		clean:      false,
		size:       windowSize,
		msgs:       make([]string, windowSize),
		lastWindow: make([]string, windowSize),
	}
}

// Add appends a new message to the debug stack.
func (ds *DebugStack) Add(msg string) {
	ds.msgs = append(ds.msgs, fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), msg))
	ds.clean = false
}

// Add appends a new message to the debug stack (formatted).
func (ds *DebugStack) Addf(format string, msg ...string) {
	ds.Add(fmt.Sprintf(format, msg))
}

// Get fetches the given indexed line from the debug stack
// (if it exists). If the index isn't populated or is out
// of range then an empty string is returned instead.
func (ds *DebugStack) Get(i int) string {
	window := ds.lastWindow
	if !ds.clean {
		// If the debug stack has changed
		// then rebuild the window.
		window = ds.Window()
	}
	if i >= len(window) {
		return ""
	}
	return window[i]
}

// Window returns a whole debug window
func (ds *DebugStack) Window() []string {
	upper := len(ds.msgs)
	lower := upper - int(ds.size)
	if lower < 0 {
		lower = 0
	}
	window := ds.msgs[lower:upper]
	ds.clean = true
	ds.lastWindow = window
	return window
}

// NullStack is used when we want to disable debugging stacks. It essentially
// returns no-ops for every func.
type NullStack struct{}

func (*NullStack) Add(string)             {}
func (*NullStack) Addf(string, ...string) {}
func (*NullStack) Get(int) string         { return "" }
func (*NullStack) Window() []string       { return nil }
