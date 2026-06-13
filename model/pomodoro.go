package model

// a pomodoro object to hold the Splits of focus and break
// and statistics for a session
type Pomodoro struct {
	Split     Time
	Task      string
	FocusTime uint64
	BreakTime uint64
}

func (p *Pomodoro) InitialisePomodoro() {
	p.Split.Initialise()
	p.AddBreakTime(0)
	p.AddBreakTime(0)
	p.SetTask("")
}

func (p *Pomodoro) SetFocus(duration uint64) {
	p.Split.SetFocus(duration)
}

func (p *Pomodoro) SetBreak(duration uint64) {
	p.Split.SetBreak(duration)
}


func (p *Pomodoro) SetTask(task string) {
	p.Task = task
}

func (p *Pomodoro) AddFocusTime(duration uint64) {
	p.FocusTime += duration
}

func (p *Pomodoro) AddBreakTime(duration uint64) {
	p.BreakTime += duration
}




