package main

import (
	"fmt"
	"os"
)

// schema update
// create route
// generate CRUD files
// generate db files

func main() {
	schemaPath := "NO_PATH"
	if len(os.Args) > 1 {
		schemaPath = os.Args[1]
	}
	err := UpdateSchema(schemaPath)
	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	err = updateRoute(schemaPath)

	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	err = generateCRUD(schemaPath)

	if err != nil {
		fmt.Println("Err: ", err)
		return
	}

	err = generateDBModel(schemaPath)

	if err != nil {
		fmt.Println("Err: ", err)
		return
	}
}
