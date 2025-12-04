package colorbear

import "fmt"

// Semantic color functions provide meaningful, intent-based coloring
// instead of technical color names. This makes code more readable and
// maintainable.
//
// IMPORTANT: For use in tables, use the TableSuccess/TableError/etc. variants
// which use ASCII text instead of Unicode symbols for consistent width.

// ============================================================================
// CONSOLE VERSIONS - Use Unicode symbols (for normal printing)
// ============================================================================

// Success returns a success message in green with a checkmark.
//
// Example:
//
//	colorbear.SuccessPrint("Deployment completed")
//	// Output: ✓ Deployment completed (in green)
func Success(text string) string {
	if text == "" {
		return colorize("✓", GreenCode)
	}
	return colorize("✓ "+text, GreenCode)
}

// SuccessPrint prints a success message
func SuccessPrint(text string) {
	fmt.Println(Success(text))
}

// Successf returns a formatted success message
func Successf(format string, args ...interface{}) string {
	return Success(fmt.Sprintf(format, args...))
}

// Error returns an error message in red with X mark
func Error(text string) string {
	if text == "" {
		return colorize("✗", RedCode, Bold)
	}
	return colorize("✗ "+text, RedCode, Bold)
}

// ErrorPrint prints an error message
func ErrorPrint(text string) {
	fmt.Println(Error(text))
}

// Errorf returns a formatted error message
func Errorf(format string, args ...interface{}) string {
	return Error(fmt.Sprintf(format, args...))
}

// Warning returns a warning message in yellow
func Warning(text string) string {
	if text == "" {
		return colorize("⚠", YellowCode)
	}
	return colorize("⚠ "+text, YellowCode)
}

// WarningPrint prints a warning message
func WarningPrint(text string) {
	fmt.Println(Warning(text))
}

// Warningf returns a formatted warning message
func Warningf(format string, args ...interface{}) string {
	return Warning(fmt.Sprintf(format, args...))
}

// Info returns an info message in cyan
func Info(text string) string {
	if text == "" {
		return colorize("ℹ", CyanCode)
	}
	return colorize("ℹ "+text, CyanCode)
}

// InfoPrint prints an info message
func InfoPrint(text string) {
	fmt.Println(Info(text))
}

// Infof returns a formatted info message
func Infof(format string, args ...interface{}) string {
	return Info(fmt.Sprintf(format, args...))
}

// Debug returns a debug message in gray
func Debug(text string) string {
	if text == "" {
		return colorize("🐛", BrightBlack)
	}
	return colorize("🐛 "+text, BrightBlack)
}

// DebugPrint prints a debug message
func DebugPrint(text string) {
	fmt.Println(Debug(text))
}

// Debugf returns a formatted debug message
func Debugf(format string, args ...interface{}) string {
	return Debug(fmt.Sprintf(format, args...))
}

// ============================================================================
// TABLE-SAFE VERSIONS - Use ASCII text (for tables)
// ============================================================================

// TableSuccess returns a table-safe success message in green.
//
// Uses [OK] instead of ✓ to ensure consistent column width in tables.
// All table-safe functions use exactly 4 ASCII characters.
//
// Example:
//
//	table.AddStyledRow("Task", colorbear.TableSuccess(""), "100%")
//	// Output: [OK] (in green) - exactly 4 characters
//
//	table.AddStyledRow("Task", colorbear.TableSuccess("Complete"), "100%")
//	// Output: [OK] Complete (in green)
func TableSuccess(text string) string {
	if text == "" {
		return colorize("[OK]", GreenCode)
	}
	return colorize("[OK] "+text, GreenCode)
}

// TableError returns a table-safe error message in red.
//
// Uses [ERR] instead of ✗ for consistent width.
func TableError(text string) string {
	if text == "" {
		return colorize("[ERR]", RedCode, Bold)
	}
	return colorize("[ERR] "+text, RedCode, Bold)
}

// TableWarning returns a table-safe warning message in yellow.
//
// Uses [!] instead of ⚠ for consistent width.
func TableWarning(text string) string {
	if text == "" {
		return colorize("[!]", YellowCode)
	}
	return colorize("[!] "+text, YellowCode)
}

// TableInfo returns a table-safe info message in cyan.
//
// Uses [i] instead of ℹ for consistent width.
func TableInfo(text string) string {
	if text == "" {
		return colorize("[i]", CyanCode)
	}
	return colorize("[i] "+text, CyanCode)
}

// TableDebug returns a table-safe debug message in gray.
//
// Uses [#] instead of 🐛 for consistent width.
func TableDebug(text string) string {
	if text == "" {
		return colorize("[#]", BrightBlack)
	}
	return colorize("[#] "+text, BrightBlack)
}
