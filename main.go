package main

import (
	"bufio"
	"log"
	"os"

	"github.com/leonhfr/go-db/repl"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := os.Stderr

	err := repl.Prompt(r, w)
	if err != nil {
		log.Fatal(err)
	}
}
