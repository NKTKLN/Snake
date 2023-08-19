package models

import "github.com/nsf/termbox-go"

type Position struct {
	XCord int
	YCord int
}

type Snake struct {
	Body        []Position
	Orientation string
}

type Fruit struct {
	Position Position
	Shape    rune
	Color    termbox.Attribute
}

var FruitTypes = []Fruit{
	{
		Shape: []rune("@")[0],
		Color: termbox.ColorLightRed,
	},
	{
		Shape: []rune("0")[0],
		Color: termbox.ColorLightYellow,
	},
	{
		Shape: []rune("*")[0],
		Color: termbox.ColorLightBlue,
	},
	{
		Shape: []rune("%")[0],
		Color: termbox.ColorLightMagenta,
	},
	{
		Shape: []rune("8")[0],
		Color: termbox.ColorLightCyan,
	},
}
