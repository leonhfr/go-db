package metacmd

import (
	"fmt"
	"strings"
)

type MetaCommand string

const (
	META_EXIT MetaCommand = ".exit" // Exits gracefully
)

func IsMetaCommand(input string) bool {
	return strings.HasPrefix(input, ".")
}

func ExecuteMetaCommand(cmd string) (bool, error) {
	switch MetaCommand(cmd) {
	case META_EXIT:
		return true, nil
	default:
		return false, fmt.Errorf("unrecognized meta command '%s'", cmd)
	}
}
