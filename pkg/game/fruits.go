package game

import (
	"math/rand"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
)

func (d *Data) CreateFruit() {
	// Receiving terminal width and height
	terminalWidth, terminalHeight := termbox.Size()

	// Selecting a random fruit
	fruitType := rand.Intn(len(models.FruitTypes))
	d.Fruit = models.FruitTypes[fruitType]

	// Checking fetal coordinates relative to the snake
cords:
	for {
		d.Fruit.Position = models.Position{
			XCord: rand.Intn(terminalWidth-2) + 1,
			YCord: rand.Intn(terminalHeight-2) + 1,
		}

		for _, block := range d.Snake.Body {
			if block == d.Fruit.Position {
				continue cords
			}
		}

		break
	}
}
