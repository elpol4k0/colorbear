package colorbear

import (
	"strings"
	"testing"
)

func TestSemanticFunctions(t *testing.T) {
	ForceColors(true)
	defer ForceColors(false)

	tests := []struct {
		name          string
		fn            func(string) string
		input         string
		shouldContain []string
	}{
		{"Success", Success, "done", []string{"✓", "done", GreenCode}},
		{"Error", Error, "failed", []string{"✗", "failed", RedCode, Bold}},
		{"Warning", Warning, "careful", []string{"⚠", "careful", YellowCode}},
		{"Info", Info, "note", []string{"ℹ", "note", CyanCode}},
		{"Debug", Debug, "debug", []string{"🐛", "debug", BrightBlack}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fn(tt.input)

			for _, substr := range tt.shouldContain {
				if !strings.Contains(result, substr) {
					t.Errorf("%s() result should contain %q, got %q", tt.name, substr, result)
				}
			}
		})
	}
}

func TestSemanticFormatting(t *testing.T) {
	ForceColors(true)
	defer ForceColors(false)

	result := Successf("Processed %d files", 42)

	if !strings.Contains(result, "42") {
		t.Errorf("Successf() should contain formatted number, got %q", result)
	}
	if !strings.Contains(result, "✓") {
		t.Errorf("Successf() should contain checkmark, got %q", result)
	}
}
