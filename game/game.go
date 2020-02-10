package game

type gameState int

const (
	stateInitialising gameState = iota
	stateRunning
	stateExiting
)

// Logger that writes logs
type Logger func(string)

// RunOptions used to define how the game works
type RunOptions struct {
	Logger Logger
}

var state gameState = stateInitialising

// Run the game
func Run(opts *RunOptions) {
	opts.Logger("initialising the game")
	defer cleanup()

	for state != stateExiting {
		state = stateExiting
	}

	opts.Logger("exiting...")
}

func cleanup() {

}
