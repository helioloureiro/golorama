package golorama

import "fmt"

const (
	CSI = "\033["
	OSC = "\033["
	BEL = '\a'
)

const (
	BRIGHT    = 1
	BOLD      = 1
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

const (
	LIGHTBLACK = iota + 90
	LIGHTRED
	LIGHTGREEN
	LIGHTYELLOW
	LIGHTBLUE
	LIGHTMAGENTA
	LIGHTCYAN
	LIGHTWHITE
)

// Color assignment
type Color int

// GetCSI: to get the color as termio format
func GetCSI(color Color) string {
	return fmt.Sprintf("%s%dm", CSI, color)
}

// Reset: set color off to the terminal
func Reset() string {
	return GetCSI(RESET_ALL)
}

// PrintColorln: a direct call to print lines in colorful way
func PrintColorln(color Color, i ...interface{}) {
	fmt.Println(GetCSI(color), i, Reset())
}
