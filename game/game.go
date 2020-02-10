package game

// Logger that writes logs
type Logger func(string)

// RunOptions used to define how the game works
type RunOptions struct {
	Logger Logger
}

// Run the game
func Run(opts *RunOptions) {

}
