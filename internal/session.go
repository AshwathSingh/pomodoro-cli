package internal

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/AshwathSingh/pomodoro-cli/ui"
)

// [TODO]
// - addding task functionality (novelty feature not really core feature)
// - ensuring that time elapsed and the progress bar exist on different lines

// StartSession starts a focus session with the given label and duration in minutes.
// It listens for user input and will interrupt the session early if the user
// types 'q' (case-insensitive) and presses Enter. The provided context can be used
// to cancel the session externally.
func StartSession(ctx context.Context, label string, durationMinutes uint64) bool {

	total := time.Duration(durationMinutes) * time.Minute
	start := time.Now()

	width := 30
	startStr := start.Format("15:04")

	// Channel used to signal an interrupt from the input goroutine
	stopCh := make(chan struct{})
	var wg sync.WaitGroup

	// Start goroutine to listen for 'q' input to interrupt the session.
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(os.Stdin)
		for {
			// Use a non-blocking check for context cancellation
			select {
			case <-ctx.Done():
				return
			default:
			}

			// Set read deadline based on context or timeout
			input, err := reader.ReadString('\n')
			if err != nil {
				// On error (including EOF), stop listening
				return
			}
			trim := strings.TrimSpace(input)
			if strings.EqualFold(trim, "q") {
				select {
				case <-stopCh:
					// Already closed, avoid double close panic
				default:
					close(stopCh)
				}
				return
			}
		}
	}()

	// Create a channel for the timer tick
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	interrupted := false

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

		// Wait either for context cancellation, session interrupt, or a tick
		select {
		case <-ctx.Done():
			// Context cancelled externally (e.g., Ctrl+C)
			fmt.Print("\033[F\033[K")
			ui.PrintProgress(
				label,
				startStr,
				(total - elapsed),
				elapsed,
				total,
				width,
			)
			interrupted = true
			break
		case <-stopCh:
			// Interrupted by user typing 'q'
			fmt.Print("\033[F\033[K")
			ui.PrintProgress(
				label,
				startStr,
				(total - elapsed),
				elapsed,
				total,
				width,
			)
			interrupted = true
			break
		case <-ticker.C:
			// Continue loop
		}

		if interrupted {
			break
		}
	}

	// Wait for input goroutine to finish
	wg.Wait()

	if interrupted {
		return false
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
