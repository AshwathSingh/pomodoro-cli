package internal

import (
	"context"

	"github.com/AshwathSingh/pomodoro-cli/model"
)

// PomodoroSession runs the pomodoro session loop with the given context.
// The context can be cancelled to gracefully exit the entire pomodoro session.
// It cycles between focus and break sessions until the user decides to quit.
func PomodoroSession(ctx context.Context, pomodoro *model.Pomodoro) {
	if pomodoro == nil {
		panic("nil Pomodoro passed to PomodoroSession")
	}

	notifyEnd(ctx, SessionFocus, pomodoro.Split.Focus)

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return
		default:
		}

		// checks if the user wants to continue with a focus session
		if !decisionNextSession(ctx, SessionFocus, pomodoro.Split.Focus) {
			break
		}

		// checks if the user wants to continue with a break session
		if !decisionNextSession(ctx, SessionBreak, pomodoro.Split.Break) {
			// if it returns as false, check if the user wants to continue with a focus PomodoroSession
			if !decisionNextSession(ctx, SessionFocus, pomodoro.Split.Focus) {
				break
			}
		}
	}
}
