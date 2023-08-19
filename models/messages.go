package models

// Version data
const Version string = "1.0.0"

// Informational messages
var InfoMessage string = "" +
	"Name:     \033[1;35mSnake\033[0m\n" +
	"Version:  \033[1;36m%s\033[0m\n" +
	"Codename: \033[1;31mChika\033[0m\n" +
	"Creator:  \033[1;34mNKTKLN\033[0m\n"

var KeysMessage string = "" +
	"\033[1;34mAll control keys in the game:\033[0m\n\n" +
	"\033[1;37m\"a\"\033[0m or \033[1;37mArrow Left\033[0m  - Moves the snake to the left.\n" +
	"\033[1;37m\"d\"\033[0m or \033[1;37mArrow Right\033[0m - Moves a snake to the right.\n" +
	"\033[1;37m\"s\"\033[0m or \033[1;37mArrow Down\033[0m  - Moves the snake down.\n" +
	"\033[1;37m\"u\"\033[0m or \033[1;37mArrow Up\033[0m    - Moves the snake up.\n" +
	"\033[1;37m\"q\"\033[0m or \033[1;37mESC\033[0m         - Quits the game.\n"

var SizeErrorMessage string = "\033[1;31mTerminal window size is smaller than 9x9!\033[0m"

// End messages
var StatisticalMessage string = "\033[1;37mScore: \033[1;35m%d\n"
var WinMessage string = "\033[1;33mYou Win!\033[0m"
var LoseMessage string = "\033[1;31mYou Lose!\033[0m"
