package game

import (
	"bytes"
	"text/template"
)

// const (
// 	windowHoriz       = "═"
// 	windowTopLeft     = "╔"
// 	windowTopRight    = "╗"
// 	windowBottomLeft  = "╚"
// 	windowBottomRight = "╝"
// 	windowSide        = "║"
// )

// const (
// 	windowWidth  = 22
// 	windowHeight = 9
// )

// Render contructs the terminal view. If debug is enabled, the debug stack
// is rendered alongside the viewport.
func (g *Game) Render() (string, error) {
	// Build the output
	// rendered := ""

	// rendered += windowTopLeft + strings.Repeat(windowHoriz, windowWidth+2) + windowTopRight + "   " + ds.Get(0) + "\n"
	// for i := 0; i < windowHeight; i++ {
	// 	// TODO - content!
	// 	rendered += windowSide + " " + strings.Repeat(" ", windowWidth) + " " + windowSide + "   " + ds.Get(i+1) + "\n"
	// }
	// rendered += windowBottomLeft + strings.Repeat(windowHoriz, windowWidth+2) + windowBottomRight + "   " + ds.Get(windowHeight+1) + "\n"

	buf := &bytes.Buffer{}

	err := tmplScreen11_9.Execute(buf, nil)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

var funcs = template.FuncMap{
	"getDebugLine": func(i int) string {
		return ds.Get(i)
	},
}

var tmplScreen11_9 = template.Must(template.New("screen").Funcs(funcs).Parse(`
  ╔════════════════════════╗   {{ getDebugLine 0 }}
  ║                        ║   {{ getDebugLine 1 }}
  ║                        ║   {{ getDebugLine 2 }}
  ║                        ║   {{ getDebugLine 3 }}
  ║                        ║   {{ getDebugLine 4 }}
  ║                        ║   {{ getDebugLine 5 }}
  ║                        ║   {{ getDebugLine 6 }}
  ║                        ║   {{ getDebugLine 7 }}
  ║                        ║   {{ getDebugLine 8 }}
  ║                        ║   {{ getDebugLine 9 }}
  ╚════════════════════════╝   {{ getDebugLine 10 }}
`))
