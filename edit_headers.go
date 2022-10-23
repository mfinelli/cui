// cui: http request/response tui
// Copyright 2022 Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
