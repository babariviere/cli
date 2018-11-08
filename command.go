package cli

// Command is a cli subcommand
type Command interface {
	// Parse is used to parse arguments
	//
	// Arguments will be those who follow the commands
	// e.j: go build main.go will return main.go for the build subcommand
	Parse([]string) error
	// Usage return the usage of the subcommand
	Usage() string
	// Spawn the subcommand
	Spawn() error
}
