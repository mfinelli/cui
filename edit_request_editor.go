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

	tmpFile, errTemp := os.CreateTemp("", tmpFileName)

	if errTemp != nil {
		return nil, errTemp
	}
	defer os.Remove(tmpFile.Name())

	_, errWrite := tmpFile.WriteString(tempText)
	if errWrite != nil {
		return nil, errWrite
	}

	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	errRun := cmd.Run()
	if errRun != nil {
		return nil, errRun
	}

	inputText, errReadTemp := os.ReadFile(tmpFile.Name())
	if errReadTemp != nil {
		return nil, errReadTemp
	}

	return inputText, nil

}
