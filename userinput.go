package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func newPipeReader() (*bufio.Reader, error) {
	// Ref: https://flaviocopes.com/go-shell-pipes/
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat stdin: %w", err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil, fmt.Errorf("not piped input")
	}
	return bufio.NewReader(os.Stdin), nil
}

func readInputFromPipe() ([]byte, error) {
	reader, err := newPipeReader()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get pipe: %v", err)
	}

	b, err := reader.ReadBytes(0x00)
	if err != nil && err != io.EOF {
		return []byte{}, fmt.Errorf("failed to read from pipe: %v", err)
	}

	return b, nil
}

func readInputFromUser() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	b, err := reader.ReadBytes(0x00)
	if err != nil && err != io.EOF {
		return []byte{}, fmt.Errorf("failed to read from input: %v", err)
	}

	return b, nil
}

func getUserInput() ([]byte, error) {
	args := os.Args[1:]

	if len(args) > 1 {
		return []byte{}, fmt.Errorf("too many arguments")
	}
	if len(args) == 1 {
		switch args[0] {
		case "--dump":
			return readInputFromUser()
		case "--help":
			fmt.Print(helpMenu)
			os.Exit(0)
		default:
			return []byte{}, fmt.Errorf("unrecognised argument: %s", args[0])
		}
	}

	return readInputFromPipe()
}
