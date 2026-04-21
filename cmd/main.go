package main

import (
	"fmt"

	"github.com/AshwathSingh/pomodoro-cli/internal/app"
	"github.com/AshwathSingh/pomodoro-cli/internal/config"
)

func main() {
	fmt.Println(config.KeyDown)
	fmt.Println(" 50/10")
	fmt.Println("please enter something here to test")
	terminalInput := app.GetInput()
	fmt.Println(terminalInput)

}
