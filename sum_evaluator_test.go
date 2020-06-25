package main

import (
	"strings"
	"testing"
)

func TestSumEvaluator(t *testing.T) {
	t.Run("x=0 and y=0 result should be 0", func(t *testing.T) {
		sb := &strings.Builder{}
		se := &SumEvaluator{X: 0, Y: 0, ResultWriter: sb}
		var evaluator Evaluator = se
		evaluator.Evaluate()

		expectedResult := SingleResult(0)
		if se.Result != expectedResult {
			t.Errorf("Expected: %d. Got: %d", expectedResult, se.Result)
		}

		expectedWrittenResult := "Result: 0\n"
		if gotWrittenResult := sb.String(); gotWrittenResult != expectedWrittenResult {
			t.Errorf("Expected: %q. Got: %q", expectedWrittenResult, gotWrittenResult)
		}
	})

	t.Run("x=2 and y=0 result should be 2", func(t *testing.T) {
		sb := &strings.Builder{}
		se := &SumEvaluator{X: 2, Y: 0, ResultWriter: sb}
		var evaluator Evaluator = se
		evaluator.Evaluate()

		expectedResult := SingleResult(2)
		if se.Result != expectedResult {
			t.Errorf("Expected: %d. Got: %d", expectedResult, se.Result)
		}

		expectedWrittenResult := "Result: 2\n"
		if gotWrittenResult := sb.String(); gotWrittenResult != expectedWrittenResult {
			t.Errorf("Expected: %q. Got: %q", expectedWrittenResult, gotWrittenResult)
		}
	})
}
