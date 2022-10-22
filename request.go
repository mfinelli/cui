// cui: http request/response tui
// Copyright 2022  Mario Finelli
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
	"fmt"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"strings"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/rivo/tview"
)

type cuiRequest struct {
	Method string
	URL    string

	Headers    map[string]string
	Body       string
	Parameters map[string]string
}

func sendRequest(req cuiRequest, cui *cuiApp, hasResponse *bool) error {
	client := &http.Client{}
	cui.ResponseBody.Clear()
	cui.ResponseHeaders.Clear()

	highlight := false
	highlightLang := ""

	r, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		return err
	}

	// query param handling
	qParams := r.URL.Query()
	for key, value := range req.Parameters {
		qParams.Add(key, value)
	}
	r.URL.RawQuery = qParams.Encode()

	// header handling
	for header, value := range req.Headers {
		// special handling for "host" header if set
		if strings.EqualFold("host", header) {
			r.Host = value
		} else {
			r.Header.Set(header, value)
		}
	}

	if req.Body != "" {
		r.Body = io.NopCloser(strings.NewReader(req.Body))
	}

	res, err := client.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	cui.ResponseStatus.SetText(fmt.Sprintf("Status: %d", res.StatusCode))

	cui.ResponseHeaders.SetCell(0, 0, tview.NewTableCell("Header"))
	cui.ResponseHeaders.SetCell(0, 1, tview.NewTableCell("Value"))

	i := 1
	for k, v := range res.Header {
		for _, vv := range v {
			if strings.EqualFold("content-type", k) {
				if strings.HasPrefix(strings.ToLower(vv), "text/html") {
					highlight = true
					highlightLang = "html"
				} else if strings.HasPrefix(strings.ToLower(vv), "application/json") {
					highlight = true
					highlightLang = "json"
				}
			}

			cui.ResponseHeaders.SetCell(i, 0, tview.NewTableCell(k))
			cui.ResponseHeaders.SetCell(i, 1, tview.NewTableCell(vv))
			i += 1
		}
	}

	if highlight {
		// TODO: need to call app.Draw() in a SetChangedFunc?
		cui.ResponseBody.SetDynamicColors(true)

		lexer := lexers.Get(highlightLang)
		style := styles.Get("fruity")
		formatter := formatters.Get("terminal")
		iterator, err := lexer.Tokenise(nil, string(body))

		w := tview.ANSIWriter(cui.ResponseBody)
		err = formatter.Format(w, style, iterator)
		if err != nil {
			return err
		}

		cui.ResponseBody.ScrollToBeginning()
	} else {
		cui.ResponseBody.SetText(string(body)).ScrollToBeginning()
	}

	*hasResponse = true

	return nil
}
