package colorbear

import (
	"fmt"
	"os"
	"strings"
)

// Table represents a terminal table with headers, rows, and optional footer.
//
// Tables provide a clean, organized way to display tabular data in the terminal
// with customizable styling, colors, and alignment options.
//
// Basic usage:
//
//	table := colorbear.NewTable()
//	table.SetHeaders("Name", "Age", "City")
//	table.AddRow("Alice", "28", "Berlin")
//	table.AddRow("Bob", "34", "Munich")
//	table.Print()
//
// Advanced usage with options:
//
//	table := colorbear.NewTable(
//	    colorbear.WithTableStyle(colorbear.TableStyleRounded),
//	    colorbear.WithHeaderColor(colorbear.CyanCode),
//	    colorbear.WithAlignment(colorbear.AlignLeft, colorbear.AlignRight, colorbear.AlignCenter),
//	)
type Table struct {
	headers      []string
	rows         [][]string
	footer       []string
	columnWidths []int
	style        *TableStyle
	options      *TableOptions
}

// TableOptions contains configuration for table appearance and behavior.
//
// Options can be set using the With* functional options when creating a table.
type TableOptions struct {
	HeaderColor  string      // Color for header text
	BorderColor  string      // Color for borders and separators
	FooterColor  string      // Color for footer text
	RowColors    []string    // Alternating row colors (cycles through)
	Alignment    []Alignment // Alignment per column
	Padding      int         // Padding inside cells
	ShowBorders  bool        // Whether to show table borders
	ShowHeader   bool        // Whether to show header row
	AutoSize     bool        // Automatically calculate column widths
	MinWidth     int         // Minimum width for all columns
	MaxWidth     int         // Maximum width for all columns
	ColumnWidths []int       // Fixed column widths (overrides auto-sizing)
	Style        *TableStyle // Table style (used internally)
}

// Alignment represents text alignment in a column.
type Alignment int

const (
	AlignLeft   Alignment = iota // Left-aligned text (default)
	AlignCenter                  // Center-aligned text
	AlignRight                   // Right-aligned text
)

// TableStyle defines the border characters for the table.
//
// Different styles create different visual appearances using Unicode box-drawing
// characters. Use one of the predefined styles or create your own.
type TableStyle struct {
	TopLeft         string // Top-left corner
	TopRight        string // Top-right corner
	BottomLeft      string // Bottom-left corner
	BottomRight     string // Bottom-right corner
	Horizontal      string // Horizontal line
	Vertical        string // Vertical line
	Cross           string // Cross intersection
	TopCross        string // Top intersection (T shape)
	BottomCross     string // Bottom intersection (inverted T)
	LeftCross       string // Left intersection
	RightCross      string // Right intersection
	HeaderSeparator string // Line under header
	Name            string // Style name (for debugging)
}

