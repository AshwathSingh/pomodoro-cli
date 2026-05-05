package internal

import (
	"fmt"
	"os"
)

// have to implement functionality to be able to interrupt pomodoro's to take a break
// additionally have the sessions running infinitely until a specifc combiniation is pressed
// has to be implemented
func PomodoroSession(timeFocus, timeBreak int64) {
	// clear the screen and start the focus session

	if timeFocus < 0 || timeBreak < 0 {
		panic("error, passing of values lead to error in storage of time focus and break")
	}

	os.Stdout.Sync()
	StartSession("FOCUS", timeFocus)
	fmt.Printf("\033[F\033[K")

	fmt.Print("do you want to take a break? (y/n) ")
	var takeBreak string
	fmt.Scan(&takeBreak)

	if takeBreak != "y" {
		return
	}
	StartSession("BREAK", timeBreak)

}
