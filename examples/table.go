//go:build ignore
// +build ignore

// examples/table.go demonstrates ColorBear's table functionality.
//
// This example shows various ways to use tables:
//   - Simple tables with headers and rows
//   - Different visual styles (rounded, double, minimal, etc.)
//   - Custom colors and alignments
//   - Styled rows with semantic functions
//   - Footers for totals and summaries
//   - Real-world use cases (status dashboards, reports, etc.)
//
// Run with: go run examples/table.go
package main

import (
	"github.com/elpol4k0/colorbear"
)

func main() {
	println("ColorBear Table Examples")
	println("========================\n")

	println("Example 1: Simple Table")
	simpleTable()

	println("\nExample 2: Table Styles")
	tableStyles()

	println("\nExample 3: Custom Colors & Alignment")
	customTable()

	println("\nExample 4: Status Table with Styled Rows")
	statusTable()

	println("\nExample 5: Table with Footer (Sales Report)")
	salesReport()

	println("\nExample 6: CI/CD Pipeline Status")
	cicdStatus()

	println("\nExample 7: System Resource Monitoring")
	systemResources()

	println("\nExample 8: Git Commit Log")
	gitLog()

	println("\nExample 9: Minimal Style (Spreadsheet-like)")
	minimalTable()

	println("\nExample 10: Table with Separators")
	tableWithSeparators()
}

func simpleTable() {
	table := colorbear.NewTable()
	table.SetHeaders("Name", "Age", "City")
	table.AddRow("Alice", "28", "Berlin")
	table.AddRow("Bob", "34", "Munich")
	table.AddRow("Charlie", "42", "Hamburg")
	table.Print()
}

func tableStyles() {
	styles := []*colorbear.TableStyle{
		colorbear.TableStyleSimple,
		colorbear.TableStyleRounded,
		colorbear.TableStyleDouble,
		colorbear.TableStyleBold,
		colorbear.TableStyleMinimal,
		colorbear.TableStyleCompact,
	}

	for _, style := range styles {
		println("\n" + style.Name + " style:")
		table := colorbear.NewTable(
			colorbear.WithTableStyle(style),
		)
		table.SetHeaders("Name", "Status")
		table.AddRow("Build", "Pass")
		table.AddRow("Test", "Pass")
		table.Print()
	}
}

func customTable() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleDouble),
		colorbear.WithHeaderColor(colorbear.GreenCode),
		colorbear.WithBorderColor(colorbear.CyanCode),
		colorbear.WithAlignment(
			colorbear.AlignLeft,
			colorbear.AlignRight,
			colorbear.AlignCenter,
		),
		colorbear.WithPadding(2),
	)

	table.SetHeaders("Product", "Price", "Stock")
	table.AddRow("Laptop", "$1,299", "In Stock")
	table.AddRow("Mouse", "$29", "Low Stock")
	table.AddRow("Keyboard", "$89", "In Stock")
	table.AddRow("Monitor", "$449", "Out of Stock")
	table.Print()
}

// statusTable demonstrates using styled rows with table-safe semantic functions.
//
// IMPORTANT: Use Table* variants (TableSuccess, TableError, etc.) in tables
// for consistent alignment. These use ASCII text: [OK], [ERR], [!], [i]
// instead of Unicode symbols which may have inconsistent width.
func statusTable() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleRounded),
	)

	table.SetHeaders("Task", "Status", "Progress", "Time")

	table.AddStyledRow(
		"Deploy Backend",
		colorbear.TableSuccess("Complete"),
		"100%",
		"2m 34s",
	)

	table.AddStyledRow(
		"Run Tests",
		colorbear.TableSuccess("Passed"),
		"100%",
		"5m 12s",
	)

	table.AddStyledRow(
		"Build Frontend",
		colorbear.TableWarning("Running"),
		"67%",
		"1m 08s",
	)

	table.AddStyledRow(
		"Update Docs",
		colorbear.TableError("Failed"),
		"0%",
		"0s",
	)

	table.AddStyledRow(
		"Deploy Frontend",
		colorbear.TableInfo("Waiting"),
		"0%",
		"-",
	)

	table.Print()
}

