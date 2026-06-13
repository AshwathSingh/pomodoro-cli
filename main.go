package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AshwathSingh/pomodoro-cli/internal"
	"github.com/AshwathSingh/pomodoro-cli/model"
	"github.com/AshwathSingh/pomodoro-cli/ui"
)

func main() {

	// Set up signal handling for graceful shutdown (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a context that can be cancelled on signal
	ctx, cancel := context.WithCancel(context.Background())

	// Handle signals in a goroutine
	go func() {
		<-sigChan
		fmt.Println("\nShutting down gracefully...")
		cancel()
	}()

	// refactoring of time into a pomodoro object
	pomodoroObject := model.Pomodoro{}

	// initialsing the pomodoroObject with initialiser
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

	internal.PomodoroSession(ctx, &pomodoroObject)
}
