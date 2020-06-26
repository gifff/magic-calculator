package evaluation

import (
	"fmt"
	"io"
)

type MultiplyEvaluator struct {
	Result       SingleResult
	X, Y         int64
	ResultWriter io.Writer
}

func (e *MultiplyEvaluator) Evaluate() error {
	e.Result = SingleResult(e.X * e.Y)
	fmt.Fprintf(e.ResultWriter, "%d\n", e.Result)
	return nil
}
