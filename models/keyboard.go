package models

import "github.com/nsf/termbox-go"

var ControlKeys = map[termbox.Key]string{
	termbox.KeyArrowUp:    "up",
	termbox.KeyArrowDown:  "down",
	termbox.KeyArrowLeft:  "left",
	termbox.KeyArrowRight: "right",
	termbox.KeyEsc:        "quit",
}

var ControlCharacters = map[string]string{
	"w": "up",
	"s": "down",
	"a": "left",
	"d": "right",
	"q": "quit",
}

var MovementLock = map[string]string{
	"down":  "up",
	"up":    "down",
	"right": "left",
	"left":  "right",
}

var MovementTypes = []string{"up", "down", "left", "right"}
