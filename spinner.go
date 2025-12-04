// spinner.go
package colorbear

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Spinner represents an animated loading spinner for terminal output.
//
// Spinners provide visual feedback for operations where the duration
// is unknown or indeterminate. They continuously animate until stopped,
// making them ideal for tasks like network requests, file processing,
// or any operation without a known completion time.
//
// Unlike progress bars, spinners don't show percentage completion.
// Use a spinner when you know something is happening but not how long it will take.
type Spinner struct {
	message    string        // Message to display next to the spinner
	frames     []string      // Animation frames (characters to cycle through)
	current    int           // Current frame index
	stop       chan bool     // Channel to signal stop
	running    bool          // Whether the spinner is currently running
	mu         sync.Mutex    // Mutex to prevent race conditions
	color      string        // Color for the spinner frames
	speed      time.Duration // Animation speed (time between frames)
	writer     io.Writer     // Output writer (default: os.Stdout)
	lastOutput string        // Last output for efficient clearing
}

// SpinnerOption is a functional option for configuring a Spinner.
//
// Use the With* functions to customize spinner behavior.
type SpinnerOption func(*Spinner)

// SpinnerStyle represents predefined spinner animation styles
type SpinnerStyle int

const (
	SpinnerDots     SpinnerStyle = iota // ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏ (default, smooth)
	SpinnerLine                         // | / - \ (classic)
	SpinnerArrow                        // ← ↖ ↑ ↗ → ↘ ↓ ↙ (circular arrow)
	SpinnerBounce                       // ⠁ ⠂ ⠄ ⠂ (bouncing)
	SpinnerCircle                       // ◐ ◓ ◑ ◒ (rotating circle)
	SpinnerSquare                       // ◰ ◳ ◲ ◱ (rotating square)
	SpinnerGrowDots                     // ⣾ ⣽ ⣻ ⢿ ⡿ ⣟ ⣯ ⣷ (growing dots)
	SpinnerPulse                        // ● ○ ● ○ (pulsing dot - smooth!)
)

// getSpinnerFrames returns the animation frames for a given style
func getSpinnerFrames(style SpinnerStyle) []string {
	switch style {
	case SpinnerDots:
		// Extended frames for smoother animation
		return []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	case SpinnerLine:
		return []string{"|", "/", "-", "\\"}
	case SpinnerArrow:
		return []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	case SpinnerBounce:
		return []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}
	case SpinnerCircle:
		return []string{"◐", "◓", "◑", "◒"}
	case SpinnerSquare:
		return []string{"◰", "◳", "◲", "◱"}
	case SpinnerGrowDots:
		return []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	case SpinnerPulse:
		// New: smooth pulsing effect
		return []string{"●", "○", "●", "○"}
	default:
		return []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	}
}

// NewSpinner creates a new spinner with the given message.
//
// The spinner uses the default "dots" animation style and cyan color.
// Customize the spinner using SpinnerOption functions.
//
// The default animation speed is optimized for smooth, fluid motion (60ms).
//
// Example:
//
//	spinner := colorbear.NewSpinner("Loading data...")
//	spinner.Start()
//	time.Sleep(3 * time.Second) // Do work
//	spinner.Stop("Data loaded!")
//
// Example with options:
//
//	spinner := colorbear.NewSpinner("Processing...",
//	    colorbear.WithSpinnerStyle(colorbear.SpinnerCircle),
//	    colorbear.WithSpinnerColor(colorbear.GreenCode),
//	)
func NewSpinner(message string, opts ...SpinnerOption) *Spinner {
	s := &Spinner{
		message: message,
		frames:  getSpinnerFrames(SpinnerDots),
		current: 0,
		stop:    make(chan bool, 1), // Buffered to prevent blocking
		running: false,
		color:   CyanCode,
		speed:   60 * time.Millisecond, // Faster for smoother animation (16.6 FPS)
		writer:  os.Stdout,
	}

	// Apply options
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithSpinnerStyle sets the animation style for the spinner.
//
// Available styles:
//   - SpinnerDots: ⠋ ⠙ ⠹ ⠸ (default, smooth and modern)
//   - SpinnerLine: | / - \ (classic, works everywhere)
//   - SpinnerArrow: ← ↖ ↑ ↗ (circular arrow)
//   - SpinnerBounce: ⠁ ⠂ ⠄ (bouncing effect)
//   - SpinnerCircle: ◐ ◓ ◑ ◒ (rotating circle)
//   - SpinnerSquare: ◰ ◳ ◲ ◱ (rotating square)
//   - SpinnerGrowDots: ⣾ ⣽ ⣻ (growing dots)
//   - SpinnerPulse: ● ○ (pulsing dot - very smooth)
//
// Example:
//
//	spinner := colorbear.NewSpinner("Loading...",
//	    colorbear.WithSpinnerStyle(colorbear.SpinnerCircle),
//	)
func WithSpinnerStyle(style SpinnerStyle) SpinnerOption {
	return func(s *Spinner) {
		s.frames = getSpinnerFrames(style)
	}
}

// WithSpinnerColor sets the color of the spinner animation.
//
// Use one of the ColorCode constants (RedCode, GreenCode, etc.).
//
// Example:
//
//	spinner := colorbear.NewSpinner("Processing...",
//	    colorbear.WithSpinnerColor(colorbear.GreenCode),
//	)
func WithSpinnerColor(color string) SpinnerOption {
	return func(s *Spinner) {
		s.color = color
	}
}

// WithSpinnerSpeed sets the animation speed (time between frames).
//
// Default is 60ms for smooth animation. Lower values = faster, higher = slower.
// Recommended range: 40ms (very fast) to 100ms (slow).
//
// Example:
//
//	spinner := colorbear.NewSpinner("Fast loading...",
//	    colorbear.WithSpinnerSpeed(40 * time.Millisecond),
//	)
func WithSpinnerSpeed(speed time.Duration) SpinnerOption {
	return func(s *Spinner) {
		s.speed = speed
	}
}

// WithCustomFrames sets custom animation frames.
//
// Provide your own sequence of characters/strings for animation.
//
// Example:
//
//	spinner := colorbear.NewSpinner("Custom...",
//	    colorbear.WithCustomFrames([]string{"◜", "◝", "◞", "◟"}),
//	)
func WithCustomFrames(frames []string) SpinnerOption {
	return func(s *Spinner) {
		s.frames = frames
	}
}

// Start begins the spinner animation.
//
// The spinner runs in a separate goroutine and animates until Stop()
// or StopWithError() is called. It's safe to call Start() multiple times
// (subsequent calls are ignored if already running).
//
// The animation is optimized for smooth, flicker-free rendering.
//
// Example:
//
//	spinner := colorbear.NewSpinner("Loading...")
//	spinner.Start()
//
//	// Do your work here
//	data := fetchData()
//
//	spinner.Stop("Done!")
func (s *Spinner) Start() {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return // Already running
	}
	s.running = true
	s.mu.Unlock()

	go func() {
		ticker := time.NewTicker(s.speed)
		defer ticker.Stop()

		// Hide cursor for smoother animation
		if isColorEnabled() {
			fmt.Fprint(s.writer, "\033[?25l") // Hide cursor
		}

		for {
			select {
			case <-s.stop:
				// Show cursor again
				if isColorEnabled() {
					fmt.Fprint(s.writer, "\033[?25h") // Show cursor
				}
				return
			case <-ticker.C:
				s.render()
			}
		}
	}()
}

