package internal

import (
	"fmt"

	"github.com/AshwathSingh/pomodoro-cli/model"
	"github.com/AshwathSingh/pomodoro-cli/ui"
)

// have to implement functionality to be able to interrupt pomodoro's to take a break
// additionally have the sessions running infinitely until a specifc combiniation is pressed
// has to be implemented
func PomodoroSession(t *model.Time) {
	// clear the screen and start the focus session

	if t == nil {
		panic("nil Time passed to PomodoroSession")
	}

	var stop int8 = 0
	var initial int8 = 0

	for stop != 1 {

		if initial == 1 {
			focus := nextSession("FOCUS")
			ui.DeleteLines()
			if focus {
				StartSession("FOCUS", t.Break)
				initial = 1

			} else {
				stop = 1
				break
			}

		} else {
			StartSession("FOCUS", t.Focus)
			initial = initial + 1
		}

		break_session := nextSession("BREAK")
		ui.DeleteLines()
		if break_session {
			StartSession("BREAK", t.Break)
			initial = 1

		} else {
			focus_session := nextSession("FOCUS")
			ui.DeleteLines()
			if focus_session {
				StartSession("FOCUS", t.Focus)
			} else {
				break

			}

		}
	}
}

func nextSession(label string) bool {
	fmt.Print("\n")
	if label == "FOCUS" {
		fmt.Printf("do you want to start another focus session? (y / n)\n")
	} else if label == "BREAK" {
		fmt.Printf("do you want to take a break (y / n)?\n")
	}

	var decision string
	fmt.Scanln(&decision)

	if decision == "y" {
		return true
	}

	return false
}
