package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gookit/color"
	"github.com/mattn/go-tty"
	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/game"
)

// var (
// 	// The tick is the interval between engine / screen updates
// 	tick = 100_000 * time.Microsecond
// )

// var message = "No message"

var (
	colorLightGrey = color.S256(254, 0).Sprint
)

func main() {

	gme := game.New(true)

	// Start rendering loop
	go gme.Loop()

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
			gme.Debug("Keypress: " + string(r))

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
