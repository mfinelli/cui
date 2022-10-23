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

func addParameter(cui *cuiApp, req *cuiRequest) {
	key := cui.RequestParameterKey.GetText()
	value := cui.RequestParameterValue.GetText()

	req.Parameters[key] = value
	cui.RequestParameterKey.SetText("")
	cui.RequestParameterValue.SetText("")
}

func setEditParametersPlain(cui *cuiApp, req *cuiRequest) {
	cui.RequestParameters.Clear().SetSelectable(true, false)
	cui.Request.Clear().SetDirection(tview.FlexRow).
		AddItem(cui.RequestParameters, 0, 1, true)

	i := 0
	for key, value := range req.Parameters {
		cui.RequestParameters.SetCell(i, 0, tview.NewTableCell(key))
		cui.RequestParameters.SetCell(i, 1, tview.NewTableCell(value))
		i += 1
	}
}

func setEditParametersAdd(cui *cuiApp, req *cuiRequest) {
	input := tview.NewFlex().
		AddItem(cui.RequestParameterKey, 0, 1, true).
		AddItem(cui.RequestParameterValue, 0, 1, false)

	cui.Request.Clear().SetDirection(tview.FlexRow).
		AddItem(input, 1, 0, true).
		AddItem(cui.RequestParameters, 0, 1, false)
}

func deleteParameter(app *tview.Application, cui *cuiApp, req *cuiRequest) {
	// TODO: keeping app around for now in case we want to implement modal confirm
	if cui.RequestParameters.GetRowCount() >= 1 {
		row, _ := cui.RequestParameters.GetSelection()
		key := cui.RequestParameters.GetCell(row, 0).Text
		delete(req.Parameters, key)
	}
}
