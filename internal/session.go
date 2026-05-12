package internal

import (
	"time"

	"github.com/AshwathSingh/pomodoro-cli/ui"
)

// [TODO]
// - ensuring breaks and focus sessions can be interrupted and set to former or latter
// - addding task functionality (novelty feature not really core feature)
// - ensuring that time elapsed and the progress bar exist on different lines

// StartSession starts a focus session with the given label and duration in minutes
func StartSession(label string, durationMinutes uint64) (bool) {

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
		ui.PrintProgress(
			label,
			startStr,
			(total - elapsed),
			elapsed,
			total,
			width,
		)

		time.Sleep(100 * time.Millisecond)
	}

	// final state
	ui.PrintProgress(
		label,
		startStr,
		total,
		total,
		total,
		width,
	)

	return true

}
