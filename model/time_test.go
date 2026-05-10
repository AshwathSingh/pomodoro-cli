package model

import "testing"

func TestTimeSettersAndInitialise(t *testing.T) {
	var tm Time
	// Initialise should set zero values
	tm.Initialise()
	if tm.Focus != 0 || tm.Break != 0 {
		t.Fatalf("Initialise did not set zeros: got Focus=%d Break=%d", tm.Focus, tm.Break)
	}

	// Setters should update values
	tm.SetFocus(25)
	tm.SetBreak(5)
	if tm.Focus != 25 || tm.Break != 5 {
		t.Fatalf("Setters failed: got Focus=%d Break=%d", tm.Focus, tm.Break)
	}
}
