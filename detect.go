package colorbear

import (
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	// forceColors allows manual override of color detection.
	// nil means auto-detect, true/false forces colors on/off.
	//
	// IMPORTANT:
	// This variable is accessed by multiple goroutines:
	// - Spinner runs color checks in the background
	// - Tests may call ForceColors() at the same time
	// Without synchronization, this caused a DATA RACE.
	forceColors *bool

	// noColor is true when the NO_COLOR environment variable is set.
	// This follows the NO_COLOR standard (https://no-color.org/).
	noColor = os.Getenv("NO_COLOR") != ""

	// mu protects forceColors from concurrent access.
	// Using RWMutex allows many readers but only one writer.
	mu sync.RWMutex
)

// ForceColors overrides automatic color detection.
//
// This is useful for testing or when you want to force colors
// on or off regardless of the environment.
//
// Example:
//
//	colorbear.ForceColors(true)  // Always use colors
//	colorbear.ForceColors(false) // Never use colors
//
// THREAD-SAFETY:
// Before this change, ForceColors wrote to a global variable,
// while Spinner was reading it in a goroutine → DATA RACE.
// We now protect this with a mutex.
func ForceColors(enabled bool) {
	mu.Lock()
	defer mu.Unlock()
	forceColors = &enabled
}

// isColorEnabled checks if colors should be used.
//
// It considers multiple factors in this priority order:
//  1. User override via ForceColors()
//  2. NO_COLOR environment variable
//  3. Whether stdout is a terminal (TTY)
//  4. CI/CD environment detection
//  5. Platform-specific support (Windows)
//
// THREAD-SAFETY:
// forceColors must be read under an RLock because Spinner Goroutines
// might read this concurrently with ForceColors() modifying it.
func isColorEnabled() bool {
	mu.RLock()
	fc := forceColors // Make a local copy while holding the lock
	mu.RUnlock()

	// User override takes precedence
	if fc != nil {
		return *fc
	}

	// NO_COLOR environment variable (https://no-color.org/)
	if noColor {
		return false
	}

	// Check if stdout is a terminal
	if !isTerminal() {
		return false
	}

	// Disable colors in CI/CD by default
	// (most CI systems don't render colors well)
	if isCI() {
		return false
	}

	// Windows has special ANSI support requirements
	if runtime.GOOS == "windows" {
		return isWindowsColorSupported()
	}

	return true
}

// isTerminal checks if stdout is a terminal (TTY).
//
// When output is piped to a file or another program,
// colors should be disabled.
func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

// isCI checks if we're running in a CI/CD environment.
//
// Most CI systems set CI=true or CI=1.
func isCI() bool {
	ci := os.Getenv("CI")
	return ci == "true" || ci == "1"
}

// isWindowsColorSupported checks if Windows terminal supports ANSI colors.
//
// Modern Windows (Windows 10+) with Windows Terminal or terminals that
// set TERM environment variable support ANSI colors.
func isWindowsColorSupported() bool {
	term := os.Getenv("TERM")
	return strings.Contains(term, "xterm") ||
		strings.Contains(term, "256color") ||
		os.Getenv("WT_SESSION") != "" // Windows Terminal
}
