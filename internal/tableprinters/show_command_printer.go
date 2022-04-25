package tableprinters

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/ykdundar/budgie/api"
)

func ShowCmdPrinter(intraday api.Intraday, head string) {
	t := table.NewWriter()
	t.SetColumnConfigs(tableConfig)
	t.SetStyle(table.StyleRounded)
	t.SetTitle(head)
	t.Style().Title.Align = text.AlignCenter
	t.Style().Title.Format = text.FormatUpper
	t.SetOutputMirror(os.Stdout)
	t.Style().Options.SeparateRows = true
	t.AppendHeader(table.Row{"SYMBOL", "OPEN", "HIGH", "LOW", "LAST", "CLOSE", "DATE", "EXCHANGE"})

	for _, v := range intraday.Data {
		formattedDate, _ := time.Parse("2006-01-02T15:04:05-0700", v.Date)
		t.AppendRow(table.Row{v.Symbol, v.Open, v.High, v.Low, v.Last,
			v.Close, formattedDate.Format("2006-01-02"), v.Exchange})
		t.AppendSeparator()
	}

	t.Render()
}
