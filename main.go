package main

import (
	"fmt"
)

const (
	helpMenu = `gostruct: Tool to convert JSON objects to Go Structs
Usage:
  Pipe input
    echo '{"id": 5, "name": "test"}' | gostruct
  Manual input (copy-paste your struct then Ctrl+D for EOF)
    gostruct --dump
  Print this menu
    gostruct --help
`
)

func main() {
	b, err := getUserInput()
	if err != nil {
		fmt.Print(helpMenu)
		fmt.Printf("\nfailed to get input: %v", err)
		return
	}

	jsonObj, err := attemptJsonDecode(b)
	if err != nil {
		fmt.Printf("failed to unmarshal JSON object: %v", err)
		return
	}

	fmt.Print(mapToStructDef(jsonObj))
}
