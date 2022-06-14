package golorama

import "fmt"

const (
	CSI = "\033["
	OSC = "\033["
	BEL = '\a'
)

const (
	BRIGHT    = 1
	DIM       = 2
	NORMAL    = 22
	RESET_ALL = 0
)

const (
	BLACK = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
	RESET
)

const (
	LIGHTBLACK_EX = iota + 90
	LIGHTRED_EX
	LIGHTGREEN_EX
	LIGHTYELLOW_EX
	LIGHTBLUE_EX
	LIGHTMAGENTA_EX
	LIGHTCYAN_EX
	LIGHTWHITE_EX
)

// GetCSI: to get the color as termio format
func GetCSI(color int) string {
	colorTAG := fmt.Sprintf("%s%dm", CSI, color)
	return colorTAG
}

// Reset: set color off to the terminal
func Reset() string {
	return "\033[0m"
}

// PrintColorln: a direct call to print lines in colorful way
func PrintColorln(color int, i ...interface{}) {
	fmt.Println(GetCSI(color), i, Reset())
}
