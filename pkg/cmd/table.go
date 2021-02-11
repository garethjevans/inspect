package cmd

import "github.com/jedib0t/go-pretty/v6/table"

var (
	tableStyle      = table.StyleDefault
	writeSeparators = true
	headers         = true
	enableMarkdown  = false
)

func EnableMarkdown() {
	enableMarkdown = true
}

func DisableHeaders() {
	headers = false
}

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

func Reset() {
	tableStyle = table.StyleDefault
	writeSeparators = true
	headers = true
	enableMarkdown = false
}
