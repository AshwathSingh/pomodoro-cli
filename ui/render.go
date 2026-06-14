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

// RenderBar renders a smooth progress bar with fractional blocks
// Uses Unicode block characters to show fine-grained progress
func renderBar(progress float64, width int) string {
	// Fractional block characters: higher indices = more filled
	blocks := []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}

	filledWidth := progress * float64(width)
	fullBlocks := int(filledWidth)
	fractionPart := filledWidth - float64(fullBlocks)

	var bar strings.Builder
	bar.WriteString("[")

	// Write full blocks
	for i := 0; i < fullBlocks; i++ {
		bar.WriteString("█")
	}

	// Write fractional block if we have one
	if fullBlocks < width {
		fractionIndex := int(fractionPart * float64(len(blocks)-1))
		if fractionIndex >= len(blocks) {
			fractionIndex = len(blocks) - 1
		}
		bar.WriteString(blocks[fractionIndex])
		fullBlocks++
	}

	// Write empty blocks
	for i := fullBlocks; i < width; i++ {
		bar.WriteString(" ")
	}

	bar.WriteString("]")
	return bar.String()
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
