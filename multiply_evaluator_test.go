package main

import (
	"strings"
	"testing"
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			me := &MultiplyEvaluator{X: tc.x, Y: tc.y, ResultWriter: sb}
			var evaluator Evaluator = me
			evaluator.Evaluate()

			if me.Result != tc.expectedResult {
				t.Errorf("Expected: %d. Got: %d", tc.expectedResult, me.Result)
			}

			if gotWrittenResult := sb.String(); gotWrittenResult != tc.expectedWrittenResult {
				t.Errorf("Expected: %q. Got: %q", tc.expectedWrittenResult, gotWrittenResult)
			}
		})
	}
}
