package game

type gameState int

const (
	stateInitialising gameState = iota
	stateRunning
	stateExiting
)

type gameStateService struct {
	currentState gameState
}

var state gameState = stateInitialising

func newGameStateService() *gameStateService {
	return &gameStateService{
		currentState: stateInitialising,
	}
}

func (s *gameStateService) HandleInput(e Event) {
	if e == KeyEsc {
		s.currentState = stateExiting
	}
}
