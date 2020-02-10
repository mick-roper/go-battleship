package game

import (
	"bufio"
	"os"
)

// Known keys
const (
	KeyLeft Event = iota
	KeyRight
	KeyUp
	KeyDown
	KeySpace
	KeyEsc
)

// Event from the user
type Event byte

// InputService gets user input
type InputService interface {
	GetEvents() (bool, Event)
}

// ConsoleInputService gets input from stdin
type ConsoleInputService struct {
	reader *bufio.Reader
}

// NewConsoleInputService creates a new instance of the ConsoleInputService
func NewConsoleInputService() *ConsoleInputService {
	return &ConsoleInputService{
		reader: bufio.NewReader(os.Stdin),
	}
}

// GetEvents from STD in
func (c *ConsoleInputService) GetEvents() (bool, Event) {
	b, err := c.reader.ReadByte()

	if err != nil {
		return false, 0
	}

	switch b {
	case 'j':
		return true, KeyLeft
	case 'k':
		return true, KeyDown
	case 'i':
		return true, KeyDown
	case 'l':
		return true, KeyRight
	case 'o':
		return true, KeySpace
	case 'x':
		return true, KeyEsc
	default:
		return false, 0
	}
}
