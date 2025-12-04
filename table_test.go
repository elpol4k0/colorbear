package colorbear

import (
	"strings"
	"testing"
)

func TestTableBasic(t *testing.T) {
	table := NewTable()
	table.SetHeaders("Name", "Age")
	table.AddRow("Alice", "28")
	table.AddRow("Bob", "34")

	output := table.String()
	if output == "" {
		t.Error("Table output should not be empty")
	}

	// Check that headers are present
	if !strings.Contains(output, "Name") || !strings.Contains(output, "Age") {
		t.Error("Table should contain headers")
	}

	// Check that data is present
	if !strings.Contains(output, "Alice") || !strings.Contains(output, "Bob") {
		t.Error("Table should contain row data")
	}
}

func TestTableStyles(t *testing.T) {
	styles := []*TableStyle{
		TableStyleSimple,
		TableStyleRounded,
		TableStyleDouble,
		TableStyleBold,
		TableStyleMinimal,
		TableStyleCompact,
	}

	for _, style := range styles {
		table := NewTable(WithTableStyle(style))
		table.SetHeaders("Col1", "Col2")
		table.AddRow("Data1", "Data2")

		output := table.String()
		if output == "" {
			t.Errorf("Table with style %s should produce output", style.Name)
		}
	}
}

func TestTableAlignment(t *testing.T) {
	table := NewTable(
		WithAlignment(AlignLeft, AlignRight, AlignCenter),
	)
	table.SetHeaders("Left", "Right", "Center")
	table.AddRow("A", "B", "C")

	output := table.String()
	if output == "" {
		t.Error("Table with alignment should produce output")
	}
}

func TestTableFooter(t *testing.T) {
	table := NewTable()
	table.SetHeaders("Item", "Price")
	table.AddRow("A", "10")
	table.AddRow("B", "20")
	table.SetFooter("Total", "30")

	output := table.String()
	if !strings.Contains(output, "Total") || !strings.Contains(output, "30") {
		t.Error("Table should contain footer")
	}
}

func TestTableChaining(t *testing.T) {
	table := NewTable()
	output := table.SetHeaders("A", "B").
		AddRow("1", "2").
		AddRow("3", "4").
		String()

	if output == "" {
		t.Error("Chained table operations should produce output")
	}
}

func TestTableClear(t *testing.T) {
	table := NewTable()
	table.SetHeaders("A", "B")
	table.AddRow("1", "2")
	table.AddRow("3", "4")

	if table.RowCount() != 2 {
		t.Errorf("Expected 2 rows, got %d", table.RowCount())
	}

	table.Clear()

	if table.RowCount() != 0 {
		t.Errorf("After clear, expected 0 rows, got %d", table.RowCount())
	}

	// Headers should still be present
	if table.ColumnCount() != 2 {
		t.Errorf("After clear, headers should remain. Expected 2 columns, got %d", table.ColumnCount())
	}
}

func TestTableColumnCount(t *testing.T) {
	table := NewTable()

	// No headers or rows
	if table.ColumnCount() != 0 {
		t.Errorf("Empty table should have 0 columns, got %d", table.ColumnCount())
	}

	// With headers
	table.SetHeaders("A", "B", "C")
	if table.ColumnCount() != 3 {
		t.Errorf("Expected 3 columns, got %d", table.ColumnCount())
	}
}

func TestTableSeparator(t *testing.T) {
	table := NewTable()
	table.SetHeaders("A", "B")
	table.AddRow("1", "2")
	table.AddSeparator()
	table.AddRow("3", "4")

	output := table.String()
	if output == "" {
		t.Error("Table with separator should produce output")
	}

	// Should have more lines due to separator
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) < 5 {
		t.Error("Table with separator should have additional line")
	}
}

func TestStripANSI(t *testing.T) {
	// Test with ANSI color codes
	colored := "\x1b[32mGreen Text\x1b[0m"
	stripped := stripANSI(colored)

	if stripped != "Green Text" {
		t.Errorf("Expected 'Green Text', got '%s'", stripped)
	}

	// Test without ANSI codes
	plain := "Plain Text"
	strippedPlain := stripANSI(plain)

	if strippedPlain != plain {
		t.Errorf("Plain text should remain unchanged")
	}
}

func TestTableWithStyledRows(t *testing.T) {
	table := NewTable()
	table.SetHeaders("Task", "Status")

	// Simulate adding styled rows (would use Success/Warning/Error in real code)
	table.AddStyledRow("Task 1", "Complete")
	table.AddStyledRow("Task 2", "In Progress")

	output := table.String()
	if !strings.Contains(output, "Task 1") || !strings.Contains(output, "Task 2") {
		t.Error("Styled rows should appear in output")
	}
}

func TestTableOptions(t *testing.T) {
	// Test that options are applied correctly
	table := NewTable(
		WithHeaderColor(CyanCode),
		WithBorderColor(GreenCode),
		WithFooterColor(YellowCode),
		WithPadding(2),
		WithAutoSize(true),
	)

	if table.options.HeaderColor != CyanCode {
		t.Error("HeaderColor option not applied")
	}

	if table.options.BorderColor != GreenCode {
		t.Error("BorderColor option not applied")
	}

	if table.options.FooterColor != YellowCode {
		t.Error("FooterColor option not applied")
	}

	if table.options.Padding != 2 {
		t.Error("Padding option not applied")
	}

	if !table.options.AutoSize {
		t.Error("AutoSize option not applied")
	}
}
