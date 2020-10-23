package terminal

import (
	"os"
	"os/exec"
)

// Clear clears the terminal
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
