package colorbear

import (
	"testing"
	"time"
)

func TestProgressBarCreation(t *testing.T) {
	bar := NewProgress(100)

	if bar.total != 100 {
		t.Errorf("Expected total 100, got %d", bar.total)
	}
	if bar.current != 0 {
		t.Errorf("Expected current 0, got %d", bar.current)
	}
	if bar.width != 40 {
		t.Errorf("Expected width 40, got %d", bar.width)
	}
}

func TestProgressBarOptions(t *testing.T) {
	bar := NewProgress(100,
		WithWidth(60),
		WithPrefix("Test:"),
		WithPercent(false),
		WithCount(true),
		WithColor(GreenCode),
	)

	if bar.width != 60 {
		t.Errorf("Expected width 60, got %d", bar.width)
	}
	if bar.prefix != "Test:" {
		t.Errorf("Expected prefix 'Test:', got %q", bar.prefix)
	}
	if bar.showPercent != false {
		t.Error("Expected showPercent false")
	}
	if bar.showCount != true {
		t.Error("Expected showCount true")
	}
	if bar.color != GreenCode {
		t.Errorf("Expected color %q, got %q", GreenCode, bar.color)
	}
}

func TestProgressBarSet(t *testing.T) {
	ForceColors(false) // Disable colors for test
	bar := NewProgress(100)

	bar.Set(50)
	if bar.current != 50 {
		t.Errorf("Expected current 50, got %d", bar.current)
	}

	bar.Set(100)
	if bar.current != 100 {
		t.Errorf("Expected current 100, got %d", bar.current)
	}
}

func TestProgressBarIncrement(t *testing.T) {
	bar := NewProgress(100)

	bar.Increment()
	if bar.current != 1 {
		t.Errorf("Expected current 1 after Increment(), got %d", bar.current)
	}

	bar.Increment()
	if bar.current != 2 {
		t.Errorf("Expected current 2 after second Increment(), got %d", bar.current)
	}
}

func TestProgressBarAdd(t *testing.T) {
	bar := NewProgress(100)

	bar.Add(10)
	if bar.current != 10 {
		t.Errorf("Expected current 10 after Add(10), got %d", bar.current)
	}

	bar.Add(25)
	if bar.current != 35 {
		t.Errorf("Expected current 35 after Add(25), got %d", bar.current)
	}
}

func TestProgressBarCapping(t *testing.T) {
	bar := NewProgress(100)

	// Test Set beyond total
	bar.Set(150)
	if bar.current != 150 {
		t.Errorf("Set should allow values beyond total, got %d", bar.current)
	}

	// Test Increment beyond total
	bar2 := NewProgress(10)
	for i := 0; i < 15; i++ {
		bar2.Increment()
	}
	if bar2.current != 10 {
		t.Errorf("Increment should cap at total, got %d", bar2.current)
	}

	// Test Add beyond total
	bar3 := NewProgress(10)
	bar3.Add(20)
	if bar3.current != 10 {
		t.Errorf("Add should cap at total, got %d", bar3.current)
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		duration time.Duration
		contains string
	}{
		{500 * time.Millisecond, "ms"},
		{2 * time.Second, "s"},
		{90 * time.Second, "m"},
	}

	for _, tt := range tests {
		result := formatDuration(tt.duration)
		if len(result) == 0 {
			t.Errorf("formatDuration(%v) returned empty string", tt.duration)
		}
	}
}
