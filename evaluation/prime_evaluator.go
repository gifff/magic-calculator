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
	fmt.Fprintf(e.ResultWriter, "Result: \n")
}
