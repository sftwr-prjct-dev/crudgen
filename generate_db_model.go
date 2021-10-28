package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func generateDBModel(schemaPath string) error {
	fmt.Println("Generating DB models")
	serviceName, serviceNameTitle := getNames(schemaPath)

	dbPath := fmt.Sprintf("src/models/%ss.go", serviceName)
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		outFile, err := os.Create(dbPath)

		if err != nil {
			return err
		}
		w := bufio.NewWriter(outFile)
		defer outFile.Close()
		content := fmt.Sprintf(modelContent, serviceNameTitle, serviceName)
		w.WriteString(content)
		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}
