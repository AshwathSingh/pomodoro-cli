package ui

import (
	"fmt"
	"strings"
)

// RenderBar renders a progress bar with the given progress and width
func RenderBar(progress float64, width int) string {
	filled := int(progress * float64(width))

	bar := strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled)

	return "[" + bar + "]"
}

// DeleteLines deletes two lines from the terminal
func DeleteLines() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}

// ClearScreen clears the screen by printing newline characters and clearing the line
func ClearScreen() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}
