package main

import (
	"context"
	"runtime"

	"github.com/necrophonic/3d-gopher-maze/gopher-maze-2/internal/game"
)

func main() {
	ctx := context.Background()

	g := game.New(true)

	// Start rendering loop
	go g.Loop(ctx)

	// Start player input loop
	go g.Player.Loop()

	// Exit the main goroutine and leave
	// our main game loops running.
	runtime.Goexit()
}
