// cui: http request/response tui
// Copyright 2022-2023 Mario Finelli
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
	"os/exec"
)

func EditRequestInEditor(currentOption string, tempText string, editor string) ([]byte, error) {
	tmpFileName := "editRequest"

	if currentOption == "JSON" {
		tmpFileName = "editRequest*.json"
	}

	tmpFile, err := os.CreateTemp("", tmpFileName)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(tempText)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	inputText, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}

	return inputText, nil
}
