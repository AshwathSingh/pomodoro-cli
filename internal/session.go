package internal

import (
	"fmt"
	"time"

	"github.com/AshwathSingh/pomodoro-cli/ui"
)

// to be implemented
// - multiple focus sessions
// - ensuring breaks and focus sessions can be interrupted and set to former or latter
// - addding task functionality (novelty feature not really core feature)
// - ensuring that time elapsed and the progress bar exist on different lines

// StartSession starts a focus session with the given label and duration in minutes
func StartSession(label string, durationMinutes int64) {
	total := time.Duration(durationMinutes) * time.Minute
	start := time.Now()

	width := 30
	startStr := start.Format("15:04")

	for {
		elapsed := time.Since(start)

		if elapsed >= total {
			break
		}
		// clear the line and render the progress bar
		printProgress(
			startStr,
			elapsed,
			total,
			width,
		)

		time.Sleep(100 * time.Millisecond)
	}

	// final state
	printProgress(
		startStr,
		total,
		total,
		width,
	)

}

// printProgress renders the progress bar for the current session
func printProgress(startStr string, elapsed time.Duration, total time.Duration, width int) {
	progress := float64(elapsed) / float64(total)
	bar := ui.RenderBar(progress, width)

	fmt.Printf(
		"\r START: %s | ELAPSED: %02d:%02d | %s",
		startStr,
		int(elapsed.Minutes()),
		int(elapsed.Seconds())%60,
		bar,
	)
}
