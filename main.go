package main

import (
	"fmt"
	"log"
	"strings"

	"time"
)

func main() {

	fmt.Printf("(1) 25 / 5 \n(2) 50 / 10 \n(3) custom\n")
	var inputOption int

	_, err := fmt.Scan(&inputOption)

	fmt.Print("\033[F\033[K")
	if err != nil {
		log.Fatal("error inputting information")
		return
	}

	var timeFocus int64
	var timeBreak int64

	timeFocus = 0
	timeBreak = 0

	switch inputOption {
	case 1:
		timeFocus = 25
		timeBreak = 5
	case 2:
		timeFocus = 50
		timeBreak = 10
	case 3:
		panic("not implemented")
	default:
		log.Fatal("invalid option")
	}
	startSession("FOCUS", timeFocus)
	startSession("BREAK", timeBreak)

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
