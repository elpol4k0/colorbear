//go:build ignore
// +build ignore

// examples/basic.go demonstrates basic ColorBear usage.
//
// This example shows:
//   - Basic color printing functions
//   - Semantic functions (Success, Error, Warning, Info, Debug)
//   - Printf-style formatting
//
// Run with: go run examples/basic.go
package main

import "github.com/elpol4k0/colorbear"

func main() {
	// Basic color printing
	// Each color has a Print function for direct output
	colorbear.RedPrint("This is red 🔴")
	colorbear.GreenPrint("This is green 🟢")
	colorbear.YellowPrint("This is yellow 🟡")
	colorbear.BluePrint("This is blue 🔵")
	colorbear.CyanPrint("This is cyan 🔷")
	colorbear.MagentaPrint("This is magenta 🟣")

	println()

	// Semantic functions (RECOMMENDED!)
	// These convey meaning, not just color
	colorbear.SuccessPrint("Deployment completed successfully")
	colorbear.ErrorPrint("Connection failed")
	colorbear.WarningPrint("Using deprecated API")
	colorbear.InfoPrint("Server started on port 8080")
	colorbear.DebugPrint("Query: SELECT * FROM users")

	println()

	// Printf-style formatting
	// All semantic functions have 'f' variants for formatting
	files := 42
	duration := "2.3s"
	colorbear.SuccessPrint(colorbear.Successf("Processed %d files in %s", files, duration))
}