// Predefined table styles
var (
	// TableStyleSimple uses basic ASCII characters.
	// Works everywhere but less visually appealing.
	//
	// Example:
	// +----------+-------+
	// | Name     | Age   |
	// +----------+-------+
	// | Alice    | 28    |
	// +----------+-------+
	TableStyleSimple = &TableStyle{
		TopLeft:         "+",
		TopRight:        "+",
		BottomLeft:      "+",
		BottomRight:     "+",
		Horizontal:      "-",
		Vertical:        "|",
		Cross:           "+",
		TopCross:        "+",
		BottomCross:     "+",
		LeftCross:       "+",
		RightCross:      "+",
		HeaderSeparator: "-",
		Name:            "simple",
	}

	// TableStyleRounded uses rounded corners with Unicode box-drawing characters.
	// This is the default style - clean and modern looking.
	//
	// Example:
	// ╭──────────┬───────╮
	// │ Name     │ Age   │
	// ├──────────┼───────┤
	// │ Alice    │ 28    │
	// ╰──────────┴───────╯
	TableStyleRounded = &TableStyle{
		TopLeft:         "╭",
		TopRight:        "╮",
		BottomLeft:      "╰",
		BottomRight:     "╯",
		Horizontal:      "─",
		Vertical:        "│",
		Cross:           "┼",
		TopCross:        "┬",
		BottomCross:     "┴",
		LeftCross:       "├",
		RightCross:      "┤",
		HeaderSeparator: "─",
		Name:            "rounded",
	}

	// TableStyleDouble uses double-line box-drawing characters.
	// Bold and formal appearance.
	//
	// Example:
	// ╔══════════╦═══════╗
	// ║ Name     ║ Age   ║
	// ╠══════════╬═══════╣
	// ║ Alice    ║ 28    ║
	// ╚══════════╩═══════╝
	TableStyleDouble = &TableStyle{
		TopLeft:         "╔",
		TopRight:        "╗",
		BottomLeft:      "╚",
		BottomRight:     "╝",
		Horizontal:      "═",
		Vertical:        "║",
		Cross:           "╬",
		TopCross:        "╦",
		BottomCross:     "╩",
		LeftCross:       "╠",
		RightCross:      "╣",
		HeaderSeparator: "═",
		Name:            "double",
	}

	// TableStyleBold uses bold single-line box-drawing characters.
	// Thicker lines than standard style.
	//
	// Example:
	// ┏━━━━━━━━━━┳━━━━━━━┓
	// ┃ Name     ┃ Age   ┃
	// ┣━━━━━━━━━━╋━━━━━━━┫
	// ┃ Alice    ┃ 28    ┃
	// ┗━━━━━━━━━━┻━━━━━━━┛
	TableStyleBold = &TableStyle{
		TopLeft:         "┏",
		TopRight:        "┓",
		BottomLeft:      "┗",
		BottomRight:     "┛",
		Horizontal:      "━",
		Vertical:        "┃",
		Cross:           "╋",
		TopCross:        "┳",
		BottomCross:     "┻",
		LeftCross:       "┣",
		RightCross:      "┫",
		HeaderSeparator: "━",
		Name:            "bold",
	}

	// TableStyleMinimal uses minimal borders - only horizontal lines.
	// Clean, spreadsheet-like appearance.
	//
	// Example:
	//  Name       Age
	// ────────────────
	//  Alice      28
	//  Bob        34
	TableStyleMinimal = &TableStyle{
		TopLeft:         "",
		TopRight:        "",
		BottomLeft:      "",
		BottomRight:     "",
		Horizontal:      "─",
		Vertical:        " ",
		Cross:           " ",
		TopCross:        " ",
		BottomCross:     " ",
		LeftCross:       "",
		RightCross:      "",
		HeaderSeparator: "─",
		Name:            "minimal",
	}

	// TableStyleCompact uses no borders at all - only spacing.
	// Ultra-minimal, maximum readability.
	//
	// Example:
	//  Name       Age
	//  Alice      28
	//  Bob        34
	TableStyleCompact = &TableStyle{
		TopLeft:         "",
		TopRight:        "",
		BottomLeft:      "",
		BottomRight:     "",
		Horizontal:      "",
		Vertical:        " ",
		Cross:           " ",
		TopCross:        "",
		BottomCross:     "",
		LeftCross:       "",
		RightCross:      "",
		HeaderSeparator: "",
		Name:            "compact",
	}
)

// TableOption is a functional option for configuring a Table.
//
// Use the With* functions to customize table behavior and appearance.
type TableOption func(*TableOptions)

// WithTableStyle sets the visual style of the table borders.
func WithTableStyle(style *TableStyle) TableOption {
	return func(o *TableOptions) {
		o.Style = style
	}
}

// WithHeaderColor sets the color for the header row.
func WithHeaderColor(color string) TableOption {
	return func(o *TableOptions) {
		o.HeaderColor = color
	}
}

// WithBorderColor sets the color for table borders and separators.
func WithBorderColor(color string) TableOption {
	return func(o *TableOptions) {
		o.BorderColor = color
	}
}

// WithFooterColor sets the color for the footer row.
func WithFooterColor(color string) TableOption {
	return func(o *TableOptions) {
		o.FooterColor = color
	}
}

// WithAlignment sets the alignment for each column.
func WithAlignment(alignments ...Alignment) TableOption {
	return func(o *TableOptions) {
		o.Alignment = alignments
	}
}

// WithPadding sets the padding inside cells.
func WithPadding(padding int) TableOption {
	return func(o *TableOptions) {
		o.Padding = padding
	}
}

// WithColumnWidths sets fixed widths for columns.
func WithColumnWidths(widths ...int) TableOption {
	return func(o *TableOptions) {
		o.ColumnWidths = widths
	}
}

// WithRowColors sets alternating row colors.
func WithRowColors(colors ...string) TableOption {
	return func(o *TableOptions) {
		o.RowColors = colors
	}
}

// WithAutoSize enables or disables automatic column sizing.
func WithAutoSize(enabled bool) TableOption {
	return func(o *TableOptions) {
		o.AutoSize = enabled
	}
}

