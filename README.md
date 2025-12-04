# ColorBear

[![Go Reference](https://pkg.go.dev/badge/github.com/elpol4k0/colorbear.svg)](https://pkg.go.dev/github.com/elpol4k0/colorbear)
[![Go Report Card](https://goreportcard.com/badge/github.com/elpol4k0/colorbear)](https://goreportcard.com/report/github.com/elpol4k0/colorbear)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Terminal Colors, The Friendly Way**

ColorBear provides beautiful, semantic terminal colors for Go. Make your CLI applications colorful and easy to read with a simple, chainable API.

## Features

- Simple API - Easy to use, hard to mess up
- Semantic Colors - Success, Error, Warning, Info
- Chainable Styles - Combine colors and styles fluently
- Progress Bars - Animated progress tracking with customization
- Spinners - Smooth loading animations for indeterminate tasks
- Tables - Beautiful tabular data display with 6 visual styles
- Smart Detection - Works in CI/CD, Windows, and everywhere
- Zero Dependencies - Pure Go, no external packages
- Lightweight - Small footprint, fast performance

## Installation
```bash
go get github.com/elpol4k0/colorbear
```

## Quick Start
```go
package main

import "github.com/elpol4k0/colorbear"

func main() {
	// Semantic functions
	colorbear.Success("Deployment completed")
	colorbear.Error("Connection failed")

	// Tables (use Table* variants!)
	table := colorbear.NewTable()
	table.SetHeaders("Task", "Status")
	table.AddStyledRow("Build", colorbear.TableSuccess("Complete"))
	table.Print()

	// Progress & Spinners
	bar := colorbear.NewProgress(100)
	spinner := colorbear.NewSpinner("Loading...")
}
```

See [examples/](examples/) directory for complete usage examples.

## Documentation

### Basic Colors
```go
colorbear.Red("text")
colorbear.Green("text")
colorbear.Yellow("text")
colorbear.Blue("text")
colorbear.Cyan("text")
colorbear.Magenta("text")
colorbear.White("text")
colorbear.Black("text")
```

Each color also has a `Print` variant for direct output:
```go
colorbear.RedPrint("text")
colorbear.GreenPrint("text")
// ... etc
```

### Semantic Functions

Semantic functions provide meaningful, intent-based coloring with appropriate icons. This is the recommended way to use ColorBear.

#### Console Semantic Functions (for regular output)
```go
colorbear.Success("Operation completed")     // ✓ green with checkmark
colorbear.Error("Something went wrong")      // ✗ red and bold with X
colorbear.Warning("Be careful")              // ⚠ yellow with warning sign
colorbear.Info("Just so you know")           // ℹ cyan with info icon
colorbear.Debug("Debug information")         // 🐛 gray with debug icon
```

Each semantic function has both a string-returning and print variant:
```go
// Returns colored string
message := colorbear.Success("Done!")

// Prints directly to stdout
colorbear.SuccessPrint("Done!")
```

#### Table Semantic Functions (for tables)

**IMPORTANT:** When using semantic functions in tables, always use the `Table*` variants for consistent alignment:

```go
colorbear.TableSuccess("text")  // [OK] text   - green
colorbear.TableError("text")    // [ERR] text  - red, bold
colorbear.TableWarning("text")  // [!] text    - yellow
colorbear.TableInfo("text")     // [i] text    - cyan
colorbear.TableDebug("text")    // [#] text    - gray
```

**Why separate functions?** Unicode symbols (✓, ✗, ⚠) can be rendered as either 1 or 2 columns depending on your terminal, causing table misalignment. The `Table*` variants use ASCII text that's always exactly 4 characters wide, ensuring perfect alignment across all terminals.

```go
//DON'T USE in tables (causes alignment issues):
table.AddStyledRow("Task", colorbear.Success("Done"), "100%")

//DO USE in tables (perfect alignment):
table.AddStyledRow("Task", colorbear.TableSuccess("Done"), "100%")
```

### Chainable API

For more complex styling, use the chainable Style API:
```go
// Create styles inline
colorbear.NewStyle().Red().Bold().Print("text")
colorbear.NewStyle().Green().Underline().Print("text")
colorbear.NewStyle().Yellow().BgBlue().Bold().Print("text")

// Use shortcuts for common starting colors
colorbear.RedStyle().Bold().Print("text")
colorbear.GreenStyle().Underline().Print("text")

// Save and reuse styles for consistency
errorStyle := colorbear.RedStyle().Bold()
errorStyle.Print("Error 1: Connection timeout")
errorStyle.Print("Error 2: File not found")
errorStyle.Print("Error 3: Permission denied")

// Build complex combinations
colorbear.NewStyle().
Cyan().
Bold().
Underline().
BgWhite().
Print("Complex styled text")
```

### Tables

Tables provide a clean, organized way to display tabular data in the terminal with customizable styling, colors, and alignment.

#### Basic Table
```go
table := colorbear.NewTable()
table.SetHeaders("Name", "Age", "City")
table.AddRow("Alice", "28", "Berlin")
table.AddRow("Bob", "34", "Munich")
table.AddRow("Charlie", "42", "Hamburg")
table.Print()
```

Output:
```
╭─────────┬─────┬─────────╮
│ Name    │ Age │ City    │
├─────────┼─────┼─────────┤
│ Alice   │ 28  │ Berlin  │
│ Bob     │ 34  │ Munich  │
│ Charlie │ 42  │ Hamburg │
╰─────────┴─────┴─────────╯
```

#### Table with Custom Styling
```go
table := colorbear.NewTable(
colorbear.WithTableStyle(colorbear.TableStyleRounded),
colorbear.WithHeaderColor(colorbear.CyanCode),
colorbear.WithBorderColor(colorbear.GreenCode),
)

table.SetHeaders("Task", "Status", "Progress")
table.AddRow("Build", "Complete", "100%")
table.AddRow("Test", "Running", "67%")
table.Print()
```

#### Table with Styled Rows

**Use `Table*` semantic functions for perfect alignment!**

```go
table := colorbear.NewTable(
colorbear.WithTableStyle(colorbear.TableStyleRounded),
)

table.SetHeaders("Task", "Status", "Time")

// Use Table* variants - they're designed for tables!
table.AddStyledRow(
"Deploy Backend",
colorbear.TableSuccess("Complete"),  // [OK] Complete
"2m 34s",
)

table.AddStyledRow(
"Run Tests",
colorbear.TableSuccess("Passed"),    // [OK] Passed
"5m 12s",
)

table.AddStyledRow(
"Build Frontend",
colorbear.TableWarning("Running"),   // [!] Running
"1m 08s",
)

table.AddStyledRow(
"Update Docs",
colorbear.TableError("Failed"),      // [ERR] Failed
"0s",
)

table.Print()
```

Output:
```
╭────────────────┬─────────────────┬────────╮
│ Task           │ Status          │ Time   │
├────────────────┼─────────────────┼────────┤
│ Deploy Backend │ [OK] Complete   │ 2m 34s │
│ Run Tests      │ [OK] Passed     │ 5m 12s │
│ Build Frontend │ [!] Running     │ 1m 08s │
│ Update Docs    │ [ERR] Failed    │ 0s     │
╰────────────────┴─────────────────┴────────╯
```

#### Table Styles

ColorBear includes 6 predefined table styles:

```go
colorbear.TableStyleRounded  // Modern rounded corners (default)
colorbear.TableStyleDouble   // Bold double lines
colorbear.TableStyleBold     // Thick bold lines
colorbear.TableStyleSimple   // Simple ASCII (works everywhere)
colorbear.TableStyleMinimal  // Minimal horizontal lines only
colorbear.TableStyleCompact  // No borders at all
```

#### Table with Footer
```go
table := colorbear.NewTable(
colorbear.WithTableStyle(colorbear.TableStyleBold),
colorbear.WithFooterColor(colorbear.YellowCode),
colorbear.WithAlignment(
colorbear.AlignLeft,
colorbear.AlignRight,
colorbear.AlignRight,
),
)

table.SetHeaders("Product", "Units", "Revenue")
table.AddRow("Laptop", "45", "$58,455")
table.AddRow("Mouse", "234", "$6,786")
table.AddRow("Keyboard", "156", "$13,884")

table.SetFooter("Total", "435", "$79,125")
table.Print()
```

#### Table with Separators
```go
table := colorbear.NewTable()
table.SetHeaders("Category", "Item", "Price")

table.AddRow("Fruits", "Apple", "$2.99")
table.AddRow("Fruits", "Banana", "$1.49")

table.AddSeparator()  // Visual separator line

table.AddRow("Vegetables", "Carrot", "$1.99")
table.AddRow("Vegetables", "Broccoli", "$2.49")

table.Print()
```

#### Table Options

Available customization options:

- `WithTableStyle(style)` - Visual style (rounded, double, bold, etc.)
- `WithHeaderColor(color)` - Header text color
- `WithBorderColor(color)` - Border and separator color
- `WithFooterColor(color)` - Footer text color
- `WithAlignment(alignments...)` - Column alignment (left, center, right)
- `WithPadding(int)` - Cell padding (default: 1)
- `WithColumnWidths(widths...)` - Fixed column widths
- `WithRowColors(colors...)` - Alternating row colors

#### Alignment Options
```go
table := colorbear.NewTable(
colorbear.WithAlignment(
colorbear.AlignLeft,    // First column: left-aligned
colorbear.AlignRight,   // Second column: right-aligned
colorbear.AlignCenter,  // Third column: center-aligned
),
)
```

#### Table Methods
```go
table.SetHeaders("A", "B", "C")      // Set column headers
table.AddRow("1", "2", "3")          // Add a single row
table.AddRows(row1, row2, row3)      // Add multiple rows
table.AddStyledRow("x", "y", "z")    // Add row with styled cells
table.SetFooter("Total", "100")      // Set footer row
table.AddSeparator()                 // Add separator line
table.Print()                        // Print to stdout
table.Clear()                        // Remove all rows
table.RowCount()                     // Get number of rows
table.ColumnCount()                  // Get number of columns
```

### Progress Bars

Progress bars provide visual feedback for tasks with known duration or item counts.

#### Basic Progress Bar
```go
// Create a progress bar for 100 items
bar := colorbear.NewProgress(100)

for i := 0; i <= 100; i++ {
bar.Set(i)
time.Sleep(50 * time.Millisecond)
}

bar.Finish("Processing complete!")
```

#### Progress Bar with Options
```go
bar := colorbear.NewProgress(1000,
colorbear.WithPrefix("Downloading:"),
colorbear.WithCount(true),           // Show (current/total)
colorbear.WithTime(true),            // Show elapsed time
colorbear.WithWidth(60),             // Custom width
colorbear.WithColor(colorbear.GreenCode), // Custom color
)

for i := 0; i <= 1000; i++ {
bar.Set(i)
// Do work...
}

bar.Finish("Download complete!")
```

#### Progress Bar Methods
```go
bar.Set(50)        // Set to specific value
bar.Increment()    // Increase by 1
bar.Add(10)        // Increase by amount

// Completion
bar.Finish("Success!")
bar.FinishWithError("Failed!")
```

#### Available Options

- `WithPrefix(string)` - Text before the bar
- `WithWidth(int)` - Bar width in characters (default: 40)
- `WithPercent(bool)` - Show/hide percentage (default: true)
- `WithCount(bool)` - Show current/total count
- `WithTime(bool)` - Show elapsed time
- `WithColor(string)` - Bar color (use ColorCode constants)

### Spinners

Spinners provide visual feedback for indeterminate operations (unknown duration).

#### Basic Spinner
```go
spinner := colorbear.NewSpinner("Loading data...")
spinner.Start()

// Do your work
time.Sleep(3 * time.Second)

spinner.Stop("Data loaded!")
```

#### Spinner with Options
```go
spinner := colorbear.NewSpinner("Processing...",
colorbear.WithSpinnerStyle(colorbear.SpinnerCircle),
colorbear.WithSpinnerColor(colorbear.GreenCode),
colorbear.WithSpinnerSpeed(50 * time.Millisecond),
)

spinner.Start()
// Do work...
spinner.Stop("Done!")
```

#### Spinner Styles

Available animation styles:

- `SpinnerDots` - ⠋ ⠙ ⠹ ⠸ (default, smooth and modern)
- `SpinnerLine` - | / - \ (classic, works everywhere)
- `SpinnerCircle` - ◐ ◓ ◑ ◒ (rotating circle)
- `SpinnerArrow` - ← ↖ ↑ ↗ (circular arrow)
- `SpinnerSquare` - ◰ ◳ ◲ ◱ (rotating square)
- `SpinnerBounce` - ⠁ ⠂ ⠄ (bouncing effect)
- `SpinnerGrowDots` - ⣾ ⣽ ⣻ (growing dots)
- `SpinnerPulse` - ● ○ (pulsing dot - ultra smooth!)

#### Dynamic Updates
```go
spinner := colorbear.NewSpinner("Step 1: Initializing...")
spinner.Start()

time.Sleep(2 * time.Second)
spinner.UpdateMessage("Step 2: Processing...")

time.Sleep(2 * time.Second)
spinner.Stop("Complete!")
```

#### Error Handling
```go
spinner := colorbear.NewSpinner("Uploading...")
spinner.Start()

err := upload()
if err != nil {
spinner.StopWithError("Upload failed: " + err.Error())
return
}

spinner.Stop("Upload complete!")
```

#### Available Options

- `WithSpinnerStyle(style)` - Animation style
- `WithSpinnerColor(string)` - Spinner color
- `WithSpinnerSpeed(duration)` - Animation speed (default: 60ms)
- `WithCustomFrames([]string)` - Custom animation frames

### Tables

Tables provide a clean way to display tabular data with customizable styling.

#### Basic Usage
```go
table := colorbear.NewTable()
table.SetHeaders("Name", "Age", "City")
table.AddRow("Alice", "28", "Berlin")
table.Print()
```

#### Table-Safe Semantic Functions

**IMPORTANT:** In tables, always use the `Table*` variants for perfect alignment:

```go
table.AddStyledRow("Task", colorbear.TableSuccess("Complete"), "2m")
table.AddStyledRow("Task", colorbear.TableError("Failed"), "0s")
```

Available functions:
- `TableSuccess(text)` → `[OK] text` (green)
- `TableError(text)` → `[ERR] text` (red)
- `TableWarning(text)` → `[!] text` (yellow)
- `TableInfo(text)` → `[i] text` (cyan)
- `TableDebug(text)` → `[#] text` (gray)

**Why two variants?** Console functions (`Success()`, `Error()`) use Unicode symbols (✓, ✗) which may have inconsistent width. Table functions use ASCII text for reliable alignment.

#### Table Styles

Six predefined styles: `TableStyleSimple`, `TableStyleRounded`, `TableStyleDouble`, `TableStyleBold`, `TableStyleMinimal`, `TableStyleCompact`

#### Available Options

- `WithTableStyle(style)` - Visual style
- `WithHeaderColor(color)` - Header color
- `WithBorderColor(color)` - Border color
- `WithAlignment(...)` - Column alignment
- `WithPadding(int)` - Cell padding
- `WithColumnWidths(...)` - Fixed widths

See [examples/table.go](examples/table.go) for complete usage examples.

### Available Style Methods

**Foreground Colors:**
`Red()`, `Green()`, `Yellow()`, `Blue()`, `Cyan()`, `Magenta()`, `White()`, `Black()`

**Text Effects:**
`Bold()`, `Underline()`, `Italic()`, `Dim()`

**Background Colors:**
`BgRed()`, `BgGreen()`, `BgYellow()`, `BgBlue()`, `BgCyan()`, `BgMagenta()`, `BgWhite()`, `BgBlack()`

**Output Methods:**
- `Apply(text)` - Returns styled string without printing
- `Print(text)` - Prints styled text with newline
- `Printf(format, args...)` - Prints formatted styled text

### Printf-Style Formatting

All semantic functions have formatting variants:
```go
count := 42
colorbear.Successf("Processed %d files", count)

host := "example.com"
colorbear.Errorf("Failed to connect to %s", host)

timeout := 30
colorbear.Warningf("Request timeout after %d seconds", timeout)

port := 8080
colorbear.Infof("Server listening on port %d", port)
```

## Color Detection

ColorBear automatically detects whether colors are supported and disables them when:

- Output is piped to a file or another program
- Running in CI/CD environments (when `CI` environment variable is set)
- `NO_COLOR` environment variable is set (following the [NO_COLOR standard](https://no-color.org/))
- Terminal doesn't support ANSI colors
- On Windows, when the terminal doesn't support ANSI escape codes

### Manual Override

You can manually override the automatic color detection:
```go
// Force colors on (useful for testing or debugging)
colorbear.ForceColors(true)

// Force colors off
colorbear.ForceColors(false)
```

## Examples

The [examples](examples/) directory contains complete usage examples:

- **basic.go** - Basic colors and semantic functions
- **chainable.go** - Chainable style API
- **progress.go** - Progress bar usage
- **spinner.go** - Spinner animations
- **table.go** - Table formatting and styles

Run any example:
```bash
go run examples/table.go
```

## Best Practices

### Use Semantic Functions

Instead of technical color names, use semantic functions that convey meaning:
```go
// Good - conveys meaning
colorbear.Success("User created successfully")
colorbear.Error("Failed to connect to database")
colorbear.Warning("API key will expire in 7 days")

// Works, but less clear
colorbear.Green("User created successfully")
colorbear.Red("Failed to connect to database")
colorbear.Yellow("API key will expire in 7 days")
```

### Use Table* Functions in Tables

Always use `Table*` semantic functions when working with tables:
```go
// ✅ Perfect alignment
table.AddStyledRow("Task", colorbear.TableSuccess("Done"), "2m")

// ❌ May cause misalignment
table.AddStyledRow("Task", colorbear.Success("Done"), "2m")
```

### Choose the Right Feedback Method

- **Semantic Functions** - Simple status messages
- **Tables** - Structured data display
- **Progress Bars** - Tasks with known item count or percentage
- **Spinners** - Tasks with unknown duration (API calls, network requests)
```go
// Known duration/count - use progress bar
bar := colorbear.NewProgress(files.Length)
for _, file := range files {
processFile(file)
bar.Increment()
}
bar.Finish("All files processed!")

// Unknown duration - use spinner
spinner := colorbear.NewSpinner("Connecting to API...")
spinner.Start()
data := fetchFromAPI()
spinner.Stop("Connected!")

// Structured data - use table
table := colorbear.NewTable()
table.SetHeaders("Name", "Status", "Time")
table.AddStyledRow("Build", colorbear.TableSuccess("Pass"), "2m")
table.Print()
```

### Reuse Styles

For consistent formatting, save and reuse styles:
```go
// Define your styles once
var (
errorStyle   = colorbear.RedStyle().Bold()
successStyle = colorbear.GreenStyle()
headerStyle  = colorbear.CyanStyle().Bold().Underline()
)

// Use throughout your application
errorStyle.Print("Connection failed")
successStyle.Print("Deployment complete")
headerStyle.Print("Build Summary")
```

### Respect Color Detection

Don't force colors in production code unless you have a specific reason. Let ColorBear handle detection automatically:
```go
// Good - respects user environment
colorbear.Success("Done")

// Avoid - might break in some environments
colorbear.ForceColors(true)
colorbear.Success("Done")
```


## API Documentation

Full API documentation is available at [pkg.go.dev/github.com/elpol4k0/colorbear](https://pkg.go.dev/github.com/elpol4k0/colorbear).

## Contributing

Contributions are welcome! Please feel free to:

- Report bugs by opening an issue
- Suggest new features or improvements
- Submit pull requests

When contributing, please:
- Add tests for new features
- Update documentation as needed

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Acknowledgments

ColorBear is inspired by popular terminal color libraries from other ecosystems:
- [chalk](https://github.com/chalk/chalk) (Node.js)
- [colorama](https://github.com/tartley/colorama) (Python)
- [fatih/color](https://github.com/fatih/color) (Go)

ColorBear aims to provide a simple, semantic API with smart defaults for the Go ecosystem.

---

Made by [elpol4k0](https://github.com/elpol4k0)