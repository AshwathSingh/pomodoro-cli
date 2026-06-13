package internal

import (
	"context"
	"testing"
	"time"
)

// TestStartSessionZeroDuration ensures StartSession returns promptly when duration is zero.
func TestStartSessionZeroDuration(t *testing.T) {
	ctx := context.Background()
	done := make(chan struct{})
	go func() {
		StartSession(ctx, "TEST", 0)
		close(done)
	}()

	select {
	case <-done:
		// success
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("StartSession with 0 duration did not return promptly")
	}
}
