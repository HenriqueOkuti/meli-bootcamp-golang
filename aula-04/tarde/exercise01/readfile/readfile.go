package readfile

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) string {
	defer func() {
		if r := recover(); r != nil {
			var errMsg = errors.New("File was not found or is damaged")
			fmt.Println(errMsg)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		panic("File was not found or is damaged")
	}

	return convertFileToString(file)
}

func convertFileToString(file *os.File) string {
	defer func() {
		if r := recover(); r != nil {
			errMsg := errors.New("Error reading file")
			fmt.Println(errMsg)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("Error reading file")
	}
	return string(bytes)
}
