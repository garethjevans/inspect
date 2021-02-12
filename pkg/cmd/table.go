package cmd

import "github.com/jedib0t/go-pretty/v6/table"

var (
	tableStyle      = table.StyleDefault
	writeSeparators = true
	headers         = true
	enableMarkdown  = false
)

// EnableMarkdown tables are outputted in markdown format, rather than ASCII.
func EnableMarkdown() {
	enableMarkdown = true
}

// DisableHeaders disables the headers on tables.
func DisableHeaders() {
	headers = false
}

// Raw enables raw output on all tables.
func Raw() {
	writeSeparators = false
	tableStyle = table.Style{
		Name:    "Raw",
		Box:     table.StyleBoxDefault,
		Color:   table.ColorOptionsDefault,
		Format:  table.FormatOptionsDefault,
		HTML:    table.DefaultHTMLOptions,
		Options: table.OptionsNoBordersAndSeparators,
		Title:   table.TitleOptionsDefault,
	}
}

// Reset resets all table formatting to the default.
func Reset() {
	tableStyle = table.StyleDefault
	writeSeparators = true
	headers = true
	enableMarkdown = false
}
