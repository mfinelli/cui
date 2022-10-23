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

func initRequest(app *tview.Application, cui *cuiApp, req *cuiRequest, method, url, body string, params, headers map[string]string) {
	methodIndex := getStrSliceIndex(&httpMethods, method)
	requestKind := getStrSliceIndex(&requestKinds, "Raw")

	setInstructions(cui, "")
	app.SetFocus(cui.Main)

	cui.ViewHasResponse = false
	cui.ResponseStatus.SetText("")
	cui.ResponseBody.Clear()
	cui.ResponseHeaders.Clear()

	cui.ViewResponse = "body"
	cui.Response.Clear().SetDirection(tview.FlexRow).
		AddItem(cui.ResponseStatus, 1, 0, false).
		AddItem(cui.ResponseBody, 0, 1, true)

	cui.ViewRequest = "RequestBody"
	cui.ViewRequestInputType = "Textarea"
	cui.MethodDropdown.SetCurrentOption(methodIndex)
	cui.RequestKindDropdown.SetCurrentOption(requestKind)

	cui.UrlInput.SetText(url)
	cui.RequestBody.SetText(body, true)

	if url == "" {
		url = "http://example.com"
	}

	req.Method = method
	req.URL = url
	req.Headers = headers
	req.Parameters = params
	req.Body = body
}
