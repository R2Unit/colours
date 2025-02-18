package messages

import (
	"fmt"

	"github.com/r2unit/go-colours/colours"
)

func Colourize(color string, message string) string {
	return fmt.Sprintf("%s%s%s", color, message, colours.Reset)
}

var (
	Success = func(message string) string {
		return fmt.Sprintf("%s[Success]%s %s%s", colours.Bold+colours.Green, colours.Reset, message, colours.Reset)
	}
	Warning = func(message string) string {
		return fmt.Sprintf("%s[Warning]%s %s%s", colours.Bold+colours.Yellow, colours.Reset, message, colours.Reset)
	}
	Error = func(message string) string {
		return fmt.Sprintf("%s[Error]%s %s%s", colours.Bold+colours.Red, colours.Reset, message, colours.Reset)
	}
	Info = func(message string) string {
		return fmt.Sprintf("%s[Info]%s %s%s", colours.Bold+colours.Blue, colours.Reset, message, colours.Reset)
	}
	Debug = func(message string) string {
		return fmt.Sprintf("%s[Debug]%s %s%s", colours.Bold+colours.Cyan, colours.Reset, message, colours.Reset)
	}
	Notice = func(message string) string {
		return fmt.Sprintf("%s[Notice]%s %s%s", colours.Bold+colours.Magenta, colours.Reset, message, colours.Reset)
	}
	Ok = func(message string) string {
		return fmt.Sprintf("%s[OK]%s %s%s", colours.Bold+colours.Green, colours.Reset, message, colours.Reset)
	}
)
