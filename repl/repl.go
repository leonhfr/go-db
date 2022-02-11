package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/leonhfr/go-db/metacmd"
	"github.com/leonhfr/go-db/sql"
)

func Prompt(r *bufio.Reader, w io.Writer) error {
	for {
		fmt.Fprint(w, "db> ")

		input, err := readInput(r)
		if err != nil {
			return err
		}
		if input == "" {
			continue
		}

		if metacmd.IsMetaCommand(input) {
			exit, err := metacmd.ExecuteMetaCommand(input)
			if err != nil {
				fmt.Println(err)
			}
			if exit {
				return nil
			}
			continue
		}

		s, err := sql.ParseStatement(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		s.Execute()
	}
}

func readInput(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}
