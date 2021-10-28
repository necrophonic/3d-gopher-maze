package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gookit/color"
	"github.com/mattn/go-tty"
	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/terminal"
)

var (
	// The tick is the interval between engine / screen updates
	tick = 100_000 * time.Microsecond
)

var message = "No message"

var (
	colorLightGrey = color.S256(254, 0).Sprint
)

func main() {

	os.Exit(0)

	// Start rendering loop
	i := 0
	go func() {
		for {
			render(i)
			time.Sleep(tick)
			i++
		}
	}()

	// Start player input loop
	go func() {
		t, err := tty.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer t.Close()
		for {
			r, err := t.ReadRune()
			if err != nil {
				log.Fatal(err)
			}
			message = "Keypress: " + string(r)

			switch r {
			case 'q':
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
	}()

	// Exit the main goroutine and leave
	// our main game loop running.
	runtime.Goexit()
}

func render(i int) {
	terminal.Clear()
	index := strconv.Itoa(i)
	fmt.Println(`╔════════════════════════╗
║ ▓▓                  ▓▓ ║
║ ▓▓▓▓              ▓▓▓▓ ║
║ ▓▓▓▓▓▓            ` + colorLightGrey("▓▓▓▓") + ` ║
║ ▓▓▓▓▓▓▓▓      ░░░░▓▓▓▓ ║
║ ▓▓▓▓▓▓▓▓▓▓  ▓▓░░░░▓▓▓▓ ║
║ ▓▓▓▓▓▓▓▓      ░░░░▓▓▓▓ ║
║ ▓▓▓▓▓▓            ▓▓▓▓ ║
║ ▓▓▓▓              ▓▓▓▓ ║
║ ▓▓                  ▓▓ ║
╚════════════════════════╝ ` + index + "\n" + message)
}
