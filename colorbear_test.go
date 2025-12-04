package colorbear

import (
	"testing"
)

func TestBasicColors(t *testing.T) {
	// Force colors for testing
	ForceColors(true)
	defer ForceColors(false)

	tests := []struct {
		name     string
		fn       func(string) string
		input    string
		expected string
	}{
		{"Red", Red, "test", RedCode + "test" + Reset},
		{"Green", Green, "test", GreenCode + "test" + Reset},
		{"Yellow", Yellow, "test", YellowCode + "test" + Reset},
		{"Blue", Blue, "test", BlueCode + "test" + Reset},
		{"Cyan", Cyan, "test", CyanCode + "test" + Reset},
		{"Magenta", Magenta, "test", MagentaCode + "test" + Reset},
		{"White", White, "test", WhiteCode + "test" + Reset},
		{"Black", Black, "test", BlackCode + "test" + Reset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.input)
			if result != tt.expected {
				t.Errorf("%s() = %q, want %q", tt.name, result, tt.expected)
			}
		})
	}
}

func TestColorize(t *testing.T) {
	ForceColors(true)
	defer ForceColors(false)

	result := colorize("text", RedCode, Bold)
	expected := RedCode + Bold + "text" + Reset

	if result != expected {
		t.Errorf("colorize() = %q, want %q", result, expected)
	}
}

func TestColorizeDisabled(t *testing.T) {
	ForceColors(false)

	result := colorize("text", RedCode, Bold)
	expected := "text" // No colors when disabled

	if result != expected {
		t.Errorf("colorize() with colors disabled = %q, want %q", result, expected)
	}
}
