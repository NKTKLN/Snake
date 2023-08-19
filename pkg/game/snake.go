package game

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
)

// This function is responsible for creating the snake at the beginning of the game
func (d *Data) CreateSnake() {
	// Receiving terminal width and height
	terminalWidth, terminalHeight := termbox.Size()

	// Creating three snake segments
	for count := 0; count < 3; count++ {
		d.Snake.Body = append(d.Snake.Body, models.Position{XCord: terminalWidth / 2, YCord: terminalHeight / 2})
	}

	// Setting the random orientation of the snake's movement
	orientationIndex := rand.Intn(len(models.MovementTypes))

	d.Snake.Orientation = models.MovementTypes[orientationIndex]
}

// This function is responsible for moving the snake around the field
func (d *Data) MoveSnake() {
	for {
		start := time.Now()

		// Receiving terminal width and height
		terminalWidth, terminalHeight := termbox.Size()

		// Creating a new snake segment
		newBlock := d.Snake.Body[len(d.Snake.Body)-1]
		switch d.Snake.Orientation {
		case "left":
			newBlock.XCord--
		case "right":
			newBlock.XCord++
		case "up":
			newBlock.YCord--
		case "down":
			newBlock.YCord++
		}

		// Obstacle check
		if (newBlock.YCord <= 0 || newBlock.YCord >= terminalHeight-1) ||
			(newBlock.XCord <= 0 || newBlock.XCord >= terminalWidth-1) {
			d.EndStats(false, true)
		}

		for _, block := range d.Snake.Body {
			if block == newBlock {
				d.EndStats(false, true)
			}
		}

		// Checking fruit eating
		if d.Fruit.Position == newBlock {
			d.Snake.Body = append(d.Snake.Body, newBlock)
			d.CreateFruit()

			d.PlayerScore++
			if d.PlayerScore == (terminalHeight-2)*(terminalWidth-2) {
				d.EndStats(true, true)
			}
		} else {
			d.Snake.Body = append(d.Snake.Body[1:], newBlock)
		}

		time.Sleep(100*time.Millisecond - time.Since(start))
	}
}
