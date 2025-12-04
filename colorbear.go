// Package colorbear provides beautiful, semantic terminal colors for Go.
//
// ColorBear makes it easy to add colors to your CLI applications with
// a simple, intuitive API. It supports basic colors, semantic functions
// (Success, Error, Warning), and chainable styles.
//
// Basic Usage:
//
//	colorbear.RedPrint("Error message")
//	colorbear.GreenPrint("Success message")
//
// Semantic Functions (Recommended):
//
//	colorbear.Success("Deployment completed")
//	colorbear.Error("Connection failed")
//	colorbear.Warning("Deprecated API")
//	colorbear.Info("Server started")
//
// Chainable Styles:
//
//	colorbear.NewStyle().Red().Bold().Print("Important!")
//	colorbear.RedStyle().Bold().Print("Error")
//
// ColorBear automatically detects if colors are supported and disables
// them in CI/CD environments, when piping output, or when NO_COLOR is set.
package colorbear

import "fmt"

// Version is the current version of ColorBear
const Version = "0.1.0"

// Basic color functions

// Red returns red colored text
func Red(text string) string {
	return colorize(text, RedCode)
}

// Green returns green colored text
func Green(text string) string {
	return colorize(text, GreenCode)
}

// Yellow returns yellow colored text
func Yellow(text string) string {
	return colorize(text, YellowCode)
}

// Blue returns blue colored text
func Blue(text string) string {
	return colorize(text, BlueCode)
}

// Cyan returns cyan colored text
func Cyan(text string) string {
	return colorize(text, CyanCode)
}

// Magenta returns magenta colored text
func Magenta(text string) string {
	return colorize(text, MagentaCode)
}

// White returns white colored text
func White(text string) string {
	return colorize(text, WhiteCode)
}

// Black returns black colored text
func Black(text string) string {
	return colorize(text, BlackCode)
}

// Print functions (print directly to stdout)

// RedPrint prints red colored text
func RedPrint(text string) {
	fmt.Println(Red(text))
}

// GreenPrint prints green colored text
func GreenPrint(text string) {
	fmt.Println(Green(text))
}

// YellowPrint prints yellow colored text
func YellowPrint(text string) {
	fmt.Println(Yellow(text))
}

// BluePrint prints blue colored text
func BluePrint(text string) {
	fmt.Println(Blue(text))
}

// CyanPrint prints cyan colored text
func CyanPrint(text string) {
	fmt.Println(Cyan(text))
}

// MagentaPrint prints magenta colored text
func MagentaPrint(text string) {
	fmt.Println(Magenta(text))
}

// Printf-style functions

// Redf returns formatted red colored text
func Redf(format string, args ...interface{}) string {
	return Red(fmt.Sprintf(format, args...))
}

// Greenf returns formatted green colored text
func Greenf(format string, args ...interface{}) string {
	return Green(fmt.Sprintf(format, args...))
}

// Yellowf returns formatted yellow colored text
func Yellowf(format string, args ...interface{}) string {
	return Yellow(fmt.Sprintf(format, args...))
}

// Bluef returns formatted blue colored text
func Bluef(format string, args ...interface{}) string {
	return Blue(fmt.Sprintf(format, args...))
}

// Cyanf returns formatted cyan colored text
func Cyanf(format string, args ...interface{}) string {
	return Cyan(fmt.Sprintf(format, args...))
}

// Magentaf returns formatted magenta colored text
func Magentaf(format string, args ...interface{}) string {
	return Magenta(fmt.Sprintf(format, args...))
}