// render draws the current spinner frame (internal method)
func (s *Spinner) render() {
	s.mu.Lock()
	defer s.mu.Unlock()

	frame := s.frames[s.current%len(s.frames)]
	s.current++

	var output string
	if isColorEnabled() {
		output = fmt.Sprintf("\r%s %s", colorize(frame, s.color), s.message)
	} else {
		output = fmt.Sprintf("\r%s %s", frame, s.message)
	}

	// Clear any leftover characters from previous render
	if len(s.lastOutput) > len(output) {
		padding := strings.Repeat(" ", len(s.lastOutput)-len(output))
		output += padding
	}

	fmt.Fprint(s.writer, output)
	s.lastOutput = output
}

// Stop stops the spinner and displays a success message.
//
// This method clears the spinner line and displays the provided
// message using the Success style (green with checkmark).
//
// Always call Stop() or StopWithError() when your operation completes
// to ensure clean terminal output.
//
// Example:
//
//	spinner := colorbear.NewSpinner("Downloading...")
//	spinner.Start()
//
//	err := download()
//	if err != nil {
//	    spinner.StopWithError("Download failed!")
//	    return
//	}
//
//	spinner.Stop("Download complete!")
func (s *Spinner) Stop(message string) {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return // Not running
	}
	s.running = false
	s.mu.Unlock()

	// Signal stop
	select {
	case s.stop <- true:
	default:
		// Channel already has a value or is closed
	}

	// Small delay to ensure goroutine has stopped
	time.Sleep(10 * time.Millisecond)

	// Clear the spinner line completely
	s.mu.Lock()
	clearLength := maxInt(len(s.lastOutput), len(s.message)+10)
	s.mu.Unlock()

	fmt.Fprint(s.writer, "\r"+strings.Repeat(" ", clearLength)+"\r")

	if message != "" {
		SuccessPrint(message)
	}
}

// StopWithError stops the spinner and displays an error message.
//
// This method clears the spinner line and displays the provided
// message using the Error style (red and bold with X mark).
//
// Use this when an operation fails during spinner animation.
//
// Example:
//
//	spinner := colorbear.NewSpinner("Connecting to server...")
//	spinner.Start()
//
//	err := connect()
//	if err != nil {
//	    spinner.StopWithError("Connection failed: " + err.Error())
//	    return
//	}
//
//	spinner.Stop("Connected!")
func (s *Spinner) StopWithError(message string) {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return // Not running
	}
	s.running = false
	s.mu.Unlock()

	// Signal stop
	select {
	case s.stop <- true:
	default:
	}

	// Small delay to ensure goroutine has stopped
	time.Sleep(10 * time.Millisecond)

	// Clear the spinner line
	s.mu.Lock()
	clearLength := maxInt(len(s.lastOutput), len(s.message)+10)
	s.mu.Unlock()

	fmt.Fprint(s.writer, "\r"+strings.Repeat(" ", clearLength)+"\r")

	ErrorPrint(message)
}

// UpdateMessage updates the spinner message while it's running.
//
// This allows you to change the message dynamically without stopping
// and restarting the spinner. The update is applied smoothly on the
// next render cycle.
//
// Example:
//
//	spinner := colorbear.NewSpinner("Step 1: Initializing...")
//	spinner.Start()
//
//	initialize()
//	spinner.UpdateMessage("Step 2: Processing...")
//
//	process()
//	spinner.Stop("Complete!")
func (s *Spinner) UpdateMessage(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.message = message
}
