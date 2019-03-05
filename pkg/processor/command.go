package processor

import (
	"fmt"
	"strings"
)

func ProcessCommand(input []string) (cmd string, args []string, err error) {
	if len(input) == 1 && input[0] == "" {
		return cmd, args, fmt.Errorf("Please provide a command.")
	}

	// extract cmd and arguments
	slices := strings.Split(input[0], " ")
	cmd = slices[0]
	args = slices[1:]

	if len(args) == 0 && cmd != "help" {
		return cmd, args, fmt.Errorf("Please provide a %s action.", cmd)
	}

	return cmd, args, nil
}
