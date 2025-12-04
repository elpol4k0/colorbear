// style.go
package colorbear

import "fmt"

// Style represents a chainable color and text style builder.
//
// Styles can be chained together to combine multiple colors and effects.
// This allows for complex styling with a fluent, readable API.
//
// Example:
//
//	// Create a style once, use multiple times
//	errorStyle := colorbear.NewStyle().Red().Bold()
//	errorStyle.Print("Error 1")
//	errorStyle.Print("Error 2")
//
//	// Or chain inline
//	colorbear.NewStyle().Cyan().Bold().Underline().Print("Important!")
//
// Note: Each method returns a new Style (styles are effectively immutable),
// so you can safely branch from base styles:
//
//	base := colorbear.NewStyle().Bold()
//	redBold := base.Red()    // Bold + Red
//	greenBold := base.Green() // Bold + Green
type Style struct {
	codes []string
}

// NewStyle creates a new empty Style.
//
// You can then chain color and style methods:
//
//	style := colorbear.NewStyle()
//	style.Red().Bold().Print("text")
func NewStyle() *Style {
	return &Style{
		codes: []string{},
	}
}

// Color methods (chainable)
//
// These methods add foreground colors to the style.
// They can be combined with other colors (last one wins)
// and style effects (Bold, Underline, etc.).

// Red adds red foreground color.
//
// Example:
//
//	colorbear.NewStyle().Red().Print("Red text")
func (s *Style) Red() *Style {
	s.codes = append(s.codes, RedCode)
	return s
}

// Green adds green foreground color.
//
// Example:
//
//	colorbear.NewStyle().Green().Print("Green text")
func (s *Style) Green() *Style {
	s.codes = append(s.codes, GreenCode)
	return s
}

// Yellow adds yellow foreground color.
//
// Example:
//
//	colorbear.NewStyle().Yellow().Print("Yellow text")
func (s *Style) Yellow() *Style {
	s.codes = append(s.codes, YellowCode)
	return s
}

// Blue adds blue foreground color.
//
// Example:
//
//	colorbear.NewStyle().Blue().Print("Blue text")
func (s *Style) Blue() *Style {
	s.codes = append(s.codes, BlueCode)
	return s
}

// Cyan adds cyan foreground color.
//
// Example:
//
//	colorbear.NewStyle().Cyan().Print("Cyan text")
func (s *Style) Cyan() *Style {
	s.codes = append(s.codes, CyanCode)
	return s
}

// Magenta adds magenta foreground color.
//
// Example:
//
//	colorbear.NewStyle().Magenta().Print("Magenta text")
func (s *Style) Magenta() *Style {
	s.codes = append(s.codes, MagentaCode)
	return s
}

// White adds white foreground color.
//
// Example:
//
//	colorbear.NewStyle().White().Print("White text")
func (s *Style) White() *Style {
	s.codes = append(s.codes, WhiteCode)
	return s
}

// Black adds black foreground color.
//
// Example:
//
//	colorbear.NewStyle().Black().Print("Black text")
func (s *Style) Black() *Style {
	s.codes = append(s.codes, BlackCode)
	return s
}

// Style/Effect methods
//
// These methods add text effects like bold, underline, italic, etc.
// Multiple effects can be combined.

// Bold adds bold/bright text style.
//
// Example:
//
//	colorbear.NewStyle().Red().Bold().Print("Bold red text")
func (s *Style) Bold() *Style {
	s.codes = append(s.codes, Bold)
	return s
}

// Underline adds underlined text style.
//
// Example:
//
//	colorbear.NewStyle().Green().Underline().Print("Underlined green text")
func (s *Style) Underline() *Style {
	s.codes = append(s.codes, Underline)
	return s
}

// Italic adds italic text style.
//
// Note: Not all terminals support italic text.
//
// Example:
//
//	colorbear.NewStyle().Blue().Italic().Print("Italic blue text")
func (s *Style) Italic() *Style {
	s.codes = append(s.codes, Italic)
	return s
}

// Dim adds dim/faint text style.
//
// This reduces the intensity of the text color.
//
// Example:
//
//	colorbear.NewStyle().White().Dim().Print("Dimmed text")
func (s *Style) Dim() *Style {
	s.codes = append(s.codes, Dim)
	return s
}

// Background color methods
//
// These methods add background colors to the style.
// They can be combined with foreground colors and effects.

// BgRed adds red background color.
//
// Example:
//
//	colorbear.NewStyle().White().BgRed().Print("White text on red background")
func (s *Style) BgRed() *Style {
	s.codes = append(s.codes, BgRed)
	return s
}

// BgGreen adds green background color.
//
// Example:
//
//	colorbear.NewStyle().Black().BgGreen().Print("Black text on green background")
func (s *Style) BgGreen() *Style {
	s.codes = append(s.codes, BgGreen)
	return s
}

