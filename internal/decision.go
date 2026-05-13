package internal

import (
	"fmt"

	"github.com/AshwathSingh/pomodoro-cli/ui"
	"github.com/gen2brain/beeep"
)

func decisionNextSession(label string, duration uint64) bool {
	session := nextSession(label)
	ui.DeleteLines()
	if session {
		notifyEnd(label, duration)
		return true
	} else {
		return false
	}

}

func nextSession(label string) bool {
	fmt.Print("\n")
	switch label {
	case SessionFocus:
		fmt.Printf("do you want to start another focus session? (y / n)\n")
	case SessionBreak:
		fmt.Printf("do you want to take a break (y / n)?\n")
	default:
		panic("hey, this isn't a label I know how to deal with")
	}

	var decision string
	fmt.Scanln(&decision)

	if decision != "y" && decision != "n" {
		fmt.Printf("invalid entry, assuming ended")
		return false
	}

	if decision == "y" {
		return true
	}

	return false
}

func notifyEnd(label string, duration uint64) {
	session := StartSession(label, duration)

	if session {
		beeep.Notify(label, "your session has been completed", "")
	}
}
