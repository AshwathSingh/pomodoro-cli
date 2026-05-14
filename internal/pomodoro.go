package internal

import (
	"github.com/AshwathSingh/pomodoro-cli/model"
)

// have to implement functionality to be able to interrupt pomodoro's to take a break
// additionally have the sessions running infinitely until a specifc combiniation is pressed
// has to be implemented
func PomodoroSession(pomodoro *model.Pomodoro) {
	// clear the screen and start the focus session

	if pomodoro == nil {
		panic("nil Time passed to PomodoroSession")
	}


	notifyEnd(SessionFocus, pomodoro.Split.Focus)

	for {
		// checks if the user wants to continue with a focus session
		if !decisionNextSession(SessionFocus, pomodoro.Split.Focus) {
			break
		}

		// checks if the user wants to continue with a break session
		if !decisionNextSession(SessionBreak, pomodoro.Split.Break) {
			// if it returns as false, check if the user wants to continue with a focus PomodoroSession
			if !decisionNextSession(SessionFocus, pomodoro.Split.Focus) {
				break
			}
		}

	}
}
