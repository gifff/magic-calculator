package main

import (
	"fmt"
	"io"
)

type FirstFibonacciEvaluator struct {
	Result       SequenceResult
	N            int64
	ResultWriter io.Writer
}

func (e *FirstFibonacciEvaluator) Evaluate() {
	e.Result = []int64{}
	if e.N == 1 {
		e.Result = append(e.Result, 0)
	}

	fmt.Fprintf(e.ResultWriter, "Result: %v\n", FormatSequenceResult(e.Result))
}
