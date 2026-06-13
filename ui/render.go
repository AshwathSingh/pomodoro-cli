package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/AshwathSingh/pomodoro-cli/model"
)

// function for case 3 of input, wherein the user wants to use a custom focus session
// takes in a pointer to internal.Time to update values in the caller
// to be implemented: further validation on inputs
func CustomInput(pomodoro *model.Pomodoro) {
	DeleteLines()
	fmt.Print("\033[F\033[K")

	fmt.Println("Custom Pomodoro Set-Up")
	fmt.Println("how long do you want to focus for?")
	fmt.Scan(&pomodoro.Split.Focus)

	fmt.Println("how long do you want to rest for?")
	fmt.Scan(&pomodoro.Split.Break)

	DeleteLines()
	ClearScreen()
}

// printProgress renders the progress bar for the current session
func PrintProgress(label string, startStr string, countdown time.Duration, elapsed time.Duration, total time.Duration, width int) {
	progress := float64(elapsed) / float64(total)
	bar := renderBar(progress, width)

	fmt.Printf(
		"\r%s | START: %s | ELAPSED: %02d:%02d | %s",
		label,
		startStr,
		int(countdown.Minutes()),
		int(countdown.Seconds())%60,
		bar,
	)
}

// RenderBar renders a progress bar with the given progress and width
func renderBar(progress float64, width int) string {
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

func ClearScreen() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}
