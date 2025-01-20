package messages

import (
	"fmt"
	"time"

	"github.com/r2unit/eyes/colours"
)

func getTimestamp() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func Colourize(color string, message string) string {
	return fmt.Sprintf("%s%s%s", color, message, colours.Reset)
}

var (
	Success = func(message string) string {
		return fmt.Sprintf("%s [%s]%s %s%s", getTimestamp(), Colourize(colours.Bold+colours.Green, "Success"), colours.Reset, message, colours.Reset)
	}
	Warning = func(message string) string {
		return fmt.Sprintf("%s [%s]%s %s%s", getTimestamp(), Colourize(colours.Bold+colours.Yellow, "Warning"), colours.Reset, message, colours.Reset)
	}
	Danger = func(message string) string {
		return fmt.Sprintf("%s [%s]%s %s%s", getTimestamp(), Colourize(colours.Bold+colours.Red, "Danger"), colours.Reset, message, colours.Reset)
	}
	Info = func(message string) string {
		return fmt.Sprintf("%s [%s]%s %s%s", getTimestamp(), Colourize(colours.Bold+colours.Blue, "Info"), colours.Reset, message, colours.Reset)
	}
	Debug = func(message string) string {
		return fmt.Sprintf("%s [%s]%s %s%s", getTimestamp(), Colourize(colours.Bold+colours.Cyan, "Debug"), colours.Reset, message, colours.Reset)
	}
	Notice = func(message string) string {
		return fmt.Sprintf("%s [%s]%s %s%s", getTimestamp(), Colourize(colours.Bold+colours.Magenta, "Notice"), colours.Reset, message, colours.Reset)
	}
)
