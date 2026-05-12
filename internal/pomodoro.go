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

	// initial is a variable to ensure first focus session runs uninterrupted
	var initial int8 = 0

	for {
		if initial == 1 {
			// checks if the user wants to continue with a focus session
			decisionContinue := decisionNextSession("FOCUS", t.Focus)
			if !decisionContinue {
				break
			}
		} else {
			notifyEnd("FOCUS", t.Focus)
			initial = initial + 1
		}

		// checks if user want to continue with a break session
		decisionContinue := decisionNextSession("BREAK", t.Break)

		if !decisionContinue {
			if !decisionNextSession("FOCUS", t.Focus) {
				break
			}
		}

	}
}
