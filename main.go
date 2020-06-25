package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func printUsage() {
	executableName := os.Args[0]
	fmt.Printf("For usage information: %s help\n", filepath.Base(executableName))
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		printUsage()
		return
	}
	cmd := args[0]
	args = args[1:]

	requiredArgs := 0
	switch cmd {
	case "sum", "multiply":
		requiredArgs = 2
	default:
		printUsage()
		return
	}

	if len(args) < requiredArgs {
		printUsage()
		return
	}

	input := make([]int64, 0, len(args))
	for _, arg := range args {
		v, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			panic(err)
		}

		input = append(input, v)
	}

	var evaluator Evaluator
	switch cmd {
	case "sum":
		evaluator = &SumEvaluator{
			X:            input[0],
			Y:            input[1],
			ResultWriter: os.Stdout,
		}
	case "multiply":
		evaluator = &MultiplyEvaluator{
			X:            input[0],
			Y:            input[1],
			ResultWriter: os.Stdout,
		}
	}

	evaluator.Evaluate()
}
