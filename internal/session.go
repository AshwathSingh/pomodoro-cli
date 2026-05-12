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
		// pritning the elasped time left
		// may have to rework the naming conventions
		ui.PrintProgress(
			label,
			startStr,
			(total - elapsed),
			elapsed,
			total,
			width,
		)

		// Wait either for the short tick or an interrupt
		select {
		case <-stopCh:
			// delete the line so that only the interrupted session line is shown rather than both
			fmt.Print("\033[F\033[K")
			// interrupted by user, print final state at current elapsed and return false
			ui.PrintProgress(
				label,
				startStr,
				(total - elapsed),
				elapsed,
				total,
				width,
			)
			return false
		case <-time.After(100 * time.Millisecond):
			// continue loop
		}
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
