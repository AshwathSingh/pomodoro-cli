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

	os.Stdout.Sync()
	ui.StartSession("FOCUS", timeFocus)
	ui.StartSession("BREAK", timeBreak)

}

func customInput(timeFocus, timeBreak *int64) {
	fmt.Println("how long do you want to focus for?")
	fmt.Scan(timeFocus)
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")

	fmt.Println("how long do you want to rest for?")
	fmt.Scan(timeBreak)
	fmt.Print("\033[F\033[K")
	fmt.Print("\033[F\033[K")

}
