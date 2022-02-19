package golorama

import "fmt"

const (
	CSI = `\033[`
	OSC = `\033]`
	BEL = `\a`
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

func PrintColorln(color int, i ...interface{}) {
	fmt.Println(CSI, color, i...)
}