package ui

import (
	"fmt"
	"strings"
)

// RenderBar renders a progress bar with the given progress and width
func RenderBar(progress float64, width int) string {
	filled := int(progress * float64(width))

	bar := strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled)

	return "[" + bar + "]"
}

// function for case 3 of input, wherein the user wants to use a custom focus session
// takes in timeFocus and timeBreak as pointers to update them in the main function caller
// to be implemented: ensuring that input passed in is correct, and not negative
func CustomInput(timeFocus, timeBreak *int64) {
	DeleteLines()
	fmt.Print("\033[F\033[K")

	fmt.Println("Custom Pomodoro Set-Up")
	fmt.Println("how long do you want to focus for?")
	fmt.Scan(timeFocus)

	DeleteLines()

	// checking the input to ensure it is not negative
	if *timeFocus < 0 {
		fmt.Println("please re-enter the focus time")
		fmt.Scan(timeFocus)

		DeleteLines()
	}

	fmt.Println("how long do you want to rest for?")
	fmt.Scan(timeBreak)

	DeleteLines()

	// checking the input to ensure it is not negative
	if *timeBreak < 0 {
		fmt.Println("please re-enter the break time")
		fmt.Scan(timeBreak)
		DeleteLines()
	}
}

// DeleteLines deletes two lines from the terminal
func DeleteLines() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}

// ClearScreen clears the screen by printing newline characters and clearing the line
func ClearScreen() {
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")
}
