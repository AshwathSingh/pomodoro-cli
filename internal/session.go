package internal

import (
	"time"

	"github.com/AshwathSingh/pomodoro-cli/ui"
)

// [TODO]
// - multiple focus sessions
// - ensuring breaks and focus sessions can be interrupted and set to former or latter
// - addding task functionality (novelty feature not really core feature)
// - ensuring that time elapsed and the progress bar exist on different lines

// StartSession starts a focus session with the given label and duration in minutes
// takes in a paramter of a label (to specify which session is being completed) and durationMinutes (as an int64)
func StartSession(label string, durationMinutes uint64) {

	total := time.Duration(durationMinutes) * time.Minute
	start := time.Now()

	width := 30
	startStr := start.Format("15:04")
	countdown := total
	for {
		elapsed := time.Since(start)

		if elapsed >= total {
			break
		}
		countdown = total - elapsed
		// clear the line and render the progress bar
		ui.PrintProgress(
			startStr,
			countdown,
			elapsed,
			total,
			width,
		)

		time.Sleep(100 * time.Millisecond)
	}

	// final state
	ui.PrintProgress(
		startStr,
		total,
		total,
		total,
		width,
	)

}
