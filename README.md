# Colours
A Library in GoLang to add information to your logs based on colours &amp; output colours.

## Features

- **Customizable ANSI Colors**: Use predefined ANSI codes for colored output.
- **Predefined Log Levels**: Styled log levels like `Success`, `Warning`, `Danger`, `Info`, `Debug`, and `Notice`.
- **Modular Design**: Colors and messages are separated into distinct packages for reusability.

### Usage

```go
import (
    "fmt"
    "github.com/r2unit/colours"
)
```

## Example

```go
package main

import (
    "fmt"
    "github.com/r2unit/colours"
)

func main() {
    // Using colors
    fmt.Println(colours.Green + "This is a green message." + colours.Reset)

    // Using predefined messages
    fmt.Println(colours.Success("Operation was successful."))
    fmt.Println(colours.Warning("This is a warning message."))
    fmt.Println(colours.Danger("Something went wrong!"))
    fmt.Println(colours.Info("Here is some informational text."))
    fmt.Println(colours.Debug("This is a debug message."))
    fmt.Println(colours.Notice("This is a notice message."))
}
```

### Output
When you run the example, the output will look something like this (formatted with colors):
```text
This is a green message.
[Success] Operation was successful.
[Warning] This is a warning message.
[Danger] Something went wrong!
[Info] Here is some informational text.
[Debug] This is a debug message.
[Notice] This is a notice message.
```

## Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue if you encounter any bugs or have feature requests.
