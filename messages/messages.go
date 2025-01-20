package messages

import (
	"fmt"

	"github.com/r2unit/colours/colours" // Import the colours package
)

// Colourize wraps a message with the given color code.
func Colourize(color string, message string) string {
	return fmt.Sprintf("%s%s%s", color, message, colours.Reset)
}

// Predefined color functions with bold log level prefixes.
var (
	Success = func(message string) string {
		return fmt.Sprintf("%s[Success]%s %s", Colourize(colours.Bold+colours.Green, "Success"), colours.Reset, message)
	}
	Warning = func(message string) string {
		return fmt.Sprintf("%s[Warning]%s %s", Colourize(colours.Bold+colours.Yellow, "Warning"), colours.Reset, message)
	}
	Danger = func(message string) string {
		return fmt.Sprintf("%s[Danger]%s %s", Colourize(colours.Bold+colours.Red, "Danger"), colours.Reset, message)
	}
	Info = func(message string) string {
		return fmt.Sprintf("%s[Info]%s %s", Colourize(colours.Bold+colours.Blue, "Info"), colours.Reset, message)
	}
	Debug = func(message string) string {
		return fmt.Sprintf("%s[Debug]%s %s", Colourize(colours.Bold+colours.Cyan, "Debug"), colours.Reset, message)
	}
	Notice = func(message string) string {
		return fmt.Sprintf("%s[Notice]%s %s", Colourize(colours.Bold+colours.Magenta, "Notice"), colours.Reset, message)
	}
)
