// Package developer contains various tools for helping develop the game
package developer

import (
	"fmt"
	"log"
	"time"
)

// DebugStack implements a stack of debugging messages that can be rendered via
// a sliding window to the terminal. Should be instantiated via NewDebugStack()
type DebugStack struct {
	size uint8
	msgs []string
}

// NewDebugStack initialises an empty debug stack with its
// window set to the size of `windowSize`.
func NewDebugStack(enable bool, windowSize uint8) *DebugStack {
	return &DebugStack{
		size: windowSize,
		msgs: make([]string, windowSize),
	}
}

// Add appends a new message to the debug stack. The message is also echoed to
// the STDOUT log.
func (ds *DebugStack) Add(msg string) {
	ds.msgs = append(ds.msgs, fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), msg))
	log.Println(msg)
}

// Add appends a new message to the debug stack (formatted). The message is
// also echoed to the STDOUT log.
func (ds *DebugStack) Addf(format string, msg ...string) {
	ds.Add(fmt.Sprintf(format, msg))
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
