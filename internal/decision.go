package internal

import (
	"context"
	"fmt"

	"github.com/AshwathSingh/pomodoro-cli/ui"
	"github.com/gen2brain/beeep"
)

func decisionNextSession(ctx context.Context, label string, duration uint64) bool {
	session, err := nextSession(ctx, label)

	if err != nil {
		fmt.Println(err)
	}
	ui.DeleteLines()
	if session {
		notifyEnd(ctx, label, duration)
		return true
	} else {
		return false
	}

}

func nextSession(ctx context.Context, label string) (bool, error) {
	fmt.Print("\n")

	if label == SessionFocus {
		fmt.Printf("do you want to start another focus session? (y / n)\n")

	} else if label == SessionBreak {
		fmt.Printf("do you want to take a break (y / n)?\n")

	} else {
		panic("hey, this isn't a label I know how to deal with")

	}
	var decision string

	// error checking for Scanln
	_, err := fmt.Scanln(&decision)

	if err != nil {
		return false, err
	}

	if decision != "y" && decision != "n" {
		fmt.Printf("invalid entry, assuming ended")
		return false, nil
	}

	if decision == "y" {
		return true, nil
	}

	return false, nil
}

func notifyEnd(ctx context.Context, label string, duration uint64) {

	session := StartSession(ctx, label, duration)

	if session {
		beeep.Notify(label, "your session has been completed", "")
	}
}
