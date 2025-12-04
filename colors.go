package colorbear

// ANSI Color Codes
//
// These constants define the ANSI escape sequences used for terminal colors.
// They follow the standard ANSI/VT100 terminal specification.
//
// Note: Color constants use the "Code" suffix to avoid naming conflicts
// with color functions (e.g., RedCode vs Red()).
const (
	Reset = "\033[0m"

	// Foreground Colors (normal intensity)
	RedCode     = "\033[31m"
	GreenCode   = "\033[32m"
	YellowCode  = "\033[33m"
	BlueCode    = "\033[34m"
	MagentaCode = "\033[35m"
	CyanCode    = "\033[36m"
	WhiteCode   = "\033[37m"
	BlackCode   = "\033[30m"

	// Bright/Bold Colors (high intensity)
	BrightBlack   = "\033[90m" // Dark gray
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// Background Colors
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	// Text Styles
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m" // Swap foreground and background
	Hidden    = "\033[8m" // Concealed/hidden text
)

// colorize applies ANSI color codes to text.
// It automatically respects the color detection settings and returns
// plain text if colors are disabled.
func colorize(text string, codes ...string) string {
	if !isColorEnabled() {
		return text
	}

	var result string
	for _, code := range codes {
		result += code
	}
	result += text + Reset
	return result
}
