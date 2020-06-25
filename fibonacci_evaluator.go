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
	fmt.Fprintf(e.ResultWriter, "Result: \n")
}
