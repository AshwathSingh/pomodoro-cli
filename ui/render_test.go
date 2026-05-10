package ui

import (
	"strings"
	"testing"
)

func TestRenderBarFullHalfEmpty(t *testing.T) {
	width := 10
	bar := renderBar(1.0, width)
	// Expect width "filled" blocks inside brackets
	if strings.Count(bar, "█") != width {
		t.Fatalf("expected %d filled blocks, got %d for bar %q", width, strings.Count(bar, "█"), bar)
	}

	bar = renderBar(0.0, width)
	if strings.Count(bar, "█") != 0 {
		t.Fatalf("expected 0 filled blocks for 0%% progress, got %d", strings.Count(bar, "█"))
	}

	bar = renderBar(0.5, width)
	if strings.Count(bar, "█") != width/2 {
		t.Fatalf("expected %d filled blocks for 50%% progress, got %d", width/2, strings.Count(bar, "█"))
	}
}
