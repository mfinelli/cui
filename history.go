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
	"time"
)

func setupRequestHistory(cui *cuiApp) error {
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

		text := fmt.Sprintf("%s: %s", req.Method, req.URL)
		second := fmt.Sprintf("%d %s", req.StatusCode, timeOutput(req.Timestamp))

		cui.RequestHistory.InsertItem(-1, text, second, 0, nil)
	}

	return nil
}

func timeOutput(timestamp int64) string {
	// TODO if it's today just show the time (e.g., 12:30)
	// if it's more than today but the same week show the day (e.g., Mon 19:00)
	// if it's more than a week show the full date (e.g., 1 Feb 6:00)
	t := time.Unix(timestamp, 0)
	return t.Format(time.UnixDate)
}
