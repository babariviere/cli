package cli

import (
	"fmt"
	"os"
)

// App is the representation of a cli App with all subcommands
type App struct {
	commands map[string]Command
	aliases  map[string]string
	usage    string
}

// NewApp creates a new app
func NewApp(usage string) *App {
	return &App{
		commands: make(map[string]Command),
		aliases:  make(map[string]string),
		usage:    usage,
	}
}

// Main parse arguments and executes subcommands
//
// e.j: app.Main(os.Args[1:])
func (a App) Main(args []string) error {
	cmd := args[0]
	if val, ok := a.aliases[cmd]; ok {
		cmd = val
	}

	if val, ok := a.commands[cmd]; ok {
		if err := val.Parse(args[1:]); err != nil {
			fmt.Println(err)
			fmt.Println(val.Usage())
			os.Exit(1)
		}

		if err := val.Spawn(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if cmd == "help" {
		if len(os.Args) > 1 {
			cmd := args[1]
			if val, ok := a.aliases[cmd]; ok {
				cmd = val
			}
			if val, ok := a.commands[cmd]; ok {
				fmt.Println(val.Usage())
			} else {
				fmt.Println(a.usage)
			}
		} else {
			fmt.Println(a.usage)
		}
	} else {
		fmt.Println(a.usage)
		os.Exit(1)
	}

	return nil
}

// Usage returns app usage
func (a App) Usage() string {
	return a.usage
}

// RegisterCommand registers a command into the app
func (a *App) RegisterCommand(name string, cmd Command) {
	a.commands[name] = cmd
}

// RegisterAlias registers an aliases
func (a *App) RegisterAlias(cmd string, alias string) {
	a.aliases[alias] = cmd
}
