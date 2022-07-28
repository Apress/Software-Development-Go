package main

import (
	"github.com/gookit/color"
)

func main() {
	color.Warn = &color.Theme{"warning", color.Style{color.BgDefault, color.FgWhite}}
	color.Error = &color.Theme{"error", color.Style{color.BgDefault, color.FgRed}}

	color.Style{color.FgGreen, color.BgWhite, color.OpItalic}.Println("Italing style")
	color.Style{color.FgDefault, color.BgDefault, color.OpStrikethrough}.Println("Strikethrough style")
	color.Style{color.FgDefault, color.BgDefault, color.OpBold}.Println("Bold style")

	color.Warn.Prompt("Warning message")
	color.Error.Prompt("Error message")
}
