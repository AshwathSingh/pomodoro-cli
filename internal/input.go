package internal

import (
	"fmt"
	"os"
)

func PomodoroSession(timeFocus, timeBreak int64) {
	// clear the screen and start the focus session
	// DeleteLines()
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
