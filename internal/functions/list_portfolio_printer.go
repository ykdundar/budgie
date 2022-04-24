package functions

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ykdundar/budgie/internal/objects"
)

func ListPortfolioPrinter(portfolio []objects.Portfolio, head string) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{head, head}, rowConfigAutoMerge)
	t.AppendHeader(table.Row{"ID", "NAME"})

	for _, v := range portfolio {
		t.AppendRow(table.Row{v.ID, v.Name})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Render()
}
