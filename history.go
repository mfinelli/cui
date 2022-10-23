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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/rivo/tview"
)

func insertHistoryItem(app *tview.Application, store *cuiStoredRequest, cui *cuiApp, req *cuiRequest) {
	text := fmt.Sprintf("%s: %s", store.Method, store.URL)
	second := fmt.Sprintf("%d %s", store.StatusCode, timeOutput(store.Timestamp))

	cui.RequestHistory.InsertItem(0, text, second, 0, func() {
		req = initRequest(app, cui, store.Method, store.URL, store.Body, store.Parameters, store.Headers)
		setInstructions(cui, "")
		app.SetFocus(cui.Main)
	})
}

func setupRequestHistory(app *tview.Application, cui *cuiApp, cuiReq *cuiRequest) error {
	cui.RequestHistory.Clear()

	// we assume this directory exists because we created it on startup
	// for the logfile
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}
	cacheDir = filepath.Join(cacheDir, "cui")

	files, err := filepath.Glob(filepath.Join(cacheDir, "request-*.json"))
	if err != nil {
		return err
	}

	sort.Strings(files)

	// TODO: we probably need to sort this list
	for _, file := range files {
		req := cuiStoredRequest{}

		b, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(b), &req)
		if err != nil {
			return err
		}

		// TODO: we could do a version check here to make sure that
		// we can hndle the format
		insertHistoryItem(app, &req, cui, cuiReq)
	}

	cui.RequestHistory.SetCurrentItem(0)

	return nil
}

func timeOutput(timestamp int64) string {
	// TODO if it's today just show the time (e.g., 12:30)
	// if it's more than today but the same week show the day (e.g., Mon 19:00)
	// if it's more than a week show the full date (e.g., 1 Feb 6:00)
	t := time.Unix(timestamp, 0)
	return t.Format(time.UnixDate)
}
