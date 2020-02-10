package game

// Logger that writes logs
type Logger func(string)

// RunOptions used to define how the game works
type RunOptions struct {
	Logger           Logger
	InputService     InputService
	RenderingService RenderingService
}

// InputService provides input
type InputService interface {
	GetEvents()
}

// RenderingService handles rendering
type RenderingService interface {
	Clear()
	Draw(x, y int, character byte)
}

type gameState int

const (
	stateInitialising gameState = iota
	stateRunning
	stateExiting
)

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
