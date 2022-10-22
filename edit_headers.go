package main

import (
	"github.com/rivo/tview"
)

func setEditHeadersPlain(cui *cuiApp, req *cuiRequest) {
	cui.RequestHeaders.Clear()
	cui.Request.Clear().SetDirection(tview.FlexRow).
		AddItem(cui.RequestHeaders, 0, 1, true)

	i := 0
	for header, value := range req.Headers {
		cui.RequestHeaders.SetCell(i, 0, tview.NewTableCell(header))
		cui.RequestHeaders.SetCell(i, 1, tview.NewTableCell(value))
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

func deleteHeader(cui *cuiApp, req *cuiRequest) {
}
