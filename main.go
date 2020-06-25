package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gifff/magic-calculator/evaluation"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printUsage()
		return
	}
	cmd := args[0]
	args = args[1:]

	requiredArgs := 0
	switch cmd {
	case "sum", "multiply":
		requiredArgs = 2
	case "nfib":
		requiredArgs = 1
	case "help":
		printHelp()
		return
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

	var evaluator evaluation.Evaluator
	switch cmd {
	case "sum":
		evaluator = &evaluation.SumEvaluator{
			X:            input[0],
			Y:            input[1],
			ResultWriter: os.Stdout,
		}
	case "multiply":
		evaluator = &evaluation.MultiplyEvaluator{
			X:            input[0],
			Y:            input[1],
			ResultWriter: os.Stdout,
		}
	case "nfib":
		evaluator = &evaluation.FirstFibonacciEvaluator{
			N:            input[0],
			ResultWriter: os.Stdout,
		}
	}

	evaluator.Evaluate()
}

func printUsage() {
	executableName := os.Args[0]
	fmt.Printf("For usage information: %s help\n", filepath.Base(executableName))
}

func printHelp() {
	executableName := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s <command> <arg1> <arg2> ... <argN>\n\n", executableName)
	fmt.Printf("Available commands:\n")
	fmt.Printf("\tsum\t\tSum of <arg1> and <arg2>\n")
	fmt.Printf("\tmultiply\tMultiply <arg1> and <arg2>\n")
	fmt.Printf("\tnfib\t\tPrint the first N Fibonacci number\n")
}
