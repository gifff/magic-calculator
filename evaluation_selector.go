package main

import (
	"io"

	"github.com/gifff/magic-calculator/evaluation"
)

func selectEvaluator(cmd string, input []int64, resultWriter io.Writer) evaluation.Evaluator {
	command, ok := commands[cmd]
	if !ok {
		panic("invalid command")
	}

	return command.EvaluatorFactoryFn(input, resultWriter)
}
