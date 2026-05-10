package main

import (
	"fmt"
	"log"

	"github.com/AshwathSingh/pomodoro-cli/internal"
	"github.com/AshwathSingh/pomodoro-cli/model"
	"github.com/AshwathSingh/pomodoro-cli/ui"
)

// initialize timeFocus and timeBreak variables
// TODO - implement this as a struct with functions to access the
// variables as good practice

func main() {

	pomodoro := model.Time{Focus: 0, Break: 0}
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

	// switch statement to set focus and break based on user input
	switch inputOption {
	case 1:
		pomodoro.SetFocus(25)
		pomodoro.SetBreak(5)
		ui.ClearScreen()
	case 2:
		pomodoro.SetFocus(50)
		pomodoro.SetBreak(10)
		ui.ClearScreen()
	case 3:
		ui.CustomInput(&pomodoro)
	default:
		log.Fatal("invalid option")
	}

	internal.PomodoroSession(&pomodoro)
}
