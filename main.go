package main

import (
	"fmt"
	"log"
	"os"

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
	case 2:
		timeFocus = 50
		timeBreak = 10
	case 3:
		ui.CustomInput(&timeFocus, &timeBreak)
	default:
		log.Fatal("invalid option")
	}

	// clear the screen and start the focus session
	ui.ClearScreen()
	os.Stdout.Sync()
	ui.StartSession("FOCUS", timeFocus)
	fmt.Printf("\033[F\033[K")

	fmt.Print("do you want to take a break? (y/n) ")
	var takeBreak string
	fmt.Scan(&takeBreak)

	if takeBreak != "y" {
		return
	}
	ui.StartSession("BREAK", timeBreak)

}
