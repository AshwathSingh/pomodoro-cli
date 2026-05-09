package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AshwathSingh/pomodoro-cli/model"
)

// have to implement functionality to be able to interrupt pomodoro's to take a break
// additionally have the sessions running infinitely until a specifc combiniation is pressed
// has to be implemented
func PomodoroSession(t *model.Time) {
	// clear the screen and start the focus session

	if t == nil {
		panic("nil Time passed to PomodoroSession")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// catch interrupt/terminate
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
	}()

	// alternate sessions until ctx is cancelled
	for ctx.Err() == nil {
		StartSession("FOCUS", t.Focus, ctx)
		if ctx.Err() != nil {
			break
		}
		StartSession("BREAK", t.Break, ctx)
	}

	// optional: cleanup/readable exit message
	fmt.Println("\nSession cancelled — exiting.")
}
