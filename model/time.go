package model

// Time holds pomodoro durations in minutes
type Time struct {
	Focus uint64
	Break uint64
}

func (t *Time) SetFocus(timer uint64) {
	t.Focus = timer
}

func (t *Time) SetBreak(timer uint64) {
	t.Break = timer
}

func (t *Time) Initialise() {
	t.Focus = 0
	t.Break = 0
}
