package main

import (
	"fmt"
	"log"

	"github.com/AshwathSingh/pomodoro-cli/ui"
)

func main() {

	// menu to select focus session type
	fmt.Printf("(1) 25 / 5 \n(2) 50 / 10 \n(3) custom\n")
	var inputOption int

	// scan the user's input for the focus session type
	_, err := fmt.Scan(&inputOption)

	fmt.Print("\033[F\033[K")
	if err != nil {
		log.Fatal("error inputting information")
		return
	}

	// initialize timeFocus and timeBreak variables
	var timeFocus int64
	var timeBreak int64

	timeFocus = 0
	timeBreak = 0

	// switch statement to set timeFocus and timeBreak based on user input
	switch inputOption {
	case 1:
		timeFocus = 25
		timeBreak = 5
		fmt.Print("\033[F\033[K")
		ui.DeleteLines()
	case 2:
		timeFocus = 50
		timeBreak = 10
		fmt.Print("\033[F\033[K")
		ui.DeleteLines()
	case 3:
		ui.CustomInput(&timeFocus, &timeBreak)
	default:
		log.Fatal("invalid option")
	}

	ui.PomodoroSession(timeFocus, timeBreak)
}
