package colorbear

import (
	"os"
	"testing"
)

func TestForceColors(t *testing.T) {
	// Test enabling
	ForceColors(true)
	if !isColorEnabled() {
		t.Error("Colors should be enabled after ForceColors(true)")
	}

	// Test disabling
	ForceColors(false)
	if isColorEnabled() {
		t.Error("Colors should be disabled after ForceColors(false)")
	}

	// Reset
	forceColors = nil
}

func TestNOCOLOR(t *testing.T) {
	// Save original
	original := os.Getenv("NO_COLOR")
	defer os.Setenv("NO_COLOR", original)

	// Test NO_COLOR set
	os.Setenv("NO_COLOR", "1")
	forceColors = nil // Reset force override

	if isColorEnabled() {
		t.Error("Colors should be disabled when NO_COLOR is set")
	}

	// Test NO_COLOR unset
	os.Unsetenv("NO_COLOR")
	noColor = false // Reset package variable

	// This might be true or false depending on terminal, just check it doesn't panic
	_ = isColorEnabled()
}

func TestCI(t *testing.T) {
	// Save original
	original := os.Getenv("CI")
	defer os.Setenv("CI", original)

	// Test CI environment
	os.Setenv("CI", "true")
	forceColors = nil
	noColor = false

	if isCI() != true {
		t.Error("isCI() should return true when CI=true")
	}

	// Reset
	os.Unsetenv("CI")
}
