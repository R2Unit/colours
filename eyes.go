package eyes

import "github.com/r2unit/eyes/messages"

var Colours = struct {
	Success func(string)
	Warning func(string)
	Danger  func(string)
	Info    func(string)
	Debug   func(string)
	Notice  func(string)
}{
	Success: func(message string) { println(messages.Success(message)) },
	Warning: func(message string) { println(messages.Warning(message)) },
	Danger:  func(message string) { println(messages.Danger(message)) },
	Info:    func(message string) { println(messages.Info(message)) },
	Debug:   func(message string) { println(messages.Debug(message)) },
	Notice:  func(message string) { println(messages.Notice(message)) },
}
