//go:build ignore
// +build ignore

// examples/progress.go demonstrates ColorBear's progress bar functionality.
//
// This example shows various ways to use progress bars:
//   - Simple progress tracking
//   - Adding prefix text and statistics
//   - Simulating file downloads
//   - Custom colors and widths
//   - Error handling
//
// Run with: go run examples/progress.go
package main

import (
	"time"

	"github.com/elpol4k0/colorbear"
)

func main() {
	println("ColorBear Progress Bar Examples")
	println("================================\n")

	// Example 1: Basic progress bar with default settings
	println("Example 1: Simple Progress Bar")
	simpleProgress()

	// Example 2: Progress bar with additional statistics
	println("\nExample 2: With Prefix and Stats")
	progressWithStats()

	// Example 3: Simulating a real-world use case (file downloads)
	println("\nExample 3: File Download Simulation")
	downloadSimulation()

	// Example 4: Custom appearance (color and width)
	println("\nExample 4: With Custom Color")
	customColorProgress()

	// Example 5: Handling errors during progress
	println("\nExample 5: Error Handling")
	progressWithError()
}

// simpleProgress demonstrates the most basic progress bar usage.
//
// This creates a progress bar with default settings and updates it
// from 0 to 100. The bar uses default cyan color and shows percentage.
func simpleProgress() {
	// Create a progress bar for 100 items
	bar := colorbear.NewProgress(100)

	// Simulate processing 100 items
	for i := 0; i <= 100; i++ {
		bar.Set(i)
		time.Sleep(30 * time.Millisecond) // Simulate work
	}

	// Finish with a success message
	bar.Finish("Complete!")
}

// progressWithStats demonstrates progress bars with additional information.
//
// This shows how to add:
//   - A prefix to indicate what's being processed
//   - Item count (current/total)
//   - Elapsed time
func progressWithStats() {
	// Create a progress bar with prefix, count, and time display
	bar := colorbear.NewProgress(
		1000, // Total items to process
		colorbear.WithPrefix("Processing:"),
		colorbear.WithCount(true), // Show (current/total)
		colorbear.WithTime(true),  // Show elapsed time
	)

	// Simulate processing 1000 items quickly
	for i := 0; i <= 1000; i++ {
		bar.Set(i)
		time.Sleep(5 * time.Millisecond) // Fast processing
	}

	bar.Finish("All items processed!")
}

// downloadSimulation demonstrates a real-world use case: downloading files.
//
// This shows how to:
//   - Use progress bars for batch operations
//   - Update the prefix to show current file being processed
//   - Track progress across multiple items
func downloadSimulation() {
	// List of files to "download"
	files := []string{
		"image1.jpg",
		"image2.jpg",
		"document.pdf",
		"video.mp4",
		"archive.zip",
	}

	// Create initial progress bar
	bar := colorbear.NewProgress(
		len(files),
		colorbear.WithPrefix("Downloading:"),
		colorbear.WithCount(true), // Show which file we're on
	)

	// Process each file
	for i, file := range files {
		// Show progress at start of current file
		bar.Set(i)
		time.Sleep(500 * time.Millisecond) // Simulate download time

		// Create new bar with updated prefix showing current file
		// Note: In a real application, you might want to use a single bar
		// and update just the prefix, but this demonstrates the flexibility
		bar = colorbear.NewProgress(
			len(files),
			colorbear.WithPrefix("Downloading: "+file),
			colorbear.WithCount(true),
		)
		bar.Set(i + 1) // Update to show completed file
	}

	bar.Finish("All files downloaded!")
}

// customColorProgress demonstrates customizing the progress bar appearance.
//
// This shows how to:
//   - Change the color of the filled portion
//   - Adjust the width of the progress bar
//   - Use different colors to indicate different types of operations
func customColorProgress() {
	// Create a green progress bar (good for "success" operations like builds)
	bar := colorbear.NewProgress(
		100,
		colorbear.WithPrefix("Building:"),
		colorbear.WithColor(colorbear.GreenCode), // Green color
		colorbear.WithWidth(50),                  // Wider bar for more detail
	)

	// Simulate a build process
	for i := 0; i <= 100; i++ {
		bar.Set(i)
		time.Sleep(20 * time.Millisecond)
	}

	bar.Finish("Build successful!")
}

// progressWithError demonstrates error handling during progress.
//
// This shows how to:
//   - Handle failures partway through a task
//   - Use FinishWithError() to display error messages
//   - Gracefully stop progress on error
//
// In a real application, you would handle errors from actual operations
// rather than simulating them.
func progressWithError() {
	// Create progress bar for an upload operation
	bar := colorbear.NewProgress(
		100,
		colorbear.WithPrefix("Uploading:"),
		colorbear.WithCount(true),
	)

	// Simulate upload that fails at 50%
	for i := 0; i <= 50; i++ {
		bar.Set(i)
		time.Sleep(30 * time.Millisecond)

		// Simulate an error occurring at 50%
		if i == 50 {
			// Stop the progress bar and show error message
			bar.FinishWithError("Upload failed: Connection timeout")
			return
		}
	}

	// This line won't be reached due to the error above
	bar.Finish("Upload complete!")
}
