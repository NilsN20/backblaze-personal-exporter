package utils

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func ReadXmlFile[T any](filePath string) T {
	file, err := os.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "it is being used by another process") {
			// backblaze currently using file, this it fine it just won't update the stats for now
			// in the future this should either retry or be shown to prometheus as the job being down
		} else {
			fmt.Printf("error: %v", err)
		}

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	var xmlModel T
	if err := xml.NewDecoder(file).Decode(&xmlModel); err != nil {
		fmt.Printf("error: %v", err)
	}

	return xmlModel
}
