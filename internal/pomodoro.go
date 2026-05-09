package internal

import (
	"fmt"
	"os"

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


	os.Stdout.Sync()
	StartSession("FOCUS", t.Focus)
	fmt.Printf("\033[F\033[K")

	fmt.Print("do you want to take a break? (y/n) ")
	var takeBreak string
	fmt.Scan(&takeBreak)

	if takeBreak != "y" {
		return
	}

	StartSession("BREAK", t.Break)
}
