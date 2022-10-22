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
