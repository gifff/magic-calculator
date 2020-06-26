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
		printHelp(false)
		return
	} else if cmd == "interactive" {
		runInteractive()
		return
	}

	if !checkRequiredArgs(cmd, args) {
		printUsage()
		return
	}

	err := evaluate(cmd, args)
	if err != nil {
		printError(err)
	}
}

func evaluate(cmd string, args []string) error {
	input, err := parseInput(cmd, args)
	if err != nil {
		return err
	}

	evaluator := selectEvaluator(cmd, input, os.Stdout)
	return evaluator.Evaluate()
}
