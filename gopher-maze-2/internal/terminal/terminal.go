//go:build !windows
// +build !windows

package terminal

import (
	"os"
	"os/exec"
)

// Clear clears the terminal
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