// BgYellow adds yellow background color.
//
// Example:
//
//	colorbear.NewStyle().Black().BgYellow().Print("Black text on yellow background")
func (s *Style) BgYellow() *Style {
	s.codes = append(s.codes, BgYellow)
	return s
}

// BgBlue adds blue background color.
//
// Example:
//
//	colorbear.NewStyle().White().BgBlue().Print("White text on blue background")
func (s *Style) BgBlue() *Style {
	s.codes = append(s.codes, BgBlue)
	return s
}

// BgCyan adds cyan background color.
//
// Example:
//
//	colorbear.NewStyle().Black().BgCyan().Print("Black text on cyan background")
func (s *Style) BgCyan() *Style {
	s.codes = append(s.codes, BgCyan)
	return s
}

// BgMagenta adds magenta background color.
//
// Example:
//
//	colorbear.NewStyle().White().BgMagenta().Print("White text on magenta background")
func (s *Style) BgMagenta() *Style {
	s.codes = append(s.codes, BgMagenta)
	return s
}

// BgWhite adds white background color.
//
// Example:
//
//	colorbear.NewStyle().Black().BgWhite().Print("Black text on white background")
func (s *Style) BgWhite() *Style {
	s.codes = append(s.codes, BgWhite)
	return s
}

// BgBlack adds black background color.
//
// Example:
//
//	colorbear.NewStyle().White().BgBlack().Print("White text on black background")
func (s *Style) BgBlack() *Style {
	s.codes = append(s.codes, BgBlack)
	return s
}

// Apply applies the style to text and returns the styled string.
//
// This doesn't print the text, just returns it with ANSI codes applied.
// Use this when you want to build strings or pass them to other functions.
//
// Example:
//
//	style := colorbear.NewStyle().Red().Bold()
//	styledText := style.Apply("Error")
//	fmt.Println(styledText)
func (s *Style) Apply(text string) string {
	return colorize(text, s.codes...)
}

// Print prints the styled text to stdout with a newline.
//
// This is the most common way to output styled text.
//
// Example:
//
//	colorbear.NewStyle().Green().Bold().Print("Success!")
func (s *Style) Print(text string) {
	fmt.Println(s.Apply(text))
}

// Printf prints formatted styled text to stdout with a newline.
//
// This combines style application with fmt.Sprintf formatting.
//
// Example:
//
//	style := colorbear.NewStyle().Red().Bold()
//	style.Printf("Error: %s (code: %d)", msg, code)
func (s *Style) Printf(format string, args ...interface{}) {
	fmt.Println(s.Apply(fmt.Sprintf(format, args...)))
}

// Shortcut functions for direct chaining
//
// These functions create a new Style with a single color already applied.
// They're convenience functions that save you from writing NewStyle().Color().
//
// Example:
//
//	colorbear.RedStyle().Bold().Print("text")
//
// is equivalent to:
//
//	colorbear.NewStyle().Red().Bold().Print("text")
//
// Use whichever reads better in your code.

// RedStyle creates a new red style.
//
// This is a shortcut for NewStyle().Red()
//
// Example:
//
//	colorbear.RedStyle().Bold().Print("Bold red text")
func RedStyle() *Style {
	return NewStyle().Red()
}

// GreenStyle creates a new green style.
//
// This is a shortcut for NewStyle().Green()
//
// Example:
//
//	colorbear.GreenStyle().Bold().Print("Bold green text")
func GreenStyle() *Style {
	return NewStyle().Green()
}

// YellowStyle creates a new yellow style.
//
// This is a shortcut for NewStyle().Yellow()
//
// Example:
//
//	colorbear.YellowStyle().Bold().Print("Bold yellow text")
func YellowStyle() *Style {
	return NewStyle().Yellow()
}

// BlueStyle creates a new blue style.
//
// This is a shortcut for NewStyle().Blue()
//
// Example:
//
//	colorbear.BlueStyle().Bold().Print("Bold blue text")
func BlueStyle() *Style {
	return NewStyle().Blue()
}

// CyanStyle creates a new cyan style.
//
// This is a shortcut for NewStyle().Cyan()
//
// Example:
//
//	colorbear.CyanStyle().Bold().Print("Bold cyan text")
func CyanStyle() *Style {
	return NewStyle().Cyan()
}

// MagentaStyle creates a new magenta style.
//
// This is a shortcut for NewStyle().Magenta()
//
// Example:
//
//	colorbear.MagentaStyle().Bold().Print("Bold magenta text")
func MagentaStyle() *Style {
	return NewStyle().Magenta()
}

// WhiteStyle creates a new white style.
//
// This is a shortcut for NewStyle().White()
//
// Example:
//
//	colorbear.WhiteStyle().Bold().Print("Bold white text")
func WhiteStyle() *Style {
	return NewStyle().White()
}

// BlackStyle creates a new black style.
//
// This is a shortcut for NewStyle().Black()
//
// Example:
//
//	colorbear.BlackStyle().Bold().Print("Bold black text")
func BlackStyle() *Style {
	return NewStyle().Black()
}
