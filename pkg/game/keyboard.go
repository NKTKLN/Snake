package game

import (
	"time"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
)

// This function is responsible for controlling the falling figure with the keyboard
func (d *Data) ReadingKeyboard() {
	for {
		// Reading data from the keyboard
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// Converting a pressed key to a command
			command, exists := models.ControlKeys[ev.Key]
			if !exists {
				command, exists = models.ControlCharacters[string(ev.Ch)]
				if !exists {
					continue
				}
			}

			// Command execution
			switch command {
			case "quit":
				d.EndStats(false, false)
			default:
				if command == models.MovementLock[d.Snake.Orientation] {
					continue
				}

				d.Snake.Orientation = command
				time.Sleep(100 * time.Microsecond)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
