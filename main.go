package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prompt()
}

func prompt() error {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprintf(os.Stderr, "db> ")

		input, err := r.ReadString('\n')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(input)
		if input == ".exit" {
			break
		}

		if input == "" {
			continue
		}

		fmt.Printf("unrecognized command '%s'\n", input)
	}

	return nil
}
