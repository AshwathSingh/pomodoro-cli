package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AshwathSingh/pomodoro-cli/ui"
)

func main() {

	fmt.Printf("(1) 25 / 5 \n(2) 50 / 10 \n(3) custom\n")
	var inputOption int

	_, err := fmt.Scan(&inputOption)

	fmt.Print("\033[F\033[K")
	if err != nil {
		log.Fatal("error inputting information")
		return
	}

	var timeFocus int64
	var timeBreak int64

	timeFocus = 0
	timeBreak = 0

	switch inputOption {
	case 1:
		timeFocus = 25
		timeBreak = 5
	case 2:
		timeFocus = 50
		timeBreak = 10
	case 3:
		customInput(&timeFocus, &timeBreak)
	default:
		log.Fatal("invalid option")
	}
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

// function for case 3 of input, wherein the user wants to use a custom focus session
// takes in timeFocus and timeBreak as pointers to update them in the main function caller
// to be implemented: ensuring that input passed in is correct, and not negative
func customInput(timeFocus, timeBreak *int64) {
	ui.DeleteLines()
	fmt.Print("\033[F\033[K")

	fmt.Println("Custom Pomodoro Set-Up")
	fmt.Println("how long do you want to focus for?")
	fmt.Scan(timeFocus)

	ui.DeleteLines()

	// checking the input to ensure it is not negative
	if *timeFocus < 0 {
		fmt.Println("please re-enter the focus time")
		fmt.Scan(timeFocus)

		ui.DeleteLines()
	}

	fmt.Println("how long do you want to rest for?")
	fmt.Scan(timeBreak)

	ui.DeleteLines()

	// checking the input to ensure it is not negative
	if *timeBreak < 0 {
		fmt.Println("please re-enter the break time")
		fmt.Scan(timeBreak)
		ui.DeleteLines()
	}
}
