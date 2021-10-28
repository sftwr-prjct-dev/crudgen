package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var routePathContent = map[string]string{
	"src/services/%s_get.go":    getRouteContent,
	"src/services/%s_post.go":   postRouteContent,
	"src/services/%s_put.go":    updateRouteContent,
	"src/services/%s_delete.go": deleteRouteContent,
}

func generateCRUD(schemaPath string) error {
	fmt.Println("Generating CRUD")
	serviceName, serviceName2 := getNames(schemaPath)

	for routePath, routeContent := range routePathContent {
		getPath := fmt.Sprintf(routePath, serviceName)
		if _, err := os.Stat(getPath); errors.Is(err, os.ErrNotExist) {
			outFile, err := os.Create(getPath)

			if err != nil {
				return err
			}
			w := bufio.NewWriter(outFile)
			defer outFile.Close()
			content := fmt.Sprintf(routeContent, serviceName2, serviceName)
			w.WriteString(content)
			if err := w.Flush(); err != nil {
				return err
			}
		}
	}
	return nil
}
