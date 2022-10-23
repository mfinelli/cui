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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/rivo/tview"
)

type cuiRequest struct {
	Method string `json:"method"`
	URL    string `json:"url"`

	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Parameters map[string]string `json:"parameters"`
}

type cuiStoredRequest struct {
	CuiVersion string `json:"version"`

	StatusCode int   `json:"status"`
	Timestamp  int64 `json:"timestamp"`

	Method string `json:"method"`
	URL    string `json:"url"`

	Parameters map[string]string `json:"parameters"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

func sendRequest(req cuiRequest, cui *cuiApp, hasResponse *bool) error {
	// we assume this directory exists because we created it on startup
	// for the logfile
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	cacheDir = filepath.Join(cacheDir, "cui")

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

	store := cuiStoredRequest{
		CuiVersion: version,
		StatusCode: res.StatusCode,
		Timestamp:  time.Now().Unix(),
		Method:     req.Method,
		URL:        req.URL,
		Parameters: req.Parameters,
		Headers:    req.Headers,
		Body:       req.Body,
	}

	jsonBytes, err := json.Marshal(store)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(cacheDir, fmt.Sprintf("request-%d.json", store.Timestamp)), jsonBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