func salesReport() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleBold),
		colorbear.WithHeaderColor(colorbear.CyanCode),
		colorbear.WithFooterColor(colorbear.YellowCode),
		colorbear.WithAlignment(
			colorbear.AlignLeft,
			colorbear.AlignRight,
			colorbear.AlignRight,
		),
	)

	table.SetHeaders("Product", "Units Sold", "Revenue")
	table.AddRow("Laptop", "45", "$58,455")
	table.AddRow("Mouse", "234", "$6,786")
	table.AddRow("Keyboard", "156", "$13,884")
	table.AddRow("Monitor", "89", "$39,961")

	table.SetFooter("Total", "524", "$119,086")

	table.Print()
}

// cicdStatus demonstrates a real-world CI/CD pipeline status table.
//
// Uses table-safe semantic functions with empty strings to show just
// the status indicator: [OK], [!], etc.
func cicdStatus() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleRounded),
		colorbear.WithAlignment(
			colorbear.AlignLeft,
			colorbear.AlignCenter,
			colorbear.AlignRight,
			colorbear.AlignLeft,
		),
	)

	table.SetHeaders("Stage", "Status", "Duration", "Details")

	table.AddStyledRow(
		"Checkout",
		colorbear.TableSuccess(""),
		"12s",
		"main@abc123",
	)

	table.AddStyledRow(
		"Build",
		colorbear.TableSuccess(""),
		"2m 34s",
		"No warnings",
	)

	table.AddStyledRow(
		"Test",
		colorbear.TableSuccess(""),
		"5m 12s",
		"127/127 passed",
	)

	table.AddStyledRow(
		"Lint",
		colorbear.TableSuccess(""),
		"45s",
		"No issues",
	)

	table.AddStyledRow(
		"Deploy",
		colorbear.TableWarning(""),
		"1m 08s",
		"production",
	)

	table.Print()
}

func systemResources() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleDouble),
		colorbear.WithHeaderColor(colorbear.MagentaCode),
		colorbear.WithAlignment(
			colorbear.AlignLeft,
			colorbear.AlignRight,
			colorbear.AlignRight,
			colorbear.AlignRight,
		),
	)

	table.SetHeaders("Resource", "Used", "Total", "Percentage")
	table.AddRow("CPU", "2.4 GHz", "3.2 GHz", "75%")
	table.AddRow("Memory", "12.8 GB", "16.0 GB", "80%")
	table.AddRow("Disk", "450 GB", "1 TB", "45%")
	table.AddRow("Swap", "0 GB", "8 GB", "0%")

	table.Print()
}

func gitLog() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleMinimal),
		colorbear.WithHeaderColor(colorbear.YellowCode),
	)

	table.SetHeaders("Hash", "Author", "Date", "Message")
	table.AddRow("abc1234", "Alice", "2 hours ago", "Fix login bug")
	table.AddRow("def5678", "Bob", "5 hours ago", "Add user dashboard")
	table.AddRow("ghi9012", "Charlie", "1 day ago", "Update dependencies")
	table.AddRow("jkl3456", "Alice", "2 days ago", "Improve performance")
	table.AddRow("mno7890", "Bob", "3 days ago", "Add tests")

	table.Print()
}

func minimalTable() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleMinimal),
	)

	table.SetHeaders("ID", "Name", "Email", "Role", "Status")
	table.AddRow("001", "Alice Johnson", "alice@example.com", "Admin", "Active")
	table.AddRow("002", "Bob Smith", "bob@example.com", "User", "Active")
	table.AddRow("003", "Charlie Brown", "charlie@example.com", "User", "Inactive")
	table.AddRow("004", "Diana Prince", "diana@example.com", "Manager", "Active")

	table.Print()
}

func tableWithSeparators() {
	table := colorbear.NewTable(
		colorbear.WithTableStyle(colorbear.TableStyleRounded),
		colorbear.WithHeaderColor(colorbear.CyanCode),
	)

	table.SetHeaders("Category", "Item", "Price")

	table.AddRow("Fruits", "Apple", "$2.99")
	table.AddRow("Fruits", "Banana", "$1.49")
	table.AddRow("Fruits", "Orange", "$3.49")

	table.AddSeparator()

	table.AddRow("Vegetables", "Carrot", "$1.99")
	table.AddRow("Vegetables", "Broccoli", "$2.49")

	table.AddSeparator()

	table.AddRow("Dairy", "Milk", "$3.99")
	table.AddRow("Dairy", "Cheese", "$5.49")

	table.Print()
}
