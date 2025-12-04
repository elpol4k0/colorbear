//go:build ignore
// +build ignore

// examples/spinner.go demonstrates ColorBear's spinner functionality.
//
// This example shows various ways to use spinners:
//   - Simple spinner with default settings
//   - Different animation styles (including new SpinnerPulse)
//   - Custom colors
//   - Error handling
//   - Updating messages during animation
//   - Different animation speeds
//   - Real-world multi-step scenarios
//
// Run with: go run examples/spinner.go
package main

import (
	"time"

	"github.com/elpol4k0/colorbear"
)

func main() {
	println("ColorBear Spinner Examples")
	println("==========================\n")

	// Example 1: Simple spinner with default settings
	println("Example 1: Simple Spinner (60ms - smooth!)")
	simpleSpinner()

	// Example 2: Different animation styles
	println("\nExample 2: Different Styles")
	differentStyles()

	// Example 3: Custom colors
	println("\nExample 3: Custom Colors")
	customColors()

	// Example 4: Error handling
	println("\nExample 4: Error Handling")
	errorHandling()

	// Example 5: Updating message during animation
	println("\nExample 5: Dynamic Message Updates")
	dynamicMessage()

	// Example 6: Different speeds
	println("\nExample 6: Animation Speed Comparison")
	differentSpeeds()

	// Example 7: Real-world scenario
	println("\nExample 7: Multi-Step Process")
	multiStepProcess()

	// Example 8: New smooth pulse spinner
	println("\nExample 8: Pulse Spinner (Extra Smooth!)")
	pulseSpinner()
}

// simpleSpinner demonstrates the most basic spinner usage.
//
// This creates a spinner with default settings:
//   - Dots animation style
//   - Cyan color
//   - 60ms speed (smooth, fluid motion)
//
// The default speed has been optimized for maximum smoothness.
func simpleSpinner() {
	// Create a spinner with a message
	spinner := colorbear.NewSpinner("Loading data...")

	// Start the animation
	spinner.Start()

	// Simulate work (replace with actual operation)
	time.Sleep(3 * time.Second)

	// Stop with success message
	spinner.Stop("Data loaded successfully!")
}

// differentStyles demonstrates all available spinner animation styles.
//
// This shows how different styles look and when to use each one.
// Each style has its own character and feel.
func differentStyles() {
	styles := []struct {
		style colorbear.SpinnerStyle
		name  string
	}{
		{colorbear.SpinnerDots, "Dots (default, modern & smooth)"},
		{colorbear.SpinnerPulse, "Pulse (ultra smooth pulsing!)"},
		{colorbear.SpinnerLine, "Line (classic, universal)"},
		{colorbear.SpinnerCircle, "Circle (rotating, professional)"},
		{colorbear.SpinnerArrow, "Arrow (directional, dynamic)"},
		{colorbear.SpinnerSquare, "Square (bold, geometric)"},
		{colorbear.SpinnerBounce, "Bounce (playful, 8 frames!)"},
		{colorbear.SpinnerGrowDots, "Growing Dots (organic feel)"},
	}

	for _, s := range styles {
		spinner := colorbear.NewSpinner(
			"Style: "+s.name,
			colorbear.WithSpinnerStyle(s.style),
		)

		spinner.Start()
		time.Sleep(2500 * time.Millisecond) // Slightly longer to appreciate smoothness
		spinner.Stop("Done!")
		time.Sleep(300 * time.Millisecond) // Brief pause between examples
	}
}

// customColors demonstrates using different colors for spinners.
//
// Different colors can convey different meanings or match your app's theme.
// Color psychology matters - choose colors that match the operation's mood.
func customColors() {
	colors := []struct {
		color string
		name  string
		use   string
	}{
		{colorbear.CyanCode, "Cyan", "default, neutral"},
		{colorbear.GreenCode, "Green", "success/positive operations"},
		{colorbear.YellowCode, "Yellow", "warning/caution needed"},
		{colorbear.MagentaCode, "Magenta", "creative/fun tasks"},
		{colorbear.BlueCode, "Blue", "professional/calm operations"},
		{colorbear.RedCode, "Red", "urgent/critical tasks"},
	}

	for _, c := range colors {
		spinner := colorbear.NewSpinner(
			"Color: "+c.name+" ("+c.use+")",
			colorbear.WithSpinnerColor(c.color),
		)

		spinner.Start()
		time.Sleep(2 * time.Second)
		spinner.Stop("Done!")
		time.Sleep(300 * time.Millisecond)
	}
}

// errorHandling demonstrates handling errors with spinners.
//
// This shows how to use StopWithError() when an operation fails,
// providing clear feedback to the user with appropriate styling.
func errorHandling() {
	// Simulate a successful operation
	spinner := colorbear.NewSpinner("Connecting to server...")
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop("Connected successfully!")

	time.Sleep(500 * time.Millisecond)

	// Simulate a failed operation
	spinner = colorbear.NewSpinner("Uploading file...")
	spinner.Start()
	time.Sleep(1500 * time.Millisecond)

	// Simulate error after some time
	spinner.StopWithError("Upload failed: Connection timeout")
}

