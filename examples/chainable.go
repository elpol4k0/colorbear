//go:build ignore
// +build ignore

// examples/chainable.go demonstrates the chainable Style API.
//
// This example shows:
//   - Combining colors and styles (Bold, Underline, etc.)
//   - Using NewStyle() vs XxxStyle() shortcuts
//   - Reusing styles for consistent formatting
//   - Complex multi-effect combinations
//
// Run with: go run examples/chainable.go
package main

import "github.com/elpol4k0/colorbear"

func main() {
	// Method 1: Using NewStyle()
	// Start with NewStyle() and chain methods
	colorbear.NewStyle().Red().Bold().Print("Bold Red Text")
	colorbear.NewStyle().Green().Underline().Print("Underlined Green")
	colorbear.NewStyle().Yellow().BgBlue().Bold().Print("Bold Yellow on Blue Background")
	colorbear.NewStyle().Cyan().Bold().Underline().Print("Bold & Underlined Cyan")

	println()

	// Method 2: Using XxxStyle() shortcuts
	// These are shorter when you start with a color
	colorbear.RedStyle().Bold().Print("Bold Red Text")
	colorbear.GreenStyle().Underline().Print("Underlined Green")
	colorbear.YellowStyle().BgBlue().Bold().Print("Bold Yellow on Blue Background")
	colorbear.CyanStyle().Bold().Underline().Print("Bold & Underlined Cyan")

	println()

	// Best Practice: Save and reuse styles
	// This ensures consistent formatting throughout your app
	errorStyle := colorbear.RedStyle().Bold()
	errorStyle.Print("Error 1: Connection timeout")
	errorStyle.Print("Error 2: File not found")
	errorStyle.Print("Error 3: Permission denied")

	println()

	// Complex combinations
	// You can combine colors, backgrounds, and multiple styles
	colorbear.NewStyle().
		Cyan().
		Bold().
		Underline().
		BgWhite().
		Print("🐻 ColorBear is awesome!")
}
