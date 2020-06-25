package main

import (
	"io"

	"github.com/gifff/magic-calculator/evaluation"
)

type EvaluatorFactory func(input []int64, resultWriter io.Writer) evaluation.Evaluator

type Command struct {
	RequiredArgs       int
	EvaluatorFactoryFn EvaluatorFactory
}

var commands = map[string]Command{
	"sum": {
		RequiredArgs: 2,
		EvaluatorFactoryFn: func(input []int64, resultWriter io.Writer) evaluation.Evaluator {
			return &evaluation.SumEvaluator{
				X:            input[0],
				Y:            input[1],
				ResultWriter: resultWriter,
			}
		},
	},
	"multiply": {
		RequiredArgs: 2,
		EvaluatorFactoryFn: func(input []int64, resultWriter io.Writer) evaluation.Evaluator {
			return &evaluation.MultiplyEvaluator{
				X:            input[0],
				Y:            input[1],
				ResultWriter: resultWriter,
			}
		},
	},
	"nfib": {
		RequiredArgs: 1,
		EvaluatorFactoryFn: func(input []int64, resultWriter io.Writer) evaluation.Evaluator {
			return &evaluation.FirstFibonacciEvaluator{
				N:            input[0],
				ResultWriter: resultWriter,
			}
		},
	},
	"nprime": {
		RequiredArgs: 1,
		EvaluatorFactoryFn: func(input []int64, resultWriter io.Writer) evaluation.Evaluator {
			return &evaluation.FirstPrimeEvaluator{
				N:            input[0],
				ResultWriter: resultWriter,
			}
		},
	},
}
