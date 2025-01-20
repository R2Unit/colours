package main

import (
	"fmt"

	"github.com/r2unit/colours/messages" // Import the messages package
)

func main() {
	fmt.Println(messages.Success("Operation was successful."))
	fmt.Println(messages.Warning("This is a warning message."))
	fmt.Println(messages.Danger("Something went wrong!"))
	fmt.Println(messages.Info("Here is some informational text."))
	fmt.Println(messages.Debug("This is a debug message."))
	fmt.Println(messages.Notice("This is a notice message."))
}
