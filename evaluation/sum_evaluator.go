package evaluation

import (
	"fmt"
	"io"
)

type SumEvaluator struct {
	Result       SingleResult
	X, Y         int64
	ResultWriter io.Writer
}

func (e *SumEvaluator) Evaluate() error {
	e.Result = SingleResult(e.X + e.Y)
	fmt.Fprintf(e.ResultWriter, "%d\n", e.Result)
	return nil
}
