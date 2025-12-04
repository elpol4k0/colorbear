package colorbear

import (
	"strings"
	"testing"
)

func TestStyleChaining(t *testing.T) {
	ForceColors(true)
	defer ForceColors(false)

	style := NewStyle().Red().Bold()
	result := style.Apply("test")

	if !strings.Contains(result, RedCode) {
		t.Error("Style should contain red color code")
	}
	if !strings.Contains(result, Bold) {
		t.Error("Style should contain bold code")
	}
	if !strings.Contains(result, "test") {
		t.Error("Style should contain text")
	}
}

func TestStyleMultipleEffects(t *testing.T) {
	ForceColors(true)
	defer ForceColors(false)

	style := NewStyle().Cyan().Bold().Underline().BgWhite()
	result := style.Apply("text")

	expectedCodes := []string{CyanCode, Bold, Underline, BgWhite}
	for _, code := range expectedCodes {
		if !strings.Contains(result, code) {
			t.Errorf("Style should contain %q", code)
		}
	}
}

func TestStyleShortcuts(t *testing.T) {
	ForceColors(true)
	defer ForceColors(false)

	tests := []struct {
		name string
		fn   func() *Style
		code string
	}{
		{"RedStyle", RedStyle, RedCode},
		{"GreenStyle", GreenStyle, GreenCode},
		{"YellowStyle", YellowStyle, YellowCode},
		{"BlueStyle", BlueStyle, BlueCode},
		{"CyanStyle", CyanStyle, CyanCode},
		{"MagentaStyle", MagentaStyle, MagentaCode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			style := tt.fn()
			result := style.Apply("test")

			if !strings.Contains(result, tt.code) {
				t.Errorf("%s() should contain color code %q", tt.name, tt.code)
			}
		})
	}
}
