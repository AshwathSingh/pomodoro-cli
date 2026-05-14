package main

import (
	"fmt"
	"log"

	"github.com/AshwathSingh/pomodoro-cli/internal"
	"github.com/AshwathSingh/pomodoro-cli/model"
	"github.com/AshwathSingh/pomodoro-cli/ui"
)


func main() {

	// refactoring of time into a pomodoro object
	pomodoroObject := model.Pomodoro{}
	pomodoroObject.InitialisePomodoro()

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
		pomodoroObject.SetFocus(25)
		pomodoroObject.SetBreak(5)
		ui.ClearScreen()
	case 2:
		pomodoroObject.SetFocus(50)
		pomodoroObject.SetBreak(10)
		ui.ClearScreen()
	case 3:
		ui.CustomInput(&pomodoroObject)
	default:
		log.Fatal("invalid option")
	}

	internal.PomodoroSession(&pomodoroObject)
}