// NewTable creates a new table with optional configuration.
func NewTable(opts ...TableOption) *Table {
	options := &TableOptions{
		HeaderColor: CyanCode,
		BorderColor: "",
		FooterColor: WhiteCode,
		Padding:     1,
		ShowBorders: true,
		ShowHeader:  true,
		AutoSize:    true,
		Style:       TableStyleRounded, // Default style
	}

	// Apply all options
	for _, opt := range opts {
		opt(options)
	}

	// Use the style from options if it was set
	style := options.Style
	if style == nil {
		style = TableStyleRounded
	}

	return &Table{
		headers:      []string{},
		rows:         [][]string{},
		footer:       []string{},
		columnWidths: []int{},
		style:        style,
		options:      options,
	}
}

// SetHeaders sets the column headers.
func (t *Table) SetHeaders(headers ...string) *Table {
	t.headers = headers
	return t
}

// AddRow adds a new data row to the table.
func (t *Table) AddRow(cells ...string) *Table {
	t.rows = append(t.rows, cells)
	return t
}

// AddRows adds multiple rows at once.
func (t *Table) AddRows(rows ...[]string) *Table {
	t.rows = append(t.rows, rows...)
	return t
}

// SetFooter sets the footer row.
func (t *Table) SetFooter(cells ...string) *Table {
	t.footer = cells
	return t
}

// AddStyledRow adds a row with pre-styled cells.
func (t *Table) AddStyledRow(cells ...string) *Table {
	return t.AddRow(cells...)
}

// AddSeparator adds a visual separator line between rows.
func (t *Table) AddSeparator() *Table {
	t.rows = append(t.rows, []string{"__SEPARATOR__"})
	return t
}

// REFACTORED: calculateColumnWidths - reduced complexity from 22 to <15
// Split into smaller, focused functions

// calculateColumnWidths determines the width of each column.
func (t *Table) calculateColumnWidths() {
	if len(t.columnWidths) > 0 {
		return
	}

	numCols := t.determineColumnCount()
	if numCols == 0 {
		return
	}

	if t.useFixedWidths() {
		return
	}

	t.initializeWidths(numCols)
	t.updateWidthsFromHeaders()
	t.updateWidthsFromRows(numCols)
	t.updateWidthsFromFooter(numCols)
	t.applyWidthConstraints()
}

// determineColumnCount returns the number of columns in the table.
func (t *Table) determineColumnCount() int {
	if len(t.headers) > 0 {
		return len(t.headers)
	}
	if len(t.rows) > 0 {
		return len(t.rows[0])
	}
	return 0
}

// useFixedWidths checks if fixed column widths are configured and applies them.
func (t *Table) useFixedWidths() bool {
	if len(t.options.ColumnWidths) == 0 {
		return false
	}
	t.columnWidths = make([]int, len(t.options.ColumnWidths))
	copy(t.columnWidths, t.options.ColumnWidths)
	return true
}

// initializeWidths creates the column width slice.
func (t *Table) initializeWidths(numCols int) {
	t.columnWidths = make([]int, numCols)
}

// updateWidthsFromHeaders updates column widths based on header text.
func (t *Table) updateWidthsFromHeaders() {
	for i, header := range t.headers {
		if i < len(t.columnWidths) {
			t.columnWidths[i] = visualWidth(header)
		}
	}
}

// updateWidthsFromRows updates column widths based on row data.
func (t *Table) updateWidthsFromRows(numCols int) {
	for _, row := range t.rows {
		if t.isSeparatorRow(row) {
			continue
		}
		t.updateWidthsFromCells(row, numCols)
	}
}

// isSeparatorRow checks if a row is a separator.
func (t *Table) isSeparatorRow(row []string) bool {
	return len(row) == 1 && row[0] == "__SEPARATOR__"
}

// updateWidthsFromCells updates widths from a single row of cells.
func (t *Table) updateWidthsFromCells(cells []string, numCols int) {
	for i, cell := range cells {
		if i >= numCols {
			break
		}
		cellLen := visualWidth(cell)
		if cellLen > t.columnWidths[i] {
			t.columnWidths[i] = cellLen
		}
	}
}

// updateWidthsFromFooter updates column widths based on footer text.
func (t *Table) updateWidthsFromFooter(numCols int) {
	t.updateWidthsFromCells(t.footer, numCols)
}

