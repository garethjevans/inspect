package cmd

import (
	"sort"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/garethjevans/inspect/pkg/util"
	"github.com/jedib0t/go-pretty/v6/table"
)

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

func writeTableForImage(labels map[string]string, t table.Writer, repo string, tag string) {
	t.SetStyle(tableStyle)

	if headers {
		t.AppendHeader(table.Row{"Label", "Value"})
	}

	t.AppendRow(table.Row{repo, tag})

	if writeSeparators {
		t.AppendSeparator()
	}

	keys := util.AllKeys(labels)
	sort.Strings(keys)

	for _, k := range keys {
		t.AppendRow(table.Row{k, labels[k]})
	}

	if writeSeparators {
		t.AppendSeparator()
	}

	url := inspect.GitHubURL(labels)
	if url != "" {
		t.AppendRow(table.Row{"GitHub URL", inspect.GitHubURL(labels)})
	}
}

func writeDiffTableForImages(labels1 map[string]string, labels2 map[string]string, t table.Writer, repo string, tag1 string, tag2 string) {
	t.SetStyle(tableStyle)

	if headers {
		t.AppendHeader(table.Row{"Image", "1", "2"})
	}

	if writeSeparators {
		t.AppendSeparator()
	}

	t.AppendRow(table.Row{repo, tag1, tag2})

	if writeSeparators {
		t.AppendSeparator()
	}

	keys := util.AllKeys(labels1, labels2)
	sort.Strings(keys)

	for _, k := range keys {
		if labels1[k] == labels2[k] {
			t.AppendRow(table.Row{k, labels1[k], labels2[k]})
		} else {
			t.AppendRow(table.Row{k, colorError(labels1[k]), colorError(labels2[k])})
		}
	}

	if writeSeparators {
		t.AppendSeparator()
	}

	ghURL1, ghURL2 := inspect.GitHubURL(labels1), inspect.GitHubURL(labels2)

	if ghURL1 != "" && ghURL2 != "" {
		t.AppendRow(table.Row{
			"GitHub URL",
			inspect.GitHubURL(labels1),
			inspect.GitHubURL(labels2),
		})
	}
}
