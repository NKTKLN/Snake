package game

import "time"

func (d *Data) Game() {
	// Launching the main fnctions
	d.CreateSnake()
	d.CreateFruit()

	go d.ReadingKeyboard()
	go d.MoveSnake()

	// Loop responsible for displaying the field in 25 fps
	for {
		start := time.Now()

		d.GameScreen()

		time.Sleep(40*time.Millisecond - time.Since(start))
	}
}
