package main

import (
	"github.com/rivo/tview"
)

func addHeader(cui *cuiApp, req *cuiRequest) {
	key := cui.RequestHeaderKey.GetText()
	value := cui.RequestHeaderValue.GetText()

	req.Headers[key] = value
	cui.RequestHeaderKey.SetText("")
	cui.RequestHeaderValue.SetText("")
}

func setEditHeadersPlain(cui *cuiApp, req *cuiRequest) {
	cui.RequestHeaders.Clear().SetSelectable(true, false)
	cui.Request.Clear().SetDirection(tview.FlexRow).
		AddItem(cui.RequestHeaders, 0, 1, true)

	i := 0
	for header, value := range req.Headers {
		cui.RequestHeaders.SetCell(i, 0, tview.NewTableCell(header))
		cui.RequestHeaders.SetCell(i, 1, tview.NewTableCell(value))
		i += 1
	}
}

func setEditHeadersAdd(cui *cuiApp, req *cuiRequest) {
	input := tview.NewFlex().
		AddItem(cui.RequestHeaderKey, 0, 1, true).
		AddItem(cui.RequestHeaderValue, 0, 1, false)

	cui.Request.Clear().SetDirection(tview.FlexRow).
		AddItem(input, 1, 0, true).
		AddItem(cui.RequestHeaders, 0, 1, false)
}

func deleteHeader(app *tview.Application, cui *cuiApp, req *cuiRequest) {
	// TODO: keeping app around for now in case we want to implement modal confirm
	if cui.RequestHeaders.GetRowCount() >= 1 {
		row, _ := cui.RequestHeaders.GetSelection()
		header := cui.RequestHeaders.GetCell(row, 0).Text
		delete(req.Headers, header)
	}
}
