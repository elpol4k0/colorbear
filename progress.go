// progress.go
package colorbear

import (
	"fmt"
	"strings"
	"time"
)

// ProgressBar represents an animated progress bar for terminal output.
//
// It provides a visual indication of progress for long-running tasks,
// with customizable appearance and optional statistics display.
//
// Progress bars automatically handle terminal width, color detection,
// and gracefully degrade in non-TTY environments.
type ProgressBar struct {
	total       int       // Total number of items to process
	current     int       // Current progress (0 to total)
	width       int       // Width of the progress bar in characters
	prefix      string    // Text to display before the bar
	showPercent bool      // Whether to show percentage
	showCount   bool      // Whether to show current/total count
	showTime    bool      // Whether to show elapsed time
	startTime   time.Time // When the progress bar was created
	lastDraw    string    // Last drawn output (for clearing)
	color       string    // ANSI color code for the filled portion
	completeMsg string    // Message to show on completion (unused currently)
}

// ProgressOption is a functional option for configuring a ProgressBar.
//
// Use the With* functions to customize progress bar behavior.
type ProgressOption func(*ProgressBar)

// NewProgress creates a new progress bar with the specified total.
//
// The total represents the number of items or steps to complete.
// Progress is updated using Set(), Increment(), or Add() methods.
//
// By default, the progress bar:
//   - Shows percentage (disable with WithPercent(false))
//   - Uses cyan color (change with WithColor())
//   - Has width of 40 characters (change with WithWidth())
//   - Has no prefix text (add with WithPrefix())
//
// Example:
//
//	bar := colorbear.NewProgress(100)
//	for i := 0; i <= 100; i++ {
//	    bar.Set(i)
//	    time.Sleep(50 * time.Millisecond)
//	}
//	bar.Finish("Done!")
//
// Example with options:
//
//	bar := colorbear.NewProgress(1000,
//	    colorbear.WithPrefix("Processing:"),
//	    colorbear.WithCount(true),
//	    colorbear.WithTime(true),
//	)
func NewProgress(total int, opts ...ProgressOption) *ProgressBar {
	pb := &ProgressBar{
		total:       total,
		current:     0,
		width:       40,
		prefix:      "",
		showPercent: true,
		showCount:   false,
		showTime:    false,
		startTime:   time.Now(),
		color:       CyanCode,
	}

	// Apply all provided options
	for _, opt := range opts {
		opt(pb)
	}

	return pb
}

// WithWidth sets the width of the progress bar in characters.
//
// The default width is 40 characters. Larger widths provide more granular
// visual feedback, while smaller widths are better for narrow terminals.
//
// Example:
//
//	bar := colorbear.NewProgress(100, colorbear.WithWidth(60))
func WithWidth(width int) ProgressOption {
	return func(pb *ProgressBar) {
		pb.width = width
	}
}

// WithPrefix sets a prefix text to display before the progress bar.
//
// The prefix is useful for indicating what task is being performed.
//
// Example:
//
//	bar := colorbear.NewProgress(100, colorbear.WithPrefix("Downloading:"))
func WithPrefix(prefix string) ProgressOption {
	return func(pb *ProgressBar) {
		pb.prefix = prefix
	}
}

// WithPercent controls whether to show the percentage.
//
// Percentage is shown by default. Set to false to hide it.
//
// Example:
//
//	bar := colorbear.NewProgress(100, colorbear.WithPercent(false))
func WithPercent(show bool) ProgressOption {
	return func(pb *ProgressBar) {
		pb.showPercent = show
	}
}

// WithCount shows the current/total count alongside the progress bar.
//
// This displays the numeric progress in the format "(current/total)".
//
// Example:
//
//	bar := colorbear.NewProgress(1000, colorbear.WithCount(true))
//	// Output: [████████░░░░] 50% (500/1000)
func WithCount(show bool) ProgressOption {
	return func(pb *ProgressBar) {
		pb.showCount = show
	}
}

// WithTime shows the elapsed time since the progress bar was created.
//
// Time is displayed in a human-readable format (e.g., "2.3s", "1m30s").
//
// Example:
//
//	bar := colorbear.NewProgress(100, colorbear.WithTime(true))
//	// Output: [████████████] 100% - 5.2s
func WithTime(show bool) ProgressOption {
	return func(pb *ProgressBar) {
		pb.showTime = show
	}
}

// WithColor sets the color of the filled portion of the progress bar.
//
// Use one of the ColorCode constants (RedCode, GreenCode, etc.).
//
// Example:
//
//	bar := colorbear.NewProgress(100, colorbear.WithColor(colorbear.GreenCode))
func WithColor(color string) ProgressOption {
	return func(pb *ProgressBar) {
		pb.color = color
	}
}

