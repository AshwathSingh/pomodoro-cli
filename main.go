package main

import (
	"flag"
	"fmt"
	"strings"

	"time"
)

func main() {
	timePtr := flag.Int("pomo", 25, "pomodoro work section")
	breakPtr := flag.Int("break", 5, "break section")

	flag.Parse()



	startSession("FOCUS", int64(*timePtr))
	startSession("BREAK", int64(*breakPtr))

}

func renderBar(progress float64, width int) string {
	filled := int(progress * float64(width))

	bar := strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled)

	return "[" + bar + "]"
}

func startSession(label string, durationMinutes int64) {
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

		bar := renderBar(progress, width)

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
		renderBar(1.0, width),
	)
}