// applyWidthConstraints applies min/max width constraints to columns.
func (t *Table) applyWidthConstraints() {
	for i := range t.columnWidths {
		t.columnWidths[i] = t.constrainWidth(t.columnWidths[i])
	}
}

// constrainWidth applies min/max constraints to a single width value.
func (t *Table) constrainWidth(width int) int {
	if t.options.MinWidth > 0 && width < t.options.MinWidth {
		return t.options.MinWidth
	}
	if t.options.MaxWidth > 0 && width > t.options.MaxWidth {
		return t.options.MaxWidth
	}
	return width
}

// stripANSI removes ANSI color codes from a string to get actual display length.
func stripANSI(str string) string {
	result := ""
	inEscape := false
	escapeSeq := ""

	for i := 0; i < len(str); i++ {
		char := str[i]

		if char == 0x1b {
			inEscape = true
			escapeSeq = ""
			continue
		}

		if inEscape {
			escapeSeq += string(char)
			if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
				inEscape = false
				escapeSeq = ""
			}
			continue
		}

		result += string(char)
	}

	return result
}

// REFACTORED: visualWidth - reduced complexity from 22 to <15
// Split wide character detection into separate functions

// visualWidth calculates the visual width of a string, accounting for wide characters.
//
// Many Unicode characters (emojis, symbols) take up 2 terminal columns instead of 1.
// However, this is terminal-dependent. Some characters like ✓, ✗, ⚡ may be rendered
// as 1 column in some terminals and 2 columns in others.
//
// This function uses a conservative approach: only count characters as wide (2 columns)
// if they are definitely wide in most terminals (CJK characters, emoji blocks).
// Common symbols like checkmarks and arrows are treated as single-width (1 column).
func visualWidth(str string) int {
	stripped := stripANSI(str)
	width := 0

	for _, r := range stripped {
		width += runeWidth(r)
	}

	return width
}

// runeWidth returns the display width of a single rune (1 or 2 columns).
func runeWidth(r rune) int {
	if isWideCharacter(r) {
		return 2
	}
	return 1
}

// isWideCharacter checks if a rune takes up 2 terminal columns.
func isWideCharacter(r rune) bool {
	return isCJKCharacter(r) || isFullwidthCharacter(r) || isEmojiCharacter(r)
}

// isCJKCharacter checks if a rune is a CJK (Chinese, Japanese, Korean) character.
func isCJKCharacter(r rune) bool {
	return isHangulJamo(r) ||
		isCJKRadical(r) ||
		isHiraganaKatakana(r) ||
		isHangulSyllable(r) ||
		isCJKCompatibility(r) ||
		isCJKExtension(r)
}

// isHangulJamo checks for Hangul Jamo range.
func isHangulJamo(r rune) bool {
	return r >= 0x1100 && r <= 0x115F
}

// isCJKRadical checks for CJK Radicals Supplement range.
func isCJKRadical(r rune) bool {
	return r >= 0x2E80 && r <= 0x303E
}

// isHiraganaKatakana checks for Hiragana, Katakana, CJK Unified Ideographs range.
func isHiraganaKatakana(r rune) bool {
	return r >= 0x3040 && r <= 0xA4CF
}

// isHangulSyllable checks for Hangul Syllables range.
func isHangulSyllable(r rune) bool {
	return r >= 0xAC00 && r <= 0xD7A3
}

// isCJKCompatibility checks for CJK Compatibility ranges.
func isCJKCompatibility(r rune) bool {
	return (r >= 0xF900 && r <= 0xFAFF) ||
		(r >= 0xFE30 && r <= 0xFE6F)
}

// isCJKExtension checks for CJK Ideograph Extensions range.
func isCJKExtension(r rune) bool {
	return r >= 0x20000 && r <= 0x2FFFD
}

// isFullwidthCharacter checks for Fullwidth Forms range.
func isFullwidthCharacter(r rune) bool {
	return (r >= 0xFF00 && r <= 0xFF60) ||
		(r >= 0xFFE0 && r <= 0xFFE6)
}

// isEmojiCharacter checks for Emoji and pictographs range.
func isEmojiCharacter(r rune) bool {
	return r >= 0x1F300 && r <= 0x1F9FF
}

// alignText aligns text within a given width according to the specified alignment.
func (t *Table) alignText(text string, width int, alignment Alignment) string {
	actualLen := visualWidth(text)

	if actualLen >= width {
		return text
	}

	padding := width - actualLen

	switch alignment {
	case AlignRight:
		return strings.Repeat(" ", padding) + text
	case AlignCenter:
		leftPad := padding / 2
		rightPad := padding - leftPad
		return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
	default:
		return text + strings.Repeat(" ", padding)
	}
}

