package main

import (
	"fmt"
	"io"
)

type SumEvaluator struct {
	Result       SingleResult
	X, Y         int64
	ResultWriter io.Writer
}

func (e *SumEvaluator) Evaluate() {
	fmt.Fprintf(e.ResultWriter, "Result: %d\n", 0)
}