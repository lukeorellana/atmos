package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// PrintAsJSON prints the provided value as YAML document to the console
func PrintAsJSON(data interface{}) error {
	j, err := json.MarshalIndent(data, "", strings.Repeat(" ", 2))
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	return nil
}

// WriteToFileAsJSON converts the provided value to YAML and writes it to the provided file
func WriteToFileAsJSON(filePath string, data interface{}, fileMode os.FileMode) error {
	j, err := json.MarshalIndent(data, "", strings.Repeat(" ", 2))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, j, fileMode)
	if err != nil {
		return err
	}
	return nil
}
