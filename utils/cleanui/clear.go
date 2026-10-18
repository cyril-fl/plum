// Package cleanui provides terminal display helpers.
package cleanui

import "fmt"

const (
	// CLEAR is the terminal escape sequence that clears the screen.
	CLEAR = "\033[H\033[2J\033[3J"
)

// Run clears the terminal screen.
func Run() {
	fmt.Print(CLEAR)
}
