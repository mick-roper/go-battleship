package game

// Logger that writes logs
type Logger func(string)

// RunOptions used to define how the game works
type RunOptions struct {
	Logger           Logger
	InputService     InputService
	RenderingService RenderingService
}

// RenderingService handles rendering
type RenderingService interface {
	Clear()
	Draw(x, y int, character byte)
}

// Run the game
func Run(opts *RunOptions) {
	opts.Logger("initialising the game")
	defer opts.Logger("exiting...")

	stateService := newGameStateService()
	inputHandlers := []InputHandler{stateService}

	for stateService.currentState != stateExiting {
		for ok, e := opts.InputService.GetEvents(); ok == true; {
			for _, h := range inputHandlers {
				h.HandleInput(e)
			}
		}
	}
}