// Set updates the progress bar to a specific value.
//
// The value should be between 0 and total (inclusive).
// This is the most common way to update progress.
//
// Example:
//
//	bar := colorbear.NewProgress(100)
//	for i := 0; i <= 100; i++ {
//	    bar.Set(i)
//	    time.Sleep(50 * time.Millisecond)
//	}
func (pb *ProgressBar) Set(current int) {
	pb.current = current
	pb.draw()
}

// Increment increases the progress by 1.
//
// This is a convenience method equivalent to Add(1).
// The progress is automatically capped at the total.
//
// Example:
//
//	bar := colorbear.NewProgress(100)
//	for i := 0; i < 100; i++ {
//	    bar.Increment()
//	    processItem(i)
//	}
func (pb *ProgressBar) Increment() {
	pb.current++
	if pb.current > pb.total {
		pb.current = pb.total
	}
	pb.draw()
}

// Add increases the progress by a specific amount.
//
// This is useful when processing items in batches.
// The progress is automatically capped at the total.
//
// Example:
//
//	bar := colorbear.NewProgress(1000)
//	for batch := range batches {
//	    processBatch(batch)
//	    bar.Add(len(batch)) // Add batch size to progress
//	}
func (pb *ProgressBar) Add(amount int) {
	pb.current += amount
	if pb.current > pb.total {
		pb.current = pb.total
	}
	pb.draw()
}

// draw renders the progress bar to the terminal.
//
// This method is called internally by Set(), Increment(), and Add().
// It handles both colored and non-colored output based on terminal capabilities.
func (pb *ProgressBar) draw() {
	// Simple fallback for non-TTY environments (piped output, CI/CD, etc.)
	if !isColorEnabled() {
		percent := float64(pb.current) / float64(pb.total) * 100
		fmt.Printf("\r%s%.0f%% (%d/%d)", pb.prefix, percent, pb.current, pb.total)
		return
	}

	// Calculate how much of the bar should be filled
	percent := float64(pb.current) / float64(pb.total) * 100
	filled := int(float64(pb.width) * float64(pb.current) / float64(pb.total))

	// Build the progress bar string
	var bar strings.Builder

	// Add prefix if provided
	if pb.prefix != "" {
		bar.WriteString(pb.prefix)
		bar.WriteString(" ")
	}

	// Draw the bar itself with filled and empty portions
	bar.WriteString("[")
	bar.WriteString(colorize(strings.Repeat("█", filled), pb.color))
	bar.WriteString(strings.Repeat("░", pb.width-filled))
	bar.WriteString("]")

	// Add percentage if enabled
	if pb.showPercent {
		bar.WriteString(fmt.Sprintf(" %.0f%%", percent))
	}

	// Add count if enabled
	if pb.showCount {
		bar.WriteString(fmt.Sprintf(" (%d/%d)", pb.current, pb.total))
	}

	// Add elapsed time if enabled
	if pb.showTime {
		elapsed := time.Since(pb.startTime)
		bar.WriteString(fmt.Sprintf(" - %s", formatDuration(elapsed)))
	}

	// Print the bar, clearing any leftover characters from previous draw
	output := bar.String()
	fmt.Print("\r" + output + strings.Repeat(" ", maxInt(0, len(pb.lastDraw)-len(output))))
	pb.lastDraw = output
}

// Finish completes the progress bar and displays a success message.
//
// This method sets progress to 100%, prints a newline, and displays
// the provided message using the Success style (green with checkmark).
//
// Always call Finish() or FinishWithError() when done to ensure
// proper terminal output.
//
// Example:
//
//	bar := colorbear.NewProgress(100)
//	for i := 0; i <= 100; i++ {
//	    bar.Set(i)
//	    processItem(i)
//	}
//	bar.Finish("All items processed!")
func (pb *ProgressBar) Finish(message string) {
	pb.current = pb.total
	pb.draw()
	fmt.Println() // Move to new line

	if message != "" {
		SuccessPrint(message)
	}
}

// FinishWithError completes the progress bar and displays an error message.
//
// This method prints a newline and displays the provided message
// using the Error style (red and bold with X mark).
//
// Use this when a task fails partway through.
//
// Example:
//
//	bar := colorbear.NewProgress(100)
//	for i := 0; i <= 100; i++ {
//	    bar.Set(i)
//	    if err := processItem(i); err != nil {
//	        bar.FinishWithError("Processing failed: " + err.Error())
//	        return
//	    }
//	}
func (pb *ProgressBar) FinishWithError(message string) {
	fmt.Println() // Move to new line
	ErrorPrint(message)
}

// formatDuration formats a duration in a human-readable way.
//
// Returns:
//   - Milliseconds for durations under 1 second (e.g., "500ms")
//   - Seconds with one decimal for durations under 1 minute (e.g., "2.3s")
//   - Minutes and seconds for longer durations (e.g., "1m30s")
func formatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%dm%ds", minutes, seconds)
}

// maxInt returns the maximum of two integers.
//
// This is a utility function used for clearing previous output.
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
