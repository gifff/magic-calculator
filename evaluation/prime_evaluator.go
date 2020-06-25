package evaluation

import (
	"fmt"
	"io"
)

type FirstPrimeEvaluator struct {
	Result       SequenceResult
	N            int64
	ResultWriter io.Writer
}

func (e *FirstPrimeEvaluator) Evaluate() {
	e.Result = SequenceResult{}
	if e.N == 1 {
		e.Result = append(e.Result, 2)
	}
	fmt.Fprintf(e.ResultWriter, "Result: %s\n", FormatSequenceResult(e.Result))
}
