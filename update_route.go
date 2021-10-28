package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getNames(schemaPath string) (string, string) {
	serviceNamePath := strings.Split(schemaPath, "/")
	serviceName := strings.Split(serviceNamePath[len(serviceNamePath)-1], ".go")[0]
	serviceName2 := fmt.Sprintf("%s%s", strings.ToUpper(string(serviceName[0])), serviceName[1:])
	return serviceName, serviceName2
}

func updateRoute(schemaPath string) error {
	fmt.Println("Updating server routes")
	serviceName, serviceName2 := getNames(schemaPath)
	update := fmt.Sprintf(routeUpdate, serviceName, serviceName2)

	serverFile := "src/server/server.go"

	file, err := os.Open(serverFile)

	if err != nil {
		return err
	}
	outFile, err := os.OpenFile(serverFile, os.O_RDWR, 0777)
	if err != nil {
		return err
	}

	defer file.Close()
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	w := bufio.NewWriter(outFile)
	newString := ""
	shouldUpdate := true
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, fmt.Sprintf("// /%[1]s routes", serviceName)) {
			shouldUpdate = false
			break
		}
		if strings.HasSuffix(line, "// ------> codegen_line_tracker ------->") {
			line = update
		}
		newString += line + "\n"
		// fmt.Fprintln(w, line)
	}
	if shouldUpdate {
		fmt.Println("Updating routes")
		w.WriteString(newString)
		if err := w.Flush(); err != nil {
			return err
		}
	}

	return nil
}
