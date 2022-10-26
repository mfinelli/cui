package main

import (
	"os"
	"path/filepath"
	"strings"
)

func SaveResponseFile(filePath string, responseBody string) error {

	var abs_filepath string

	if strings.HasPrefix(filePath, "~/") {
		dirname, _ := os.UserHomeDir()
		abs_filepath = filepath.Join(dirname, filePath[2:])

		_, errStat := os.Stat(abs_filepath)

		if os.IsNotExist(errStat) {
			f, err := os.Create(abs_filepath)
			if err != nil {
				return err
			}
			defer f.Close()

			_, errWrite := f.WriteString(responseBody)
			if errWrite != nil {
				return errWrite
			}
		}
	} else {

		// fmt.Printf("%v file exists", abs_filepath)
		// return errStat
	}

	return nil

}
