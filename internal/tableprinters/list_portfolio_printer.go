package tableprinters

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/ykdundar/budgie/internal/objects"
)

func ListPortfolioPrinter(portfolio []objects.Portfolio, head string) {
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:      1,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
		{
			Number:      2,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
		},
	})
	t.SetStyle(table.StyleRounded)
	t.SetTitle(head)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Format = text.FormatUpper
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "NAME"})
	t.Style().Options.SeparateRows = true

	for _, v := range portfolio {
		t.AppendRow(table.Row{v.ID, v.Name})
		t.AppendSeparator()
	}

	t.Render()
}
