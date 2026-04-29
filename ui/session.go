package ui

import (
	"fmt"
	"time"
)

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

		progress := float64(elapsed) / float64(total)

		bar := RenderBar(progress, width)

		// elapsed time formatting
		elapsedMin := int(elapsed.Minutes())
		elapsedSec := int(elapsed.Seconds()) % 60

		fmt.Printf(
			"\r%s: %d | START: %s | ELAPSED: %02d:%02d | %s",
			label,
			durationMinutes,
			startStr,
			elapsedMin,
			elapsedSec,
			bar,
		)

		time.Sleep(100 * time.Millisecond)
	}

	// final state
	elapsed := total
	fmt.Printf(
		"\r%s | START: %s | ELAPSED: %02d:%02d | %s\n",
		label,
		startStr,
		int(elapsed.Minutes()),
		int(elapsed.Seconds())%60,
		RenderBar(1.0, width),
	)
}
