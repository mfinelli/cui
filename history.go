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
	"os"
	"path/filepath"
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
		// TODO: main text Method URL
		cui.RequestHistory.InsertItem(-1, file, "TODO: Status: Timestamp", 0, nil)
	}

	return nil
}