// dynamicMessage demonstrates updating the spinner message while running.
//
// This is useful for multi-step operations where you want to show
// progress through different stages without stopping and restarting
// the spinner animation.
func dynamicMessage() {
	spinner := colorbear.NewSpinner("Step 1: Initializing...")
	spinner.Start()

	// Step 1
	time.Sleep(1500 * time.Millisecond)

	// Step 2 - message updates smoothly
	spinner.UpdateMessage("Step 2: Processing data...")
	time.Sleep(1500 * time.Millisecond)

	// Step 3
	spinner.UpdateMessage("Step 3: Validating results...")
	time.Sleep(1500 * time.Millisecond)

	// Step 4
	spinner.UpdateMessage("Step 4: Finalizing...")
	time.Sleep(1500 * time.Millisecond)

	spinner.Stop("All steps completed successfully!")
}

// differentSpeeds demonstrates how animation speed affects the spinner.
//
// The default speed (60ms) is optimized for smooth, fluid animation.
// You can adjust it based on your preference:
//   - Faster (40ms): More energetic, attention-grabbing
//   - Normal (60ms): Smooth, balanced (recommended)
//   - Slower (100ms+): Calmer, less distracting
func differentSpeeds() {
	// Ultra-fast spinner
	spinner := colorbear.NewSpinner(
		"Ultra-fast animation (40ms) - Very energetic!",
		colorbear.WithSpinnerSpeed(40*time.Millisecond),
	)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop("Ultra-fast done!")

	time.Sleep(500 * time.Millisecond)

	// Fast spinner
	spinner = colorbear.NewSpinner(
		"Fast animation (50ms) - Energetic",
		colorbear.WithSpinnerSpeed(50*time.Millisecond),
	)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop("Fast done!")

	time.Sleep(500 * time.Millisecond)

	// Default spinner (60ms - recommended!)
	spinner = colorbear.NewSpinner(
		"Default animation (60ms) - Smooth & balanced (RECOMMENDED)",
		colorbear.WithSpinnerSpeed(60*time.Millisecond),
	)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop("Default done!")

	time.Sleep(500 * time.Millisecond)

	// Slow spinner
	spinner = colorbear.NewSpinner(
		"Slow animation (100ms) - Calm",
		colorbear.WithSpinnerSpeed(100*time.Millisecond),
	)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop("Slow done!")

	time.Sleep(500 * time.Millisecond)

	// Very slow spinner
	spinner = colorbear.NewSpinner(
		"Very slow animation (150ms) - Very calm",
		colorbear.WithSpinnerSpeed(150*time.Millisecond),
	)
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop("Very slow done!")
}

// multiStepProcess demonstrates a real-world scenario with multiple steps.
//
// This simulates a typical deployment or build process with different stages,
// showing how spinners can provide feedback throughout a complex operation.
// Using different spinner styles for different phases can help convey
// the type of work being done.
func multiStepProcess() {
	// Build phase - use Circle (professional)
	spinner := colorbear.NewSpinner(
		"Building project...",
		colorbear.WithSpinnerStyle(colorbear.SpinnerCircle),
		colorbear.WithSpinnerColor(colorbear.BlueCode),
	)
	spinner.Start()
	time.Sleep(2 * time.Second)

	// Test phase
	spinner.UpdateMessage("Running tests...")
	time.Sleep(2 * time.Second)

	// Deploy phase
	spinner.UpdateMessage("Deploying to production...")
	time.Sleep(2 * time.Second)

	// Verification phase
	spinner.UpdateMessage("Verifying deployment...")
	time.Sleep(1500 * time.Millisecond)

	spinner.Stop("Deployment complete!")
}

// pulseSpinner demonstrates the new ultra-smooth SpinnerPulse style.
//
// The Pulse spinner is specifically designed for maximum smoothness
// with a simple, elegant pulsing animation. Perfect for modern UIs.
func pulseSpinner() {
	spinner := colorbear.NewSpinner(
		"Processing with ultra-smooth pulse animation...",
		colorbear.WithSpinnerStyle(colorbear.SpinnerPulse),
		colorbear.WithSpinnerColor(colorbear.MagentaCode),
	)

	spinner.Start()
	time.Sleep(3 * time.Second)
	spinner.Stop("Pulse animation complete!")

	time.Sleep(500 * time.Millisecond)

	// Show pulse with different colors
	colors := []struct {
		color string
		name  string
	}{
		{colorbear.CyanCode, "Cyan pulse"},
		{colorbear.GreenCode, "Green pulse"},
		{colorbear.YellowCode, "Yellow pulse"},
	}

	for _, c := range colors {
		spinner := colorbear.NewSpinner(
			c.name+" - smooth pulsing...",
			colorbear.WithSpinnerStyle(colorbear.SpinnerPulse),
			colorbear.WithSpinnerColor(c.color),
		)
		spinner.Start()
		time.Sleep(2 * time.Second)
		spinner.Stop("Done!")
		time.Sleep(300 * time.Millisecond)
	}
}
