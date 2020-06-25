package main

import "strconv"

func checkRequiredArgs(cmd string, args []string) bool {
	command, ok := commands[cmd]
	if !ok {
		return false
	}

	return command.RequiredArgs >= len(args)
}

func parseInput(cmd string, args []string) ([]int64, error) {
	input := make([]int64, 0, len(args))
	for _, arg := range args {
		v, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return nil, err
		}

		input = append(input, v)
	}

	return input, nil
}