// getAlignment returns the alignment for a specific column.
func (t *Table) getAlignment(col int) Alignment {
	if col < len(t.options.Alignment) {
		return t.options.Alignment[col]
	}
	return AlignLeft
}

// tableIsColorEnabled checks if colors should be enabled for tables.
// This is a table-specific implementation that can be overridden by
// a package-level isColorEnabled() function if it exists.
func tableIsColorEnabled() bool {
	// Check if output is a terminal
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		return false
	}

	// Check NO_COLOR environment variable
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	// Check CI environment
	if os.Getenv("CI") != "" {
		return false
	}

	return true
}

// tableColorize applies color to text if colors are enabled.
// This is a table-specific implementation that can use the package-level
// colorize() function if it exists, or fall back to this implementation.
func tableColorize(text, color string) string {
	if !tableIsColorEnabled() || color == "" {
		return text
	}
	return color + text + "\033[0m"
}

// colorize applies a color code to text if colors are enabled.
func (t *Table) colorize(text, color string) string {
	if color == "" {
		return text
	}
	return tableColorize(text, color)
}

// buildBorder creates a border line (top, middle, or bottom).
func (t *Table) buildBorder(left, cross, right, horizontal string) string {
	if !t.options.ShowBorders || left == "" {
		return ""
	}

	parts := []string{left}
	for i, width := range t.columnWidths {
		parts = append(parts, strings.Repeat(horizontal, width+2*t.options.Padding))
		if i < len(t.columnWidths)-1 {
			parts = append(parts, cross)
		}
	}
	parts = append(parts, right)

	line := strings.Join(parts, "")
	return t.colorize(line, t.options.BorderColor)
}

// buildRow creates a formatted row of data.
func (t *Table) buildRow(cells []string, color string) string {
	if len(cells) == 0 {
		return ""
	}

	if len(cells) == 1 && cells[0] == "__SEPARATOR__" {
		return t.buildBorder(
			t.style.LeftCross,
			t.style.Cross,
			t.style.RightCross,
			t.style.Horizontal,
		)
	}

	verticalBorder := t.getVerticalBorder()
	parts := t.buildRowParts(cells, color, verticalBorder)

	return strings.Join(parts, "")
}

// getVerticalBorder returns the vertical border string with color applied.
func (t *Table) getVerticalBorder() string {
	if !t.options.ShowBorders {
		return ""
	}
	return t.colorize(t.style.Vertical, t.options.BorderColor)
}

// buildRowParts constructs the parts of a row (borders and cells).
func (t *Table) buildRowParts(cells []string, color, verticalBorder string) []string {
	parts := []string{}

	if t.options.ShowBorders {
		parts = append(parts, verticalBorder)
	}

	parts = t.appendCellParts(parts, cells, color, verticalBorder)

	if t.options.ShowBorders {
		parts = append(parts, verticalBorder)
	}

	return parts
}

// appendCellParts adds formatted cells to the parts slice.
func (t *Table) appendCellParts(parts []string, cells []string, color, verticalBorder string) []string {
	for i := 0; i < len(t.columnWidths); i++ {
		cell := t.getCellContent(cells, i)
		displayCell := t.colorizeCell(cell, color)
		formattedCell := t.formatCell(displayCell, i)

		parts = append(parts, formattedCell)

		if t.shouldAddVerticalBorder(i) {
			parts = append(parts, verticalBorder)
		}
	}
	return parts
}

// getCellContent retrieves cell content or returns empty string.
func (t *Table) getCellContent(cells []string, index int) string {
	if index < len(cells) {
		return cells[index]
	}
	return ""
}

// colorizeCell applies color to a cell if needed.
func (t *Table) colorizeCell(cell, color string) string {
	if color == "" || strings.Contains(cell, "\x1b[") {
		return cell
	}
	return t.colorize(cell, color)
}

// formatCell formats a cell with padding and alignment.
func (t *Table) formatCell(displayCell string, columnIndex int) string {
	padding := strings.Repeat(" ", t.options.Padding)
	aligned := t.alignText(displayCell, t.columnWidths[columnIndex], t.getAlignment(columnIndex))
	return padding + aligned + padding
}

