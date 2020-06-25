package main

import (
	"strings"
	"testing"
)

func TestSumEvaluator(t *testing.T) {

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
			name:                  "x=2 and y=0 result should be 2",
			x:                     2,
			y:                     0,
			expectedResult:        SingleResult(2),
			expectedWrittenResult: "Result: 2\n",
		},
		{
			name:                  "x=2 and y=3 result should be 5",
			x:                     2,
			y:                     3,
			expectedResult:        SingleResult(5),
			expectedWrittenResult: "Result: 5\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			se := &SumEvaluator{X: tc.x, Y: tc.y, ResultWriter: sb}
			var evaluator Evaluator = se
			evaluator.Evaluate()

			if se.Result != tc.expectedResult {
				t.Errorf("Expected: %d. Got: %d", tc.expectedResult, se.Result)
			}

			if gotWrittenResult := sb.String(); gotWrittenResult != tc.expectedWrittenResult {
				t.Errorf("Expected: %q. Got: %q", tc.expectedWrittenResult, gotWrittenResult)
			}
		})
	}
}
