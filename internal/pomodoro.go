package internal

import (
	"fmt"
	"os"

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
		os.Stdout.Sync()

		if initial == 1 {
			fmt.Println("\n do you want to start another focus session (y / n)")
			var takeBreak string
			fmt.Scan(&takeBreak)

			ui.DeleteLines()
			if takeBreak != "y" {
				stop = 1
				break
			} else {
				StartSession("FOCUS", t.Break)
				initial = 1
			}

		} else {
			StartSession("FOCUS", t.Focus)
			initial = initial + 1
		}
		fmt.Println("\n do you want to take a break? (y/n) ")
		var takeBreak string
		fmt.Scan(&takeBreak)

		ui.DeleteLines()
		if takeBreak != "y" {
			stop = 1
			break
		} else {
			StartSession("BREAK", t.Break)
			initial = 1
		}
	}
}