// shouldAddVerticalBorder checks if a vertical border should be added after this column.
func (t *Table) shouldAddVerticalBorder(columnIndex int) bool {
	return t.options.ShowBorders && columnIndex < len(t.columnWidths)-1
}

// REFACTORED: String - reduced complexity from 16 to <15
// Split into smaller functions for each table section

// String returns the table as a formatted string.
func (t *Table) String() string {
	t.calculateColumnWidths()

	var output strings.Builder

	t.writeTopBorder(&output)
	t.writeHeader(&output)
	t.writeRows(&output)
	t.writeFooter(&output)
	t.writeBottomBorder(&output)

	return output.String()
}

// writeTopBorder writes the top border of the table.
func (t *Table) writeTopBorder(output *strings.Builder) {
	if !t.options.ShowBorders || t.style.TopLeft == "" {
		return
	}
	output.WriteString(t.buildBorder(
		t.style.TopLeft,
		t.style.TopCross,
		t.style.TopRight,
		t.style.Horizontal,
	))
	output.WriteString("\n")
}

// writeHeader writes the header row and its separator.
func (t *Table) writeHeader(output *strings.Builder) {
	if !t.options.ShowHeader || len(t.headers) == 0 {
		return
	}

	headerRow := t.buildRow(t.headers, t.options.HeaderColor)
	output.WriteString(headerRow)
	output.WriteString("\n")

	t.writeHeaderSeparator(output)
}

// writeHeaderSeparator writes the separator line after the header.
func (t *Table) writeHeaderSeparator(output *strings.Builder) {
	if !t.options.ShowBorders || t.style.HeaderSeparator == "" {
		return
	}
	output.WriteString(t.buildBorder(
		t.style.LeftCross,
		t.style.Cross,
		t.style.RightCross,
		t.style.HeaderSeparator,
	))
	output.WriteString("\n")
}

// writeRows writes all data rows.
func (t *Table) writeRows(output *strings.Builder) {
	for i, row := range t.rows {
		t.writeRow(output, row, i)
	}
}

// writeRow writes a single data row.
func (t *Table) writeRow(output *strings.Builder, row []string, rowIndex int) {
	if t.isSeparatorRow(row) {
		output.WriteString(t.buildRow(row, ""))
		output.WriteString("\n")
		return
	}

	rowColor := t.getRowColor(rowIndex)
	output.WriteString(t.buildRow(row, rowColor))
	output.WriteString("\n")
}

// getRowColor returns the color for a row based on alternating colors.
func (t *Table) getRowColor(rowIndex int) string {
	if len(t.options.RowColors) == 0 {
		return ""
	}
	return t.options.RowColors[rowIndex%len(t.options.RowColors)]
}

// writeFooter writes the footer row with its separator.
func (t *Table) writeFooter(output *strings.Builder) {
	if len(t.footer) == 0 {
		return
	}

	t.writeFooterSeparator(output)

	output.WriteString(t.buildRow(t.footer, t.options.FooterColor))
	output.WriteString("\n")
}

// writeFooterSeparator writes the separator line before the footer.
func (t *Table) writeFooterSeparator(output *strings.Builder) {
	if !t.options.ShowBorders || t.style.HeaderSeparator == "" {
		return
	}
	output.WriteString(t.buildBorder(
		t.style.LeftCross,
		t.style.Cross,
		t.style.RightCross,
		t.style.HeaderSeparator,
	))
	output.WriteString("\n")
}

// writeBottomBorder writes the bottom border of the table.
func (t *Table) writeBottomBorder(output *strings.Builder) {
	if !t.options.ShowBorders || t.style.BottomLeft == "" {
		return
	}
	output.WriteString(t.buildBorder(
		t.style.BottomLeft,
		t.style.BottomCross,
		t.style.BottomRight,
		t.style.Horizontal,
	))
	output.WriteString("\n")
}

// Print outputs the table to stdout.
func (t *Table) Print() {
	fmt.Print(t.String())
}

// Clear removes all rows but keeps headers and configuration.
func (t *Table) Clear() *Table {
	t.rows = [][]string{}
	t.columnWidths = []int{}
	return t
}

// RowCount returns the number of data rows (excluding header and footer).
func (t *Table) RowCount() int {
	return len(t.rows)
}

// ColumnCount returns the number of columns based on headers or first row.
func (t *Table) ColumnCount() int {
	if len(t.headers) > 0 {
		return len(t.headers)
	}
	if len(t.rows) > 0 {
		return len(t.rows[0])
	}
	return 0
}
