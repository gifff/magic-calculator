package evaluation

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyEvaluator(t *testing.T) {

	testCases := []struct {
		name                  string
		x, y                  int64
		expectedResult        SingleResult
		expectedWrittenResult string
	}{
		{
			name:                  "x=0 and y=0 result should be 0",
			x:                     0,
			y:                     0,
			expectedResult:        SingleResult(0),
			expectedWrittenResult: "Result: 0\n",
		},
		{
			name:                  "x=1 and y=1 result should be 1",
			x:                     1,
			y:                     1,
			expectedResult:        SingleResult(1),
			expectedWrittenResult: "Result: 1\n",
		},
		{
			name:                  "x=1 and y=-5 result should be -5",
			x:                     1,
			y:                     -5,
			expectedResult:        SingleResult(-5),
			expectedWrittenResult: "Result: -5\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			me := &MultiplyEvaluator{X: tc.x, Y: tc.y, ResultWriter: sb}
			var evaluator Evaluator = me
			err := evaluator.Evaluate()

			assert.Nil(t, err)
			assert.Equal(t, tc.expectedResult, me.Result)
			assert.Equal(t, tc.expectedWrittenResult, sb.String())
		})
	}
}
