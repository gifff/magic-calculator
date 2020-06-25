package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printUsage()
		return
	}
	cmd := args[0]
	args = args[1:]

	if cmd == "help" {
		printHelp()
		return
	}

	if !checkRequiredArgs(cmd, args) {
		printUsage()
		return
	}

	input, err := parseInput(cmd, args)
	if err != nil {
		printError(err)
	}

	evaluator := selectEvaluator(cmd, input, os.Stdout)
	err = evaluator.Evaluate()
	if err != nil {
		printError(err)
	}
}
