package game

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
)

// Displaying the playing field
func (d *Data) GameScreen() {
	// Cleaning the terminal
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	// Receiving terminal width and height
	terminalWidth, terminalHeight := termbox.Size()

	// Edge displaying
	for xCord := 0; xCord < terminalWidth; xCord++ {
		termbox.SetCell(xCord, 0, []rune("-")[0],
			termbox.AttrBold, termbox.ColorDefault)
		termbox.SetCell(xCord, terminalHeight-1, []rune("-")[0],
			termbox.AttrBold, termbox.ColorDefault)
	}

	for yCord := 0; yCord < terminalHeight; yCord++ {
		termbox.SetCell(0, yCord, []rune("|")[0],
			termbox.AttrBold, termbox.ColorDefault)
		termbox.SetCell(terminalWidth-1, yCord, []rune("|")[0],
			termbox.AttrBold, termbox.ColorDefault)
	}

	for _, xCord := range []int{0, terminalWidth - 1} {
		for _, yCord := range []int{0, terminalHeight - 1} {
			termbox.SetCell(xCord, yCord, []rune("+")[0],
				termbox.AttrBold, termbox.ColorDefault)
		}
	}

	// Displaying the snake on the field
	for _, block := range d.Snake.Body {
		termbox.SetCell(block.XCord, block.YCord, []rune("#")[0],
			termbox.ColorLightGreen, termbox.ColorDefault)
	}

	// Displaying the fruit on the field
	termbox.SetCell(d.Fruit.Position.XCord, d.Fruit.Position.YCord, d.Fruit.Shape,
		d.Fruit.Color, termbox.ColorDefault)

	termbox.Flush()
}

// Displaying the player's score after the end of the game
func (d *Data) EndStats(isWin, isDelay bool) {
	if isDelay {
		time.Sleep(time.Second)
	}

	termbox.Close()

	if isWin {
		fmt.Println(models.WinMessage)
	} else {
		fmt.Println(models.LoseMessage)
	}

	fmt.Printf(models.StatisticalMessage, d.PlayerScore)
	os.Exit(0)
}
