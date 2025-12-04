package colorbear

import (
	"testing"
	"time"
)

func TestSpinnerCreation(t *testing.T) {
	spinner := NewSpinner("Loading...")

	if spinner.message != "Loading..." {
		t.Errorf("Expected message 'Loading...', got %q", spinner.message)
	}
	if spinner.speed != 60*time.Millisecond {
		t.Errorf("Expected speed 60ms, got %v", spinner.speed)
	}
	if spinner.color != CyanCode {
		t.Errorf("Expected color %q, got %q", CyanCode, spinner.color)
	}
}

func TestSpinnerOptions(t *testing.T) {
	spinner := NewSpinner("Test",
		WithSpinnerStyle(SpinnerCircle),
		WithSpinnerColor(GreenCode),
		WithSpinnerSpeed(100*time.Millisecond),
	)

	if len(spinner.frames) != 4 { // Circle has 4 frames
		t.Errorf("Expected 4 frames for Circle style, got %d", len(spinner.frames))
	}
	if spinner.color != GreenCode {
		t.Errorf("Expected color %q, got %q", GreenCode, spinner.color)
	}
	if spinner.speed != 100*time.Millisecond {
		t.Errorf("Expected speed 100ms, got %v", spinner.speed)
	}
}

func TestSpinnerStyles(t *testing.T) {
	styles := []struct {
		style      SpinnerStyle
		frameCount int
		name       string
	}{
		{SpinnerDots, 10, "Dots"},
		{SpinnerLine, 4, "Line"},
		{SpinnerCircle, 4, "Circle"},
		{SpinnerArrow, 8, "Arrow"},
		{SpinnerSquare, 4, "Square"},
		{SpinnerBounce, 8, "Bounce"},
		{SpinnerGrowDots, 8, "GrowDots"},
		{SpinnerPulse, 4, "Pulse"},
	}

	for _, tt := range styles {
		t.Run(tt.name, func(t *testing.T) {
			frames := getSpinnerFrames(tt.style)
			if len(frames) != tt.frameCount {
				t.Errorf("%s: expected %d frames, got %d", tt.name, tt.frameCount, len(frames))
			}
		})
	}
}

func TestSpinnerStartStop(t *testing.T) {
	ForceColors(false) // Disable colors for cleaner test output

	spinner := NewSpinner("Testing...")

	// Test Start
	spinner.Start()
	if !spinner.running {
		t.Error("Spinner should be running after Start()")
	}

	// Let it run briefly
	time.Sleep(200 * time.Millisecond)

	// Test Stop
	spinner.Stop("")
	time.Sleep(50 * time.Millisecond) // Give it time to stop

	if spinner.running {
		t.Error("Spinner should not be running after Stop()")
	}
}

func TestSpinnerUpdateMessage(t *testing.T) {
	spinner := NewSpinner("Original")

	spinner.UpdateMessage("Updated")
	if spinner.message != "Updated" {
		t.Errorf("Expected message 'Updated', got %q", spinner.message)
	}
}

func TestSpinnerCustomFrames(t *testing.T) {
	customFrames := []string{"A", "B", "C"}
	spinner := NewSpinner("Test", WithCustomFrames(customFrames))

	if len(spinner.frames) != 3 {
		t.Errorf("Expected 3 custom frames, got %d", len(spinner.frames))
	}
	if spinner.frames[0] != "A" {
		t.Errorf("Expected first frame 'A', got %q", spinner.frames[0])
	}
}
