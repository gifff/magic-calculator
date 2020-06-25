package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciEvaluator(t *testing.T) {
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
		{
			name:                  "n=1 result should be '0'",
			n:                     1,
			expectedResult:        SequenceResult{0},
			expectedWrittenResult: "Result: 0\n",
		},
		{
			name:                  "n=2 result should be '0, 1'",
			n:                     2,
			expectedResult:        SequenceResult{0, 1},
			expectedWrittenResult: "Result: 0, 1\n",
		},
		{
			name:                  "n=5 result should be '0, 1, 1, 2, 3'",
			n:                     5,
			expectedResult:        SequenceResult{0, 1, 1, 2, 3},
			expectedWrittenResult: "Result: 0, 1, 1, 2, 3\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sb := &strings.Builder{}
			e := &FirstFibonacciEvaluator{N: tc.n, ResultWriter: sb}
			var evaluator Evaluator = e
			evaluator.Evaluate()

			assert.NotNil(t, e.Result)
			assert.ElementsMatch(t, tc.expectedResult, e.Result)
			assert.Equal(t, tc.expectedWrittenResult, sb.String())
		})
	}

	t.Run("n=-3 should be panic", func(t *testing.T) {
		sb := &strings.Builder{}
		e := &FirstFibonacciEvaluator{N: -3, ResultWriter: sb}
		var evaluator Evaluator = e

		assert.PanicsWithError(t, ErrInvalidInput.Error(), func() {
			evaluator.Evaluate()
		})
	})
}
