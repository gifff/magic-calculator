package evaluation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstPrimeEvaluator(t *testing.T) {
	testCases := []struct {
		name                  string
		n                     int64
		expectedResult        SequenceResult
		expectedWrittenResult string
	}{
		{
			name:                  "n=0 result should be empty",
			n:                     0,
			expectedResult:        SequenceResult{},
			expectedWrittenResult: "Result: \n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			e := &FirstPrimeEvaluator{N: tc.n, ResultWriter: sb}
			var evaluator Evaluator = e
			evaluator.Evaluate()

			assert.NotNil(t, e.Result)
			assert.ElementsMatch(t, tc.expectedResult, e.Result)
			assert.Equal(t, tc.expectedWrittenResult, sb.String())
		})
	}
}
