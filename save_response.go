package main

import (
	"os"
	"path/filepath"
	"strings"
)

func SaveResponseFile(filePath string, responseBody string) error {

	var abs_filepath string = filePath

	if strings.HasPrefix(filePath, "~/") {
		dirname, _ := os.UserHomeDir()
		abs_filepath = filepath.Join(dirname, filePath[2:])

	}

	// relative paths not supported

	_, errStat := os.Stat(abs_filepath)

	if errStat == nil {
		return os.ErrExist
	}

	f, err := os.Create(abs_filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, errWrite := f.WriteString(responseBody)
	if errWrite != nil {
		return errWrite
	}

	return nil

}

// func ReplaceFile() {}
