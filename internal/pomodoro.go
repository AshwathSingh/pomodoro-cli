package internal

import (
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


	notifyEnd(SessionFocus, t.Focus)

	for {
		// checks if the user wants to continue with a focus session
		if !decisionNextSession(SessionFocus, t.Focus) {
			break
		}

		// checks if the user wants to continue with a break session
		if !decisionNextSession(SessionBreak, t.Break) {
			// if it returns as false, check if the user wants to continue with a focus PomodoroSession
			if !decisionNextSession(SessionFocus, t.Focus) {
				break
			}
		}

	}
}
