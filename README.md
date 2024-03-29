# 3D Gopher Maze

[![Go Report Card](https://goreportcard.com/badge/github.com/necrophonic/3d-gopher-maze)](https://goreportcard.com/report/github.com/necrophonic/3d-gopher-maze)

Find the gopher, escape the maze!

Inspired by 80s games like 3d monster maze and phantom slayer, this is a little
maze crawler game using just the terminal for rendering.

You're placed randomly in a maze and you need to find the gopher and escape by
finding the exit.

```
╔════════════════════════╗
║ ▓▓                  ▓▓ ║
║ ▓▓▓▓              ▓▓▓▓ ║
║ ▓▓▓▓▓▓            ▓▓▓▓ ║
║ ▓▓▓▓▓▓▓▓      ░░░░▓▓▓▓ ║
║ ▓▓▓▓▓▓▓▓▓▓  ▓▓░░░░▓▓▓▓ ║
║ ▓▓▓▓▓▓▓▓      ░░░░▓▓▓▓ ║
║ ▓▓▓▓▓▓            ▓▓▓▓ ║
║ ▓▓▓▓              ▓▓▓▓ ║
║ ▓▓                  ▓▓ ║
╚════════════════════════╝
Facing: N

Which way?: 
```

## Playing

### Launching

Open your terminal of choice and run:

```shell
go run cmd/maze/main.go
```

There are a few environment vars supported for changing behaviours:

| Env var | Default | Description                                                                                        |
| ------- | ------- | -------------------------------------------------------------------------------------------------- |
| DEBUG   | `false` | Switches on developer debugging. This will stop the screen auto-refreshing and move to slide mode. |


### Controls

To move around, use `W`,`S`,`A` & `D` (case insensitve) keys. Quit using `q` and display help with `h`

| Key | Action                   |
| --- | ------------------------ |
| `w` | Move forward one space   |
| `s` | Move backwards one space |
| `a` | Rotate 90&deg; left      |
| `d` | Rotate 90&deg; right     |
| `q` | Quit game                |
| `h` | Display help             |

Currently the game doesn't support just press-move; you need to type the
letter of the command and hit `<return>`!

### How to win

> Note, the win conditions aren't actually implemented yet!

- Find the gopher hidden somewhere in the maze
- Then find the way out! (not implemented yet!)

### Exiting the game

Either type `q` as a command, or use `ctrl-c`


## Contributing

- Any pull requests raised must:
  - Be formated with `go fmt`
  - Contain a good description
  - Have sensible commit messages (not just "fixed things")
  - Where appropriate should include/update unit tests verifying the change/fix

- Any spam will be deleted

## License

Copyright &copy; 2020 J Gregory

Released under MIT license, see [LICENSE](LICENSE) for details.
