package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UpdateSchema(schemaPath string) error {
	schemaPathText := fmt.Sprintf("Updating the schema in: %s", schemaPath)
	fmt.Println(schemaPathText)
	file, err := os.Open(schemaPath)

	if err != nil {
		return err
	}
	outFile, err := os.OpenFile(schemaPath, os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	defer file.Close()
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	w := bufio.NewWriter(outFile)
	updated := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, "``") {
			updated = true
			name := (strings.Fields(line)[0])
			name2 := name
			if name == "ID" || name == "UID" {
				name = strings.ToLower(name)
				name2 = name
				if name == "id" {
					name2 = "_" + name + ",omitempty"
					name = "-"
				}
			} else {
				name = fmt.Sprintf("%s%s", strings.ToLower(string(name[0])), name[1:])
				name2 = name
			}
			replacement := fmt.Sprintf("`json:\"%s,omitempty\" bson:\"%s\"`", name, name2)
			line = strings.ReplaceAll(line, "``", replacement)
		}
		fmt.Fprintln(w, line)
	}
	if updated {
		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}
