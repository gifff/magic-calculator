package evaluation

import (
	"fmt"
	"io"
)

type FirstFibonacciEvaluator struct {
	Result       SequenceResult
	N            int64
	ResultWriter io.Writer
}

func (e *FirstFibonacciEvaluator) Evaluate() error {
	e.Result = []int64{}

	if e.N < 0 {
		return ErrInvalidInput
	}

	if e.N >= 1 {
		e.Result = append(e.Result, 0)
	}
	if e.N >= 2 {
		e.Result = append(e.Result, 1)
	}

	for i := int64(2); i < e.N; i++ {
		e.Result = append(e.Result, e.Result[i-1]+e.Result[i-2])
	}

	fmt.Fprintf(e.ResultWriter, "%v\n", FormatSequenceResult(e.Result))
	return nil
}
