package ui

import (
	"fmt"
	"strings"
)

func RenderBar(progress float64, width int) string {
	filled := int(progress * float64(width))

	bar := strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled)

	return "[" + bar + "]"
}

func DeleteLines() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}

func ClearScreen() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}
