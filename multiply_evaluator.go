package main

import (
	"fmt"
	"io"
)

type MultiplyEvaluator struct {
	Result       SingleResult
	X, Y         int64
	ResultWriter io.Writer
}

func (e *MultiplyEvaluator) Evaluate() {
	fmt.Fprintf(e.ResultWriter, "Result: %d\n", 0)
}
