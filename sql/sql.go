package sql

import (
	"fmt"
	"strings"
)

type statementKind string

const (
	STATEMENT_INSERT statementKind = "insert" // Insert statement
	STATEMENT_SELECT statementKind = "select" // Select statement
)

type Statement struct {
	kind statementKind
}

func ParseStatement(input string) (*Statement, error) {
	if strings.HasPrefix(input, string(STATEMENT_INSERT)) {
		return &Statement{STATEMENT_INSERT}, nil
	}

	if strings.HasPrefix(input, string(STATEMENT_SELECT)) {
		return &Statement{STATEMENT_SELECT}, nil
	}

	return nil, fmt.Errorf("unrecognized statement '%s'", input)
}

func (s *Statement) Execute() {
	switch s.kind {
	case STATEMENT_INSERT:
		fmt.Println("This is where we would do an insert.")
	case STATEMENT_SELECT:
		fmt.Println("This is where we would do a select.")
	}
}
